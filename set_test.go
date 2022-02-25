package commoncollections

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	if set.Contains(0) {
		t.Fatalf("Empty set contains 0")
	}
	for i := 0; i < 33; i++ {
		set.Insert(i)
	}
	for i := 0; i < 33; i++ {
		if !set.Contains(i) {
			t.Fatalf("Set doesnt contain %d", i)
		}
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
