package safecollections

import "testing"

func filled(n uint64) *Map[uint64] {
	m := &Map[uint64]{}
	for i := uint64(0); i < n; i++ {
		m, _, _ = m.Put(i, i)
	}
	return m
}

func TestMap_Put(t *testing.T) {
	m := filled(200)
	m, val, ok := m.Put(100, 666)
	if val != 100 || !ok {
		t.Errorf("Map.Put() got = %v,%v, want %v,%v", val, ok, 100, true)
	}
	m, val, ok = m.Put(100, 777)
	if val != 666 || !ok {
		t.Errorf("Map.Put() got = %v,%v, want %v,%v", val, ok, 666, true)
	}
	m, val, ok = m.Put(200, 888)
	if val != 0 || ok {
		t.Errorf("Map.Put() got = %v,%v, want %v,%v", val, ok, 0, false)
	}
}

func TestMap_Get(t *testing.T) {
	m := filled(200)
	for i := uint64(0); i < 200; i++ {
		v, ok := m.Get(i)
		if v != uint64(i) || !ok {
			t.Errorf("Map.Get() got = %v,%v, want %v,%v", v, ok, i, true)
		}
	}
	v, ok := m.Get(201)
	if v != 0 || ok {
		t.Errorf("Map.Get() got = %v,%v, want %v,%v", v, ok, 0, true)
	}
	v, ok = m.Get(1 << 16)
	if v != 0 || ok {
		t.Errorf("Map.Get() got = %v,%v, want %v,%v", v, ok, 0, true)
	}
}

func TestMap_Delete(t *testing.T) {
	n := uint64(200)
	m := filled(n)
	var v uint64
	var ok bool
	m, v, ok = m.Delete(1 << 16)
	if v != 0 || ok {
		t.Errorf("Map.Delete() got = %v,%v, want %v,%v", v, ok, 0, false)
	}
	for i := uint64(0); i < n; i++ {
		m, v, ok = m.Delete(i)
		if v != uint64(i) || !ok {
			t.Errorf("Map.Delete() got = %v,%v, want %v,%v", v, ok, i, true)
		}
	}
	for i := uint64(0); i < n; i++ {
		m, v, ok = m.Delete(i)
		if v != uint64(0) || ok {
			t.Errorf("Map.Delete() got = %v,%v, want %v,%v", v, ok, 0, false)
		}
	}
}
