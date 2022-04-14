package safecollections

const keyPartLen = 5
const bucketSize = 1 << keyPartLen

type mapValue[T any] struct {
	key uint64
	val T
}

// Map is an immutable map. The null map is an empty map and usable.
// Put and Delete operations are significant cheaper in this Map
// than in a copy-on-write map but much more expensive than in a mutable map.
// Its lookup speed is O(log32 n)
type Map[T any] struct {
	bucket [bucketSize]interface{}
}

// Put inserts a new value into the map.
// If the key already exists, the old value is overwritten.
// Returns the new map, the old value, and a boolean
// indicating whether the key was present.
func (m *Map[T]) Put(key uint64, val T) (*Map[T], T, bool) {
	idx := key % bucketSize
	var self [bucketSize]interface{}
	copy(self[:], m.bucket[:])
	bucket := self[idx]
	switch bucket.(type) {
	case *Map[T]:
		n, val, ok := bucket.(*Map[T]).Put(key>>keyPartLen, val)
		self[idx] = n
		return &Map[T]{self}, val, ok
	case *mapValue[T]:
		if bucket.(*mapValue[T]).key == key {
			self[idx] = &mapValue[T]{key, val}
			return &Map[T]{self}, bucket.(*mapValue[T]).val, true
		} else {
			n := &Map[T]{}
			n, _, _ = n.Put(bucket.(*mapValue[T]).key>>5, bucket.(*mapValue[T]).val)
			n, val, ok := n.Put(key>>keyPartLen, val)
			self[idx] = n
			return &Map[T]{self}, val, ok
		}
	}
	self[idx] = &mapValue[T]{key, val}
	return &Map[T]{self}, *new(T), false
}

func (s *Map[T]) empty() bool {
	for i := uint8(0); i < bucketSize; i++ {
		if s.bucket[i] != nil {
			return false
		}
	}
	return true
}

// Delete removes a value from the map.
// Returns the new map, the old value, and a boolean
// indicating whether the key was present.
func (m *Map[T]) Delete(key uint64) (*Map[T], T, bool) {
	idx := key % bucketSize
	bucket := m.bucket[idx]
	switch bucket.(type) {
	case *Map[T]:
		old := bucket.(*Map[T])
		n, val, ok := old.Delete(key >> keyPartLen)
		if !ok {
			return m, *new(T), false
		}
		var self [bucketSize]interface{}
		copy(self[:], m.bucket[:])
		if n.empty() {
			self[idx] = nil
		} else {
			self[idx] = n
		}
		return &Map[T]{self}, val, ok
	case *mapValue[T]:
		if bucket.(*mapValue[T]).key == key {
			val := bucket.(*mapValue[T]).val
			var self [bucketSize]interface{}
			copy(self[:], m.bucket[:])
			self[idx] = nil
			return &Map[T]{self}, val, true
		}
	}
	return m, *new(T), false
}

// Get returns the value associated with the given key.
// Returns the value and a boolean indicating
// whether the key was present.
func (m *Map[T]) Get(key uint64) (T, bool) {
	bucket := m.bucket[key%bucketSize]
	switch bucket.(type) {
	case *Map[T]:
		return bucket.(*Map[T]).Get(key >> keyPartLen)
	case *mapValue[T]:
		value := bucket.(*mapValue[T])
		if value.key == key {
			return value.val, true
		}
		return *new(T), false
	default:
		return *new(T), false
	}
}
