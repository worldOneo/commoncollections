package commoncollections

import (
	"testing"
)

func filled(n uint64) *IntMap[uint64] {
	m := NewIntMap[uint64]()
	for i := uint64(0); i < n; i++ {
		m.Put(i, i)
	}
	return m
}

func filledStd(n uint64) map[uint64]uint64 {
	m := make(map[uint64]uint64)
	for i := uint64(0); i < n; i++ {
		m[i] = i
	}
	return m
}

func BenchmarkStdMap_Put(b *testing.B) {
	filledStd(uint64(b.N))
}

func BenchmarkIntMap_Put(b *testing.B) {
	filled(uint64(b.N))
}

func BenchmarkStdMap_Get(b *testing.B) {
	m := filledStd(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		_, _ = m[i]
	}
}

func BenchmarkIntMap_Get(b *testing.B) {
	m := filled(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		_, _ = m.Get(i)
	}
}

func BenchmarkStdMap_Delete(b *testing.B) {
	m := filledStd(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		delete(m, i)
	}
}

func BenchmarkIntMap_Delete(b *testing.B) {
	m := filled(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		_, _ = m.Delete(i)
	}
}