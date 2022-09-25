package commoncollections

import "testing"

func TestBkTree(t *testing.T) {
	tree := NewBKTree(func(a, b int) int {
		if a-b < 0 {
			return b - a
		}
		return a - b
	})

	if len(tree.Find(0, 1000)) != 0 {
		t.Fatalf("BKTree.Find(0) want %d", 0)
	}

	for i := 0; i < 1000; i++ {
		tree.Insert(i)
		if !tree.Insert(i) {
			t.Fatalf("BKTree.Insert wanted %t, got %t", true, false)
		}
	}

	for i := 1; i < 999; i++ {
		r := tree.Find(i, 1)
		if len(r) != 3 {
			t.Fatalf("Expected %d, got %d", 3, len(r))
		}
		if r[0] != i-1 || r[1] != i || r[2] != i+1 {
			t.Fatalf("Expected %d, got %d", []int{i - 1, i, i + 1}, r)
		}
	}
}
