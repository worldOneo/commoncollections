package commoncollections

import "testing"

func BenchmarkSliceEquals(b *testing.B) {
	a := make([]int, 100, 100)
	c := make([]int, 100, 100)
	for i := range a {
		a[i] = i
	}
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEquals(a, c)
	}
}

func BenchmarkSliceFastNequals(b *testing.B) {
	a := make([]int, 100, 100)
	c := make([]int, 101, 101)
	for i := range a {
		a[i] = i
	}
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEquals(a, c)
	}
}