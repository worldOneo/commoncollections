package commoncollections

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	if set.Contains(0) {
		t.Fatalf("Empty set contains 0")
	}
	keys := make([]int, 33)
	for i := 0; i < 33; i++ {
		keys[i] = i
		set.Insert(i)
		if set.Len() != i+1 {
			t.Fatalf("Size of set wrong, expected %v got %v", i+1, set.Len())
		}
	}
	for i := 0; i < 33; i++ {
		if !set.Contains(i) {
			t.Fatalf("Set doesnt contain %d", i)
		}
	}
	i := 0
	for _, v := range set.Values() {
		i += 1 << v
	}
	if i != 0x1ffffffff {
		t.Fatalf("Not all values present in set value: %v", set.Values())
	}
	for i := 0; i < 33; i++ {
		set.Remove(i)
	}
	for i := 0; i < 33; i++ {
		if set.Contains(i) {
			t.Fatalf("Set contains %d", i)
		}
	}
}
