package commoncollections

import (
	"testing"
)

func BenchmarkOptLock_LockSync(b *testing.B) {
	lock := NewOptLock()
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}

func BenchmarkOptLock_Lock(b *testing.B) {
	lock := NewOptLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			lock.Unlock()
		}
	})
}

func BenchmarkOptLock_RLock(b *testing.B) {
	lock := NewOptLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			state, _ := lock.RLock()
			lock.RVerify(state)
		}
	})
}

func BenchmarkOptLock_20_80(b *testing.B) {
	lock := NewOptLock()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i++
			if i%5 == 0 {
				lock.Lock()
				lock.Unlock()
				continue
			}
			state, _ := lock.RLock()
			lock.RVerify(state)
		}
	})
}

func BenchmarkOptLock_50_50(b *testing.B) {
	lock := NewOptLock()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i++
			if i%2 == 0 {
				lock.Lock()
				lock.Unlock()
				continue
			}
			state, _ := lock.RLock()
			lock.RVerify(state)
		}
	})
}
