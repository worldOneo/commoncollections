package commoncollections

import (
	"testing"
)

func TestMap_Put(t *testing.T) {
	n := uint64(10_000_000)
	filledMap(n)
}

func TestMap_Get(t *testing.T) {
	n := uint64(10_000_000)
	m := filledMap(n)
	for _, i := range randomInts(n) {
		v, ok := m.Get(i)
		if v != uint64(i) || !ok {
			t.Errorf("Map.Get() got = %v,%v, want %v,%v", v, ok, i, true)
		}
	}
	v, ok := m.Get(n + 1)
	if v != 0 || ok {
		t.Errorf("Map.Get() got = %v,%v, want %v,%v", v, ok, 0, true)
	}
}

func TestMap_Collission(t *testing.T) {
	m := NewMap[uint64, uint64](func(u uint64) uint64 { return u })
	collissions := []uint64{}
	prev := 0
	for i := uint64(0); i < 64; i++ {
		prev <<= 1
		prev++
		collissions = append(collissions, uint64(prev))
	}
	for _, i := range collissions {
		m.Put(i, i)
		m.Get(i)
	}
}

func TestMap_Delete(t *testing.T) {
	n := uint64(10_000_000)
	m := filledMap(n)
	for i := uint64(0); i < n-1; i++ {
		v, ok := m.Delete(i)
		if v != uint64(i) || !ok {
			t.Errorf("Map.Delete() got = %v,%v, want %v,%v", v, ok, i, true)
		}
	}
	for i := uint64(0); i < n-1; i++ {
		v, ok := m.Delete(i)
		if v != uint64(0) || ok {
			t.Errorf("Map.Delete() got = %v,%v, want %v,%v", v, ok, 0, false)
		}
	}
	v, ok := m.Delete(n + 1)
	if v != 0 || ok {
		t.Errorf("Map.Delete() got = %v,%v, want %v,%v", v, ok, 0, false)
	}
	for i := uint64(0); i < n-1; i++ {
		_, ok := m.Get(i)
		if ok {
			t.Errorf("Map.Delete() got = %v, want %v", ok, false)
		}
	}
	v, ok = m.Get(n - 1)
	if v != n-1 || !ok {
		t.Errorf("Map.Delete() got = %v, %v, want %v, %v", v, ok, n-1, true)
	}
}
