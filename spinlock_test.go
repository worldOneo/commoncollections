package commoncollections

import (
	"testing"
)

func TestSpinLockLock(t *testing.T) {
	lock := SpinLock(0)
	lock.Lock()
	lock.Unlock()
}

func TestSpinLockUnlock(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("unlocked spinlock unlocked")
		}
	}()
	lock := SpinLock(0)
	lock.Unlock()
}
