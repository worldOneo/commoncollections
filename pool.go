package commoncollections

import "sync"

// Pool stores a set of items.
// It can returned a pooled or a new one when an item is retrieved.
// A pool might be synced for parallel access see NewSyncPool.
type Pool[T any] struct {
	queue   *Queue[T]
	sync    bool
	lock    sync.Mutex
	factory func() T
}

func newPool[T any](factory func() T, isSync bool) *Pool[T] {
	value := factory()
	queue := NewQueue(value)
	queue.Push(value)
	return &Pool[T]{
		queue:   queue,
		sync:    isSync,
		lock:    sync.Mutex{},
		factory: factory,
	}
}

// NewSyncPool creates a new Pool which is safe to access
// from multiple goroutines simultaneously.
//
// The Pool requires locking for this and might therefore
// be slower for
func NewSyncPool[T any](factory func() T) *Pool[T] {
	return newPool(factory, true)
}

// NewPool creates a new Pool which is not safe to access
// from multiple goroutines.
// For a safe implementation us NewSyncPool
func NewPool[T any](factory func() T) *Pool[T] {
	return newPool(factory, false)
}

// Get returns an item from the pool if any is available
// or creates a new one.
func (P *Pool[T]) Get() T {
	if P.sync {
		P.lock.Lock()
		defer P.lock.Unlock()
	}
	val, ok := P.queue.Pop()
	if !ok {
		return P.factory()
	}
	return val
}

// Put adds a the given element to the Pool.
func (P *Pool[T]) Put(elem T) {
	if P.sync {
		P.lock.Lock()
		defer P.lock.Unlock()
	}
	P.queue.Push(elem)
}
