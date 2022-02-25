package commoncollections

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue(0)
	_, ok := queue.Peek()
	if ok {
		t.Fatalf("Could peek into empty queue")
	}
	_, ok = queue.Pop()
	if ok {
		t.Fatalf("Could pop from empty queue")
	}
	for i := 0; i < 33; i++ {
		queue.Push(i)
	}
	val, ok := queue.Peek()
	if val != 0 || !ok {
		t.Fatalf("Could not peek into full queue")
	}
	for i := 0; i < 33; i++ {
		val, ok := queue.Pop()
		if !ok || val != i {
			t.Fatalf("Test failed, got %v %v expected %v %v", val, ok, i, true)
		}
	}
	_, ok = queue.Pop()
	if ok {
		t.Fatalf("Test failed, got %v expected %v", ok, true)
	}
	_, ok = queue.Peek()
	if ok {
		t.Fatalf("Could peek into emptied queue")
	}
}
