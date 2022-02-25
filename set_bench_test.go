package commoncollections

import "testing"

func BenchmarkSetInsert(b *testing.B) {
	set := NewSet[int]()
	for i := 0; i < b.N; i++ {
		set.Insert(i)
	}
}

func BenchmarkSetContains(b *testing.B) {
	set := NewSet[int]()
	for i := 0; i < b.N; i++ {
		set.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(i)
	}
}

func BenchmarkSetRemove(b *testing.B) {
	set := NewSet[int]()
	for i := 0; i < b.N; i++ {
		set.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(i)
	}
}