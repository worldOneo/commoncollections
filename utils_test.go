package commoncollections

import "testing"

func TestSliceEquals(t *testing.T) {
	if SliceEquals([]int{1, 2, 3, 4}, []int{1, 2, 3, 5}) || SliceEquals([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("Uneqaul arrays are equall")
	}
	if !SliceEquals([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}) {
		t.Fatalf("Equal arrays are unequall")
	}
}
