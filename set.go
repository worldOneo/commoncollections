package commoncollections

// Set type using builtin map as container.
// A set contains every item at most once.
type Set[T comparable] map[T]struct{}

// NewSet creates a new set
func NewSet[T comparable]() Set[T] {
	return Set[T](make(map[T]struct{}))
}

// Insert adds the item elem to the set.
func (set Set[T]) Insert(elem T) {
	set[elem] = struct{}{}
}

// Contains returns true if the set contains elem and false otherwise.
func (set Set[T]) Contains(elem T) bool {
	_, ok := set[elem]
	return ok
}

// Remove removes the given element from the map
func (set Set[T]) Remove(elem T) {
	delete(set, elem)
}

// Values returns all values in the set
// It returns a new slice with a copy of the elements in it
func (set Set[T]) Values() []T {
	keys := make([]T, 0, len(set))
	for key := range set {
		keys = append(keys, key)
	}
	return keys
}

// Len returs the size of the set
func (set Set[T]) Len() int {
	return len(set)
}
