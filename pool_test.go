package commoncollections

import "testing"

func testPool(t *testing.T, pool *Pool[int]) {
	for i := 0; i < 33; i++ {
		pool.Get()
	}
	for i := 0; i < 33; i++ {
		pool.Put(i)
	}
	for i := 0; i < 33; i++ {
		pool.Get()
	}
}

func TestFreePool(t *testing.T) {
	c := 0
	factory := func() int {
		if c > 33 {
			panic("To many items created")
		}
		val := c
		c++
		return val
	}
	pool := NewPool(factory)
	testPool(t, &pool)
}

func TestSyncPool(t *testing.T) {
	c := 0
	factory := func() int {
		if c > 33 {
			panic("To many items created")
		}
		val := c
		c++
		return val
	}
	pool := NewSyncPool(factory)
	testPool(t, &pool)
}