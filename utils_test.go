package commoncollections

import (
	"testing"
)

const FNV_ITERS = 1_000_000

func TestSliceEquals(t *testing.T) {
	if SliceEquals([]int{1, 2, 3, 4}, []int{1, 2, 3, 5}) || SliceEquals([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("Uneqaul arrays are equall")
	}
	if !SliceEquals([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}) {
		t.Fatalf("Equal arrays are unequall")
	}
}

func TestFNV64(t *testing.T) {
	mp := NewSet[uint64]()
	bytes := [8]byte{}
	slice := bytes[:]
	for i := 0; i < FNV_ITERS; i++ {
		if i%100_000 == 0 {
			t.Logf("Iter %d\n", i)
		}
		randomString(slice)
		hash := FNV64(slice)
		if mp.Contains(hash) {
			t.Fatalf("Collision %v %d", string(slice), i)
		}
		mp.Insert(hash)
	}
}

func TestMax(t *testing.T) {
	if Max(1, 1) != 1 || Max(1, 2) != 2 || Max(2, 1) != 2 {
		t.Fatalf("Invalid value returned from max")
	}
}

func TestMin(t *testing.T) {
	if Min(1, 1) != 1 || Min(1, 2) != 1 || Min(2, 1) != 1 {
		t.Fatalf("Invalid value returned from max")
	}
}