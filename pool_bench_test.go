package commoncollections

import "testing"

func BenchmarkPoolPut(b *testing.B) {
	pool := NewPool(func() int {
		return 0
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(i)
	}
}

func BenchmarkPoolSycGet(b *testing.B) {
	pool := NewSyncPool(func() int {
		return 0
	})
	for i := 0; i < b.N; i++ {
		pool.Put(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkPoolSyncPut(b *testing.B) {
	pool := NewSyncPool(func() int {
		return 0
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(i)
	}
}

func BenchmarkPoolRoutinesPut(b *testing.B) {
	pool := NewSyncPool(func() int {
		return 0
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for i := 0; p.Next(); i++ {
			pool.Put(i)
		}
	})
}
