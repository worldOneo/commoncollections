package commoncollections

// BKTree implementation.
type BKTree[T comparable] struct {
	dist     func(T, T) int
	children map[int]*BKTree[T]
	value    T
	full     bool
}

// NewBKTree creates a new BKTree.
// The dist function returns the distance between two elements
// of type T.
func NewBKTree[T comparable](dist func(T, T) int) BKTree[T] {
	return BKTree[T]{
		dist:     dist,
		children: make(map[int]*BKTree[T]),
	}
}

// Insert adds a new element to the tree.
// Returns if the element was already present.
func (b *BKTree[T]) Insert(val T) bool {
	if !b.full {
		b.value = val
		b.full = true
		return false
	}

	if b.value == val {
		return true
	}

	dist := Abs(b.dist(b.value, val))
	tree, ok := b.children[dist]
	if !ok {
		newTree := NewBKTree(b.dist)
		tree = &newTree
		b.children[dist] = &newTree
	}
	return tree.Insert(val)
}

// Find returns every element of the tree that are fewer than
// maxDist avay from val.
// If nothing is found, an empty slice is returned.
func (b *BKTree[T]) Find(val T, maxDist int) []T {
	if !b.full {
		return []T{}
	}
	var results []T
	dist := Abs(b.dist(b.value, val))
	if dist <= maxDist {
		results = append(results, b.value)
	}
	for i := dist - maxDist; i <= dist+maxDist; i++ {
		tree := b.children[i]
		if tree != nil {
			results = append(results, tree.Find(val, maxDist)...)
		}
	}
	return results
}
