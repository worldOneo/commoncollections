package commoncollections

import "sync"

// Pool stores a set of items.
// It can returned a pooled or a new one when an item is retrieved.
// A pool might be synced for parallel access see NewSyncPool.
type Pool[T any] struct {
	queue   *Queue[T]
	lock    sync.Locker
	factory func() T
}

func newPool[T any](factory func() T, lock sync.Locker) Pool[T] {
	value := factory()
	queue := NewQueue[T]()
	queue.Push(value)
	return Pool[T]{
		queue:   &queue,
		lock:    lock,
		factory: factory,
	}
}

// NewSyncPool creates a new Pool which is safe to access
// from multiple goroutines simultaneously.
//
// The Pool requires locking for this and might therefore
// be slower for
func NewSyncPool[T any](factory func() T) Pool[T] {
	return newPool(factory, NewSpinLock())
}

// NewPool creates a new Pool which is not safe to access
// from multiple goroutines.
// For a safe implementation us NewSyncPool
func NewPool[T any](factory func() T) Pool[T] {
	return newPool(factory, NewNoLock())
}

// Get returns an item from the pool if any is available
// or creates a new one.
func (pool *Pool[T]) Get() T {
	pool.lock.Lock()
	val, ok := pool.queue.Pop()
	pool.lock.Unlock()
	if !ok {
		return pool.factory()
	}
	return val
}

// Put adds a the given element to the Pool.
func (pool *Pool[T]) Put(elem T) {
	pool.lock.Lock()
	pool.queue.Push(elem)
	pool.lock.Unlock()
}
