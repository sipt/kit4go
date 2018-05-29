package stack

import (
	"testing"
)

func TestReverseStack(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	ReverseStack(stack)
	var v int
	for i := 0; i < 10; i++ {
		v = stack.Pop().(int)
		if v != i {
			t.Errorf("value not equal: %d != %d", v, i)
			return
		}
	}
}
