package commoncollections

import (
	"sync"
	"testing"
)

func BenchmarkPoolPut(b *testing.B) {
	pool := NewPool(func() int {
		return 0
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(i)
	}
}

func BenchmarkPoolSyncGet(b *testing.B) {
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

func BenchmarkPoolRoutinesGet(b *testing.B) {
	pool := NewSyncPool(func() int {
		return 0
	})
	for i := 0; i < b.N; i++ {
		pool.Put(i)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			pool.Get()
		}
	})
}

func BenchmarkStdPoolRoutinesGet(b *testing.B) { 
	pool := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	for i := 0; i < b.N; i++ {
		pool.Put(i)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			pool.Get()
		}
	})
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

func BenchmarkStdPoolRoutinesPut(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for i := 0; p.Next(); i++ {
			pool.Put(i)
		}
	})
}
