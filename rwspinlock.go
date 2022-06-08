package commoncollections

import (
	"runtime"
	"sync/atomic"
)

// RWSpinLock type is a read/write lock.
// It is more efficient than sync.RWMutex
// but has less safety checks.
type RWSpinLock int32

const rwUnhold = int32(1_000_000_000)

// NewRWSpinLock creates a new RWSpinLock.
func NewRWSpinLock() RWSpinLock {
	return RWSpinLock(rwUnhold)
}

func spin(backoff int) {
	for i := 0; i < backoff; i++ {
		runtime.Gosched()
	}
}

// RLock locks the RWSpinLock for reading.
// Blocks if the lock is held for writing until
// the lock is unlocked.
func (lock *RWSpinLock) RLock() {
	for {
		if atomic.AddInt32((*int32)(lock), -1) > 0 {
			return
		}
		atomic.AddInt32((*int32)(lock), 1)
		backoff := 1
		for atomic.LoadInt32((*int32)(lock)) <= 0 {
			spin(backoff)
			backoff = Min(backoff<<1, maxSchedules)
		}
	}
}

// RUnlock unlocks the RWSpinLock for reading.
// Panics if the RWSpinLock is not locked for reading.
func (lock *RWSpinLock) RUnlock() {
	if atomic.AddInt32((*int32)(lock), 1) > rwUnhold {
		panic("RWSpinLock: runlock of unheld lock")
	}
}

// Lock locks the RWSpinLock for writing.
// It blocks until no readers are reading.
func (lock *RWSpinLock) Lock() {
	backoff := 1
	for {
		if atomic.AddInt32((*int32)(lock), -rwUnhold) == 0 {
			return
		}
		atomic.AddInt32((*int32)(lock), rwUnhold)
		spin(backoff)
		backoff = Min(backoff<<1, maxSchedules)
	}
}

// Unlock unlocks the RWSpinLock for writing.
func (lock *RWSpinLock) Unlock() {
	atomic.AddInt32((*int32)(lock), rwUnhold)
}
