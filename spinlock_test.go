package commoncollections

import (
	"runtime"
	"testing"
)

func TestSpinLockLock(t *testing.T) {
	lock := SpinLock(0)
	lock.Lock()
	notify := make(chan struct{})
	go func() {
		notify <- struct{}{}
		lock.Lock()
		notify <- struct{}{}
	}();
	<-notify
	runtime.Gosched()
	lock.Unlock()
	<-notify
}

func TestSpinLockUnlock(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("unlocked spinlock unlocked")
		}
	}()
	lock := SpinLock(0)
	lock.Unlock()
	t.Fatalf("unlocked spinlock unlocked")
}
