package safecollections

import "testing"

func BenchmarkStack_Push(b *testing.B) {
	fillStack(uint64(b.N))
}

func BenchmarkStack_Pop(b *testing.B) {
	stack := fillStack(uint64(b.N))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}
