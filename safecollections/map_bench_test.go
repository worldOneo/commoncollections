package safecollections

import "testing"

func BenchmarkMap_Put(b *testing.B) {
	filled(uint64(b.N))
}

func BenchmarkMap_Get(b *testing.B) {
	m := filled(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		_, _ = m.Get(i)
	}
}

func BenchmarkMap_Delete(b *testing.B) {
	m := filled(uint64(b.N))
	b.ResetTimer()
	for i := uint64(0); i < uint64(b.N); i++ {
		m, _, _ = m.Delete(i)
	}
}
