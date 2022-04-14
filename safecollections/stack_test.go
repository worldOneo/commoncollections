package safecollections

import "testing"

func fillStack(n int) *Stack[int] {
	stack := &Stack[int]{}
	for i := 0; i < n; i++ {
		stack = stack.Push(i)
	}
	return stack
} 

func TestStack_Push(t *testing.T) {
	stack := fillStack(200)
	if stack.Size() != 200 {
		t.Errorf("Stack.Push() got = %v, want %v", stack.Size(), 200)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := fillStack(200)
	var val int
	for i := 0; i < 200; i++ {
		stack, val = stack.Pop()
		if val != 199-i {
			t.Errorf("Stack.Pop() got = %v, want %v", val, 199-i)
		}
	}
	if stack.Size() != 0 {
		t.Errorf("Stack.Pop() got = %v, want %v", stack.Size(), 0)
	}
}

func TestStack_Peak(t *testing.T) {
	stack := fillStack(200)
	if stack.Peak() != 199 {
		t.Errorf("Stack.Peak() got = %v, want %v", stack.Peak(), 199)
	}
}

func TestStack_Get(t *testing.T) {
	stack := fillStack(200)
	for i := 0; i < 200; i++ {
		if stack.Get(i) != 199-i {
			t.Errorf("Stack.Get() got = %v, want %v", stack.Get(i), 199-i)
		}
	}
}

func TestStack_GetBoundsOver(t *testing.T) {
	stack := fillStack(200)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Stack.Get() did not panic")
		}
	}()
	stack.Get(201)
}

func TestStack_GetBoundsUnder(t *testing.T) {
	stack := fillStack(200)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Stack.Get() did not panic")
		}
	}()
	stack.Get(-1)
}