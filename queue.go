package commoncollections

// Queue type for an unbound FIFO queue
type Queue[T any] struct {
	nilvalue T
	buff     []T
	read     int
	write    int
}

// NewQueue creates a new Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		nilvalue: *new(T),
		buff:     make([]T, 16),
		read:     0,
		write:    0,
	}
}

// Push adds an element to the back of the queue.
// This might increase the size of the internal buffer.
func (queue *Queue[T]) Push(elem T) {
	queue.buff[queue.write] = elem
	queue.write++
	cap := len(queue.buff)
	queue.write &= cap - 1 // cap power 2 means Q.write %= cap
	if queue.write == queue.read {
		old := queue.buff
		queue.buff = make([]T, cap*2)
		copy(queue.buff, old[:queue.read])
		copy(queue.buff[queue.write+cap:], old[queue.read:])
		queue.read += cap
	}
}

// Pop reads the last element from the queue removing it.
// It returns the element and true if an element is present
// and returns the nilvalue and false if no element is left.
func (queue *Queue[T]) Pop() (T, bool) {
	if queue.read == queue.write {
		return queue.nilvalue, false
	}
	val := queue.buff[queue.read]
	queue.read++
	if queue.read >= len(queue.buff) {
		queue.read = 0
	}
	return val, true
}

// Peek reads the last element from the queue without changing
// the queue.
// Returns the element and true if an element is present
// and returns the nilvalue and false if no element is left.
func (queue *Queue[T]) Peek() (T, bool) {
	if queue.read == queue.write {
		return queue.nilvalue, false
	}
	return queue.buff[queue.read], true
}

func (queue *Queue[T]) Len() int {
	diff := queue.write - queue.read
	if diff < 0 {
		return len(queue.buff) + diff
	}
	return diff
}
