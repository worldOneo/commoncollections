package commoncollections

import (
	"sync/atomic"
)

type channelNode[T any] struct {
	next atomic.Pointer[channelNode[T]]
	val  T
}

// Channel provides a unbound channel.
// To be unbound is the main benefit over std chan but
// performance gains might be possible in high contention situations.
// The nil value is not usable.
type Channel[T any] struct {
	head atomic.Pointer[channelNode[T]]
	tail atomic.Pointer[channelNode[T]]
}

// NewChannel creates a new channel
func NewChannel[T any]() Channel[T] {
	base := atomic.Pointer[channelNode[T]]{}
	base.Store(&channelNode[T]{})
	return Channel[T]{
		head: base,
		tail: base,
	}
}

// Send sends a message through the channel.
// This operation might block if the channel is bussy.
func (c *Channel[T]) Send(val T) {
	q := &channelNode[T]{
		val: val,
	}
	for {
		p := c.tail.Load()
		if !p.next.CompareAndSwap(nil, q) {
			c.tail.CompareAndSwap(p, p.next.Load())
		} else {
			c.tail.CompareAndSwap(p, q)
			break
		}
	}
}

// TryRecv reads a value from the channel.
//
// Returns value, true if successful, nil, false otherwise.
func (c *Channel[T]) TryRecv() (T, bool) {
	var t T
	for {
		p := c.head.Load()
		next := p.next.Load()
		if next == nil {
			return t, false
		}
		if c.head.CompareAndSwap(p, next) {
			return next.val, true
		}
	}
}

// Recv receives a message from the channel or blocks
// until one is available.
func (c *Channel[T]) Recv() T {
	val, recieved := c.TryRecv()
	backoff := 0
	for !recieved {
		backoff = spin(backoff)
		val, recieved = c.TryRecv()
	}
	return val
}
