package commoncollections

import "testing"

func botherRLock(lock *RWSpinLock, ch chan bool, i int) {
	for j := 0; j < i; j++ {
		lock.RLock()
		lock.RUnlock()
	}
	ch <- true
}

func botherLock(lock *RWSpinLock, ch chan bool, i int) {
	for j := 0; j < i; j++ {
		lock.Lock()
		lock.Unlock()
	}
	ch <- true
}

func TestRWLock_Lock(t *testing.T) {
	lock := NewRWSpinLock()
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go botherLock(&lock, ch, 1000)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}

func TestRWLock_RLock(t *testing.T) {
	lock := NewRWSpinLock()
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go botherRLock(&lock, ch, 1000)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}

func TestRWLock(t *testing.T) {
	lock := NewRWSpinLock()
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go botherRLock(&lock, ch, 1000)
	}
	for i := 0; i < 10; i++ {
		go botherLock(&lock, ch, 1000)
	}
	for i := 0; i < 20; i++ {
		<-ch
	}
}
