package commoncollections

// IntMap is an int to any map which might provide performance
// benefits over the std map.
// Insert and lookup speeds are often slightly faster.
// Also delete returns the value it held avoiding aditional lookups
// in some situations.
type IntMap[V any] struct {
	keys     []uint64
	values   []V
	keymask  uint64
	hasFree  bool
	free     V
	cap      int
	size     int
	nilvalue V
}

func scramble(k uint64) uint64 {
	hash := k * 0x9E3779B9
	return hash * (hash >> 16)
}

const (
	initMapSize = 16
	freeKey     = 0
)

// NewIntMap initialises a new intmap
func NewIntMap[V any](nilvalue V) *IntMap[V] {
	return &IntMap[V]{
		keys:     make([]uint64, initMapSize),
		values:   make([]V, initMapSize),
		keymask:  initMapSize - 1,
		hasFree:  false,
		cap:      (initMapSize / 6) * 7,
		size:     0,
		nilvalue: nilvalue,
	}
}

func (im *IntMap[V]) index(k uint64) uint64 {
	return scramble(k) & im.keymask
}

func (im *IntMap[V]) next(k uint64) uint64 {
	return (k + 1) & im.keymask
}

// Put adds an item to the int map
func (im *IntMap[V]) Put(key uint64, val V) {
	if key == freeKey {
		im.free = val
		im.hasFree = true
		return
	}
	index := im.index(key)
	for {
		definedKey := im.keys[index]
		if key == definedKey || definedKey == freeKey {
			if definedKey == freeKey {
				im.size++
				im.keys[index] = key
			}
			im.values[index] = val
			break
		}
		index = im.next(index)
	}
	im.expand()
}

// Get retrieves an item from the intmap and returns value, true or
// 0, false if the item isn't in this map.
func (im *IntMap[V]) Get(key uint64) (V, bool) {
	if key == freeKey {
		if im.hasFree {
			return im.free, true
		}
		return im.nilvalue, false
	}
	index := im.index(key)
	for {
		definedKey := im.keys[index]
		if definedKey == freeKey {
			return im.nilvalue, false
		}
		if key == definedKey {
			return im.values[index], true
		}
		index = im.next(index)
	}
}

// Delete removes a value from this map returns value,true or
// 0, false if the key wasnt in this map
func (im *IntMap[V]) Delete(key uint64) (V, bool) {
	if key == freeKey {
		if im.hasFree {
			im.hasFree = false
			return im.free, true
		}
		return im.nilvalue, false
	}
	index := im.index(key)
	for {
		definedKey := im.keys[index]
		if definedKey == freeKey {
			return im.nilvalue, false
		}
		if key == definedKey {
			data := im.values[index]
			im.unshift(index)
			im.size--
			return data, true
		}
		index = im.next(index)
	}
}

func (im *IntMap[V]) unshift(current uint64) {
	var key uint64
	for {
		last := current
		current = im.next(current)
		for {
			key = im.keys[current]
			if key == freeKey {
				im.keys[last] = freeKey
				return
			}
			slot := im.index(key)
			if last <= current {
				if last >= slot || slot > current {
					break
				}
			} else if last >= slot && slot > current {
				break
			}
			current = im.next(current)
		}
		im.keys[last] = key
		im.values[last] = im.values[current]
	}
}

func (im *IntMap[V]) expand() {
	if im.size < im.cap {
		return
	}

	oldLen := uint64(len(im.keys))
	oldKeys := im.keys
	oldVal := im.values
	im.cap *= 2
	l := oldLen * 2
	im.keymask = l - 1
	im.size = 0
	im.keys = make([]uint64, l)
	im.values = make([]V, l)
	for i := uint64(0); i < oldLen; i++ {
		im.Put(oldKeys[i], oldVal[i])
	}
}
