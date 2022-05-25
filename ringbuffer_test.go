package commoncollections

import "testing"

func TestRingBuffer_Size(t *testing.T) {
	buffer := NewRingBuffer[int](3)
	buffer.Put(1)
	buffer.Put(2)
	if buffer.Size() != 2 {
		t.Errorf("buffer.Size() != 2, got %d", buffer.Size())
	}
	buffer.Put(3)
	buffer.Put(4)
	if buffer.Size() != 3 {
		t.Errorf("buffer.Size() != 3, got %d", buffer.Size())
	}
	buffer.Get()
	if buffer.Size() != 2 {
		t.Errorf("buffer.Size() != 2, got %d", buffer.Size())
	}
	buffer.Get()
	buffer.Get()
	if buffer.Size() != 0 {
		t.Errorf("buffer.Size() != 0, got %d", buffer.Size())
	}
	if buffer.Cap() != 3 {
		t.Errorf("buffer.Cap() != 3, got %d", buffer.Cap())
	}
}

func TestRingBuffer_Content(t *testing.T) {
	buffer := NewRingBuffer[int](3)
	for i := 0; i < 4; i++ {
		buffer.Put(i)
	}
	for i := 0; i < 3; i++ {
		value, ok := buffer.Get()
		if !ok {
			t.Errorf("buffer.Get() returned false")
		}
		if value != i+1 {
			t.Errorf("buffer.Get() != %d, got %d", i+1, value)
		}
	}
	_, ok := buffer.Get()
	if ok {
		t.Errorf("buffer.Get() returned true")
	}
}

func TestRingBuffer_SizeZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewRingBuffer[int](0)
}
