package commoncollections

import "testing"

func BenchmarkQueuePush(b *testing.B) {
	queue := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
}


func BenchmarkQueuePop(b *testing.B) {
	queue := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Pop()
	}
}
