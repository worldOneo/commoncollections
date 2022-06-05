package commoncollections

import (
	"runtime"
	"testing"
)

func TestOptLock_RWLock(t *testing.T) {
	lock := NewOptLock()
	state, ok := lock.RLock()
	if !ok {
		t.Errorf("lock.RLock() returned false")
	}
	lock.Lock()
	if lock.RVerify(state) {
		t.Errorf("lock.RVerify(%d) returned true", state)
	}
	_, dead := lock.RLock()
	if dead {
		t.Errorf("lock.RLock() returned true")
	}
	lock.Unlock()
	if lock.RVerify(state) {
		t.Errorf("lock.RVerify() returned true")
	}
	state, ok = lock.RLock()
	if !ok {
		t.Errorf("lock.RLock() returned false")
	}
	if !lock.RVerify(state) {
		t.Errorf("lock.RVerify() returned false")
	}
}

func TestOptLock_DoubleUnlock(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	lock := NewOptLock()
	lock.Unlock()
}

func TestOptLock_Spin(t *testing.T) {
	lock := NewOptLock()
	lock.Lock()
	notify := make(chan struct{})
	go func() {
		notify <- struct{}{}
		lock.Lock()
		notify <- struct{}{}
	}()
	<-notify
	runtime.Gosched()
	lock.Unlock()
	<-notify
}
