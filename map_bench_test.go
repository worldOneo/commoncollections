package commoncollections

import (
	"math/rand"
	"testing"
)

func randomInts(n uint64) []uint64 {
	input := []uint64{}
	for i := uint64(0); i < n; i++ {
		input = append(input, i)
	}
	rand.Shuffle(int(n), func(i, j int) { tmp := input[i]; input[i] = input[j]; input[j] = tmp })
	return input
}

func filled(n uint64) *IntMap[uint64] {
	m := NewIntMap[uint64]()
	input := randomInts(n)
	for i := uint64(0); i < n; i++ {
		m.Put(input[i], input[i])
	}
	return &m
}

func filledMap(n uint64) *Map[uint64, uint64] {
	m := NewMap[uint64, uint64](func(i uint64) uint64 { return i })
	input := randomInts(n)
	for i := uint64(0); i < n; i++ {
		m.Put(input[i], input[i])
	}
	return &m
}

func filledStd(n uint64) map[uint64]uint64 {
	m := make(map[uint64]uint64)
	input := randomInts(n)
	for i := uint64(0); i < n; i++ {
		m[input[i]] = input[i]
	}
	return m
}

func BenchmarkStdMap_Put(b *testing.B) {
	filledStd(uint64(b.N))
}

func BenchmarkIntMap_Put(b *testing.B) {
	filled(uint64(b.N))
}

func BenchmarkMap_Put(b *testing.B) {
	filledMap(uint64(b.N))
}

func BenchmarkStdMap_Get(b *testing.B) {
	m := filledStd(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		_ = m[i]
	}
}

func BenchmarkIntMap_Get(b *testing.B) {
	m := filled(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		_, _ = m.Get(i)
	}
}

func BenchmarkMap_Get(b *testing.B) {
	m := filledMap(uint64(b.N))
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

func BenchmarkMap_Delete(b *testing.B) {
	m := filledMap(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		_, _ = m.Delete(i)
	}
}
