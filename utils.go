package commoncollections

import (
	"golang.org/x/exp/constraints"
)

// SliceEquals returns if two slices are equal
func SliceEquals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

const (
	// Offset64 is the offset for FNV64
	Offset64 = 14695981039346656037
	// Prime64 is the prime for FNV64
	Prime64 = 1099511628211
)

// FNV64 hash function
func FNV64(key []byte) uint64 {
	var hash uint64 = Offset64
	l := len(key)
	for i := 0; i < l; i++ {
		hash ^= uint64(key[i])
		hash *= Prime64
	}
	return hash
}

// Max function for any type.
// Returns the bigger value of a and b.
// Returns b if equal.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min function for any type.
// Returns the smaller value of a and b.
// Returns b if equal.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
