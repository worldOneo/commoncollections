package commoncollections

type binaryNode[T any] struct {
	left, right *binaryNode[T]
	value       T
	present     bool
	weight      int8
}

type BinaryTree[T any] struct {
	comparator func(a, b T) int
	root       binaryNode[T]
}

func NewBinaryTree[T any](comparator func(a, b T) int) BinaryTree[T] {
	return BinaryTree[T]{
		comparator: comparator,
		root:       binaryNode[T]{},
	}
}

func (b *binaryNode[T]) Insert(v T) {

}