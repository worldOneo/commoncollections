package commoncollections

import (
	"testing"
	"time"
)

var seed = uint64(time.Now().Unix())

func fastrand() uint64 {
	seed = (214013*seed + 2531011)
	return (seed >> 16) & 0x7FFF
}

func randomString(into []byte) {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	for i := 0; i < len(into); i++ {
		into[i] = letters[fastrand()%uint64(len(letters))]
	}
}

func BenchmarkSliceEquals(b *testing.B) {
	a := make([]int, 100)
	c := make([]int, 100)
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
	a := make([]int, 100)
	c := make([]int, 101)
	for i := range a {
		a[i] = i
	}
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEquals(a, c)
	}
}

func BenchmarkFNV64(b *testing.B) {
	slice := make([]byte, b.N)
	randomString(slice)
	b.ResetTimer()
	FNV64(slice)
}
