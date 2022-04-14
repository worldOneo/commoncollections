package safecollections

// Stack is an immutable stack.
// The null stack is an empty stack and usable.
type Stack[T any] struct {
	tail *Stack[T]
	val  T
	size int
}


// Push returns a new stack with the value pushed onto the top.
func (s *Stack[T]) Push(val T) *Stack[T] {
	return &Stack[T]{s, val, s.size + 1}
}

// Pop returns a new stack with the top value popped off and the popped value.
// Popping from an empty stack returns nil and null value of T.
func (s *Stack[T]) Pop() (*Stack[T], T) {
	return s.tail, s.val
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return s.size
}

// Peak returns the top value in the stack.
func (s *Stack[T]) Peak() T {
	return s.val
}

// Get returns the nth value in the stack.
// Panics if n is out of bounds.
func (s *Stack[T]) Get(i int) T {
	if i < 0 || i >= s.size {
		panic("index out of range")
	}
	if i == 0 {
		return s.val
	}
	return s.tail.Get(i - 1)
}