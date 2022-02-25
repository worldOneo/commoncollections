package commoncollections

// Set type using builtin map as container.
// A set contains every item at most once.
type Set[T comparable] map[T]struct{}

// NewSet creates a new set
func NewSet[T comparable]() Set[T] {
	return Set[T](make(map[T]struct{}))
}

// Insert adds the item elem to the set.
func (S Set[T]) Insert(elem T) {
	S[elem] = struct{}{}
}

// Contains returns true if the set contains elem and false otherwise.
func (S Set[T]) Contains(elem T) bool {
	_, ok := S[elem]
	return ok
}

// Remove removes the given element from the map
func (S Set[T]) Remove(elem T) {
	delete(S, elem)
}

// Values returns all values in the set
// It returns a new slice with a copy of the elements in it
func (S Set[T]) Values() []T {
	keys := make([]T, 0, len(S))
	for key := range S {
		keys = append(keys, key)
	}
	return keys
}
