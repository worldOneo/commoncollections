package commoncollections

import "fmt"

type entry[K comparable, V any] struct {
	key   K
	value V
}

type group[K comparable, V any] struct {
	fingerprintsJumps [4][2]byte // [][fingerprint, jump]
	entries           [4]entry[K, V]
}

const (
	mapPresent  byte = 0b1000_0000
	mapFPMask        = ^mapPresent
	mapMaxJumps      = 125
	mapNope          = ^uint64(0)
)

// Stolen with good faith
// https://github.com/skarupke/flat_hash_map/blob/master/bytell_hash_map.hpp
//
//	Copyright Malte Skarupke 2017.
//
// Distributed under the Boost Software License, Version 1.0.
//
//	(See http://www.boost.org/LICENSE_1_0.txt)
// var mapJumpDist = [126]uint64{
// 	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,

// 	21, 28, 36, 45, 55, 66, 78, 91, 105, 120, 136, 153, 171, 190, 210, 231,
// 	253, 276, 300, 325, 351, 378, 406, 435, 465, 496, 528, 561, 595, 630,
// 	666, 703, 741, 780, 820, 861, 903, 946, 990, 1035, 1081, 1128, 1176,
// 	1225, 1275, 1326, 1378, 1431, 1485, 1540, 1596, 1653, 1711, 1770, 1830,
// 	1891, 1953, 2016, 2080, 2145, 2211, 2278, 2346, 2415, 2485, 2556,

// 	3741, 8385, 18915, 42486, 95703, 215496, 485605, 1091503, 2456436,
// 	5529475, 12437578, 27986421, 62972253, 141700195, 318819126, 717314626,
// 	1614000520, 3631437253, 8170829695, 18384318876, 41364501751,
// 	93070021080, 209407709220, 471167588430, 1060127437995, 2385287281530,
// 	5366895564381, 12075513791265, 27169907873235, 61132301007778,
// 	137547673121001, 309482258302503, 696335090510256, 1566753939653640,
// 	3525196427195653, 7931691866727775, 17846306747368716,
// 	40154190394120111, 90346928493040500, 203280588949935750,
// 	457381324898247375, 1029107980662394500, 2315492957028380766,
// 	5209859150892887590,
// }

// Map is a generic any to any map
type Map[K comparable, V any] struct {
	hasher           func(K) uint64
	itemCount        uint64
	size             uint64
	maxItemCount     uint64
	groupLenMinusOne uint64
	groups           []group[K, V]
}

const groupMak = 0b11
const groupJumpShift = 2

// Get reads an item from the map and returns
// the value and true if it is present or
// the nil value and false otherwise
func (m *Map[K, V]) Get(k K) (V, bool) {
	hash := m.hasher(k)
	location := hash
	fp := byte(hash&0b1111_1111) & mapFPMask

	for {
		groupIndex := (location >> groupJumpShift) & m.groupLenMinusOne
		group := &m.groups[groupIndex]
		step := location & groupMak
		otherFp := group.fingerprintsJumps[step][0]
		otherJmp := group.fingerprintsJumps[step][1]
		jumpDistance := uint64(otherJmp) & uint64(mapFPMask)
		if otherFp&mapPresent == 0 && otherJmp&mapPresent == 0 {
			var v V
			return v, false
		}
		if otherFp&mapPresent == 0 {
			location = hash + jumpDistance
			continue
		}
		if otherFp&mapFPMask == fp {
			if group.entries[step].key == k {
				return group.entries[step].value, true
			}
		}
		if jumpDistance == 0 {
			var v V
			return v, false
		}
		fmt.Printf("Following steps: %v %v %v %v %v\n", k, location, groupIndex, step, jumpDistance)
		location += jumpDistance
	}
}

// Put inserts a value into the map
func (m *Map[K, V]) Put(k K, v V) {
	hash := m.hasher(k)
	fp := byte(hash&0b1111_1111) & mapFPMask
	location := hash
	prevGIndex := mapNope
	prevStep := mapNope
	for i := uint64(1); i < mapMaxJumps; i++ {
		groupIndex := (location >> groupJumpShift) & m.groupLenMinusOne
		step := location & groupMak
		group := &m.groups[groupIndex]
		otherFp := group.fingerprintsJumps[step][0]
		otherJmp := group.fingerprintsJumps[step][1]

		if otherFp&mapPresent == 0 {
			slot := &group.entries[step]
			slot.key = k
			slot.value = v
			head := &group.fingerprintsJumps[step]
			head[0] = fp | mapPresent
			if prevGIndex != mapNope {
				m.groups[prevGIndex].fingerprintsJumps[prevStep][1] = byte(i) | mapPresent
			}
			m.itemCount++
			if m.itemCount == m.maxItemCount {
				m.grow()
			}
			return
		}

		if otherFp == fp {
			if group.entries[step].key == k {
				group.entries[step].value = v
				return
			}
		}

		jumpDistance := uint64(otherJmp) & uint64(mapFPMask)
		if otherJmp&mapPresent == 0 || otherFp != fp {
			location++
		} else {
			location += jumpDistance
		}
		prevGIndex = groupIndex
		prevStep = step
	}
	m.grow()
	m.Put(k, v)
}

// Delete removes an item from the map and returns
// the value and true if it was present or
// the nil value and false otherwise.
func (m *Map[K, V]) Delete(k K) (V, bool) {
	hash := m.hasher(k)
	location := hash
	fp := byte(hash&0b1111_1111) & mapFPMask

	for {
		groupIndex := (location >> groupJumpShift) & m.groupLenMinusOne
		group := &m.groups[groupIndex]
		step := location & groupMak
		otherFp := group.fingerprintsJumps[step][0]
		otherJmp := group.fingerprintsJumps[step][1]
		jumpDistance := uint64(otherJmp) & uint64(mapFPMask)
		if otherFp&mapPresent == 0 && otherJmp&mapPresent == 0 {
			var v V
			return v, false
		}
		if otherFp&mapPresent == 0 {
			location = hash + jumpDistance
			continue
		}
		if otherFp&mapFPMask == fp {
			if group.entries[step].key == k {
				m.itemCount--
				group.fingerprintsJumps[step][0] = 0
				return group.entries[step].value, true
			}
		}
		if jumpDistance == 0 {
			var v V
			return v, false
		}
		location += jumpDistance
	}
}

func (m *Map[K, V]) grow() {
	oldData := m.groups
	oldSize := m.size
	m.size *= 2
	m.groups = make([]group[K, V], m.size)
	m.groupLenMinusOne = m.size - 1
	m.itemCount = 0
	m.maxItemCount *= 2
	for i := uint64(0); i < oldSize; i++ {
		group := &oldData[i]
		for j := uint64(0); j < 4; j++ {
			if group.fingerprintsJumps[j][0]&mapPresent != 0 {
				m.Put(group.entries[j].key, group.entries[j].value)
			}
		}
	}
}

// NewMap creates a new map :D
func NewMap[K comparable, V any](hasher func(K) uint64) Map[K, V] {
	return Map[K, V]{
		hasher:           hasher,
		itemCount:        0,
		size:             4,
		maxItemCount:     15, // 93.75% load factor
		groupLenMinusOne: 3,
		groups:           make([]group[K, V], 4),
	}
}

func crc8(data uint64) byte {
	crc := byte(0xff)
	var i, j uint64
	for i = 0; i < 8; i++ {
		crc ^= byte(data >> (8 * i))
		for j = 0; j < 8; j++ {
			if (crc & 0x80) != 0 {
				crc = (byte)((crc << 1) ^ 0x31)
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}
