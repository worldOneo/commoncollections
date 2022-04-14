package safecollections

import "testing"

func BenchmarkStack_Push(b *testing.B) {
	fillStack(b.N)
}

func BenchmarkStack_Pop(b *testing.B) {
	stack := fillStack(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}
