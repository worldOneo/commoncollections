package commoncollections

type Queue[T any] struct {
	nilvalue T
	buff     []T
	read     int
	write    int
}

func NewQueue[T any](nilvalue T) *Queue[T] {
	return &Queue[T]{
		nilvalue: nilvalue,
		buff:     make([]T, 16),
		read:     0,
		write:    0,
	}
}

func (Q *Queue[T]) Push(elem T) {
	Q.buff[Q.write] = elem
	Q.write++
	cap := len(Q.buff)
	if Q.write >= cap {
		Q.write = 0
	}
	if Q.write == Q.read {
		old := Q.buff
		Q.buff = make([]T, cap*2)
		copy(Q.buff, old[:Q.read])
		copy(Q.buff[Q.write+cap:], old[Q.read:])
		Q.read += cap
	}
}

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
