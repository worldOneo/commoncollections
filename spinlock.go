package commoncollections

import (
	"sync/atomic"
)

// SpinLock implements and statisfies sync.Locker
// using a SpinLock as described here https://en.wikipedia.org/wiki/Spinlock
// with an exponential backoff described here https://en.wikipedia.org/wiki/Exponential_backoff
//
// The zero-value of the spinlock is valid and can be initialised with SpinLock(0)
// SpinLock must not be copied after first use.
type SpinLock uint32

const maxSchedules = 16

// Lock locks the SpinLock or waits until its available
// using exponential backoff.
func (s *SpinLock) Lock() {
	schedule := 1
	for !atomic.CompareAndSwapUint32((*uint32)(s), 0, 1) {
		schedule = spin(schedule)
	}
}

// Unlock unlocks the locks.
// Panics when the lock wasn't locked
func (s *SpinLock) Unlock() {
	new := atomic.SwapUint32((*uint32)(s), 0)
	if new != 1 {
		panic("SpinLock: Unlock of unlocked lock")
	}
}

// NewSpinLock creates a new spinlock
func NewSpinLock() *SpinLock {
	lock := SpinLock(0)
	return &lock
}
