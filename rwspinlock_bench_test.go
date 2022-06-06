package commoncollections

import (
	"sync"
	"testing"
)

func BenchmarkRWSpinLock_LockSync(b *testing.B) {
	lock := NewRWSpinLock()
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}

func BenchmarkRWMutex_LockSync(b *testing.B) {
	lock := sync.RWMutex{}
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}

func BenchmarkRWMutex_Lock(b *testing.B) {
	lock := sync.RWMutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			lock.Unlock()
		}
	})
}

func BenchmarkRWSpinLock_Lock(b *testing.B) {
	lock := NewRWSpinLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			lock.Unlock()
		}
	})
}

func BenchmarkRWMutex_RLock(b *testing.B) {
	lock := sync.RWMutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.RLock()
			lock.RUnlock()
		}
	})
}

func BenchmarkRWSpinLock_RLock(b *testing.B) {
	lock := NewRWSpinLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.RLock()
			lock.RUnlock()
		}
	})
}

func BenchmarkRWSpinlock_20_80(b *testing.B) {
	lock := NewRWSpinLock()
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i++
			if i%2 == 0 {
				lock.Lock()
				lock.Unlock()
			}
			lock.RLock()
			lock.RUnlock()
		}
	})
}

func BenchmarkRWMutex_20_80(b *testing.B) {
	lock := sync.RWMutex{}
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i++
			if i%2 == 0 {
				lock.Lock()
				lock.Unlock()
			}
			lock.RLock()
			lock.RUnlock()
		}
	})
}
