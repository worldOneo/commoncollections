package commoncollections

// Queue type for a FIFO queue based on a circular buffer
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
func (Q *Queue[T]) Push(elem T) {
	Q.buff[Q.write] = elem
	Q.write++
	cap := len(Q.buff)
	Q.write &= cap - 1 // cap power 2 means Q.write %= cap
	if Q.write == Q.read {
		old := Q.buff
		Q.buff = make([]T, cap*2)
		copy(Q.buff, old[:Q.read])
		copy(Q.buff[Q.write+cap:], old[Q.read:])
		Q.read += cap
	}
}

// Pop reads the last element from the queue removing it.
// It returns the element and true if an element is present
// and returns the nilvalue and false if no element is left.
func (Q *Queue[T]) Pop() (T, bool) {
	if Q.read == Q.write {
		return Q.nilvalue, false
	}
	val := Q.buff[Q.read]
	Q.read++
	if Q.read >= len(Q.buff) {
		Q.read = 0
	}
	return val, true
}

// Peek reads the last element from the queue without changing
// the queue.
// Returns the element and true if an element is present
// and returns the nilvalue and false if no element is left.
func (Q *Queue[T]) Peek() (T, bool) {
	if Q.read == Q.write {
		return Q.nilvalue, false
	}
	return Q.buff[Q.read], true
}

func (Q *Queue[T]) Len() int {
	diff := Q.write - Q.read
	if diff < 0 {
		return len(Q.buff) + diff
	}
	return diff
}
