package commoncollections

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue(0)
	for i := 0; i < 33; i++ {
		queue.Push(i)
	}
	for i := 0; i < 33; i++ {
		val, ok := queue.Pop()
		if !ok || val != i {
			t.Fatalf("Test failed, got %v %v expected %v %v", val, ok, i, true)
		}
	}
	_, ok := queue.Pop()
	if ok {
		t.Fatalf("Test failed, got %v expected %v", ok, true)
	}
}
