package commoncollections

import (
	"sync"
	"testing"
)

func BenchmarkLockUnlockMutex(b *testing.B) {
	lock := sync.Mutex{}
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			lock.Lock()
			lock.Unlock()
		}
	})
}

func BenchmarkLockUnlockSpinLock(b *testing.B) {
	lock := SpinLock(0)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			lock.Lock()
			lock.Unlock()
		}
	})
}
