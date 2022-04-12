package commoncollections

type Stack[T any] struct {
	buff []T
	ptr  int
}
