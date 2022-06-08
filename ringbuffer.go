package commoncollections

const maxUint64 = ^uint64(0)

// A RingBuffer is a FIFO queue with a fixed size.
// If the buffer is full, the oldest element is overwritten.
type RingBuffer[T any] struct {
	data []T
	cap  uint64
	wrap bool
	head uint64
	tail uint64
}

// NewRingBuffer creates a new RingBuffer.
// The capacity must be greater than zero.
func NewRingBuffer[T any](capacity uint64) RingBuffer[T] {
	if capacity == 0 {
		panic("size must be greater than zero")
	}
	return RingBuffer[T]{
		data: make([]T, capacity),
		cap:  capacity,
		wrap: false,
		head: 0,
		tail: 0,
	}
}

// Puts adds an element to the buffer.
func (ring *RingBuffer[T]) Put(value T) {
	ring.data[ring.head] = value
	head := ring.head
	ring.head = (ring.head + 1) % ring.cap
	if head == ring.tail && ring.wrap {
		ring.tail = (ring.tail + 1) % ring.cap
	}
	if ring.head == 0 {
		ring.wrap = true
	}
}

// Get reads the last element from the queue removing it.
// If no item is present an invalid value and false is returned.
// Otherwise the last item is returned and true.
func (ring *RingBuffer[T]) Get() (T, bool) {
	value := ring.data[ring.tail]
	if ring.tail == ring.head && !ring.wrap {
		return value, false
	}
	ring.tail = (ring.tail + 1) % ring.cap
	if ring.tail == 0 {
		ring.wrap = false
	}
	return value, true
}

// Size returns the number of elements in the buffer.
func (ring *RingBuffer[T]) Size() uint64 {
	if ring.wrap {
		return ring.head + ring.cap - ring.tail
	}
	return ring.head - ring.tail
}

// Cap returns the maximum number of elements in the buffer.
func (ring *RingBuffer[T]) Cap() uint64 {
	return ring.cap
}
