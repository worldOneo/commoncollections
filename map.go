package commoncollections

const (
	mapPresent    = 1 << 63
	mapMask       = mapPresent - 1
	mapMaxJumps   = 32
	mapBucketSize = 4
)

type entry[K, V any] struct {
	distance uint64
	hash     uint64
	key      K
	value    V
}

// Map is a generic any to any map
type Map[K comparable, V any] struct {
	hasher       func(K) uint64
	size         uint64
	sizeMinusOne uint64
	itemCount    uint64
	maxItemCount uint64
	items        []entry[K, V]
}

// Get reads an item from the map and returns
// the value and true if it is present or
// the nil value and false otherwise
func (m *Map[K, V]) Get(k K) (V, bool) {
	hash := scramble(m.hasher(k))
	slot := &m.items[hash&m.sizeMinusOne]
	present := slot.distance&mapPresent != 0
	offset := uint64(0)
	for present {
		if slot.hash == hash {
			if slot.key == k {
				return slot.value, true
			}
		}
		offset++
		slot = &m.items[(hash+offset)&m.sizeMinusOne]
		present = slot.distance&mapPresent != 0
	}
	var v V
	return v, false
}

// Put inserts a value into the map
// retuns the previous value and true if it was present
// false otherwise
func (m *Map[K, V]) Put(k K, v V) (V, bool) {
	hash := scramble(m.hasher(k))
	for jump := uint64(0); jump < mapMaxJumps; jump++ {
		slot := &m.items[(hash+jump)&m.sizeMinusOne]
		present := slot.distance&mapPresent != 0
		distance := slot.distance & mapMask
		if !present {
			slot.distance = jump | mapPresent
			slot.hash = hash
			slot.key = k
			slot.value = v
			m.itemCount++
			if m.itemCount == m.maxItemCount {
				m.grow()
			}
			var v V
			return v, false
		}
		if jump > distance {
			ph, pk, pv := slot.hash, slot.key, slot.value
			slot.distance = jump | mapPresent
			slot.hash = hash
			slot.key = k
			slot.value = v
			offset := jump + 1
			for {
				slot := &m.items[(hash+offset)&m.sizeMinusOne]
				present := slot.distance&mapPresent != 0
				tph, tpk, tpv := slot.hash, slot.key, slot.value
				tdistance := slot.distance & mapMask
				slot.distance = (distance + 1) | mapPresent
				slot.hash = ph
				slot.key = pk
				slot.value = pv
				if !present {
					break
				}
				ph, pk, pv = tph, tpk, tpv
				distance = tdistance
				offset++
			}
			m.itemCount++
			if m.itemCount == m.maxItemCount {
				m.grow()
			}
			var v V
			return v, false
		}
		if slot.hash == hash {
			if slot.key == k {
				ov := slot.value
				slot.value = v
				return ov, true
			}
		}
	}
	m.grow()
	return m.Put(k, v)
}

// Delete removes an item from the map and returns
// the value and true if it was present or
// the nil value and false otherwise.
func (m *Map[K, V]) Delete(k K) (V, bool) {
	hash := scramble(m.hasher(k))
	offset := uint64(0)
	for {
		slot := &m.items[(hash+offset)&m.sizeMinusOne]
		present := slot.distance&mapPresent != 0
		if !present {
			break
		}
		if slot.hash == hash && slot.key == k {
			v := slot.value
			prev := (hash + offset) & m.sizeMinusOne
			for {
				offset++
				slot = &m.items[(hash+offset)&m.sizeMinusOne]
				present = slot.distance&mapPresent != 0
				distance := slot.distance & mapMask
				if !present || distance == 0 {
					break
				}
				m.items[prev].distance = (prev - (slot.hash & m.sizeMinusOne)) | mapPresent
				m.items[prev].hash = slot.hash
				m.items[prev].key = slot.key
				m.items[prev].value = slot.value
				prev = (hash + offset) & m.sizeMinusOne
			}
			m.items[prev].distance = 0
			return v, true
		}
		offset++
	}
	var v V
	return v, false
}

func (m *Map[K, V]) grow() {
	oldData := m.items
	oldSize := m.size
	m.size *= 2
	m.sizeMinusOne = m.size - 1
	m.items = make([]entry[K, V], m.size)
	m.itemCount = 0
	m.maxItemCount *= 2
	for i := uint64(0); i < oldSize; i++ {
		bucket := &oldData[i]
		if bucket.distance&mapPresent != 0 {
			m.Put(bucket.key, bucket.value)
		}
	}
}

// NewMap creates a new map :D
func NewMap[K comparable, V any](hasher func(K) uint64) Map[K, V] {
	return Map[K, V]{
		hasher:       hasher,
		itemCount:    0,
		size:         16,
		maxItemCount: 14, // 75% load factor
		sizeMinusOne: 15,
		items:        make([]entry[K, V], 16),
	}
}
