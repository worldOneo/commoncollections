package commoncollections

import "testing"

func BenchmarkBKTree_Insert(b *testing.B) {
	tree := NewBKTree(func(a, b int) int {
		return a - b
	})

	for i := 0; i < b.N; i++ {
		tree.Insert(i)
	}
}

func BenchmarkBKTree_Find(b *testing.B) {
	tree := NewBKTree(func(a, b int) int {
		return a - b
	})

	for i := 0; i < b.N; i++ {
		tree.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Find(i, 1)
	}
}