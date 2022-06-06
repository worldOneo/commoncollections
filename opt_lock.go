package commoncollections

import (
	"runtime"
	"sync/atomic"
)

// OptLock type is used for an optimistic concurrency design.
// The lock provides a Read/Write facility.
// Writers are not prevented from acquiring the lock while readers
// are holding it.
// Readers must verify that the lock is still valid
// after reading their data.
//
// Readers must first acquire the state by calling RLock()
// and then verify that the state is still valid after reading
// by calling RVerify().
//
// Writers must first acquire the state by calling Lock()
// and release their lock after writing by calling Unlock().
//
// The null value of the lock is usable.
//
// Data protected by the lock must not panic when reads and writes
// happen concurrently even if those actions result in invalid data.
// The std map therefore cant be protected by the lock.
type OptLock uint32

// NewOptLock creates a new OptLock.
func NewOptLock() *OptLock {
	lock := OptLock(0)
	return &lock
}

// RLock acquires the read lock.
// It returns the state and if the lock was acquired.
// After reading readers must verify that the state is still valid
// by calling RVerify().
func (lock *OptLock) RLock() (uint32, bool) {
	val := atomic.LoadUint32((*uint32)(lock))
	if val&optLockBit != 0 {
		return val, false
	}
	return val, true
}

// RVerify verifies that the state is still valid.
// It returns if read data is still valid (true) or might
// be corrupted by reads (false).
func (lock *OptLock) RVerify(expected uint32) bool {
	return atomic.LoadUint32((*uint32)(lock)) == expected
}

const optLockBit = 1
const optLockMask = ^uint32(1)
const opetLockBackoff = 16

// Lock acquires the write lock.
// This operation blocks until the lock is acquired
// using exponential backoff.
func (lock *OptLock) Lock() {
	backoff := 1
	for {
		old := atomic.LoadUint32((*uint32)(lock)) & optLockMask
		new := old | optLockBit
		if atomic.CompareAndSwapUint32((*uint32)(lock), old, new) {
			return
		}
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		backoff = Min(backoff<<1, opetLockBackoff)
	}
}

// Unlock releases the write lock.
// It panics when the lock wasn't locked.
func (lock *OptLock) Unlock() {
	old := atomic.LoadUint32((*uint32)(lock))
	new := old & optLockMask
	check := old & optLockBit
	if check != optLockBit {
		panic("OptLock.Unlock() called without OptLock.Lock()")
	}
	atomic.StoreUint32((*uint32)(lock), new+2)
}
