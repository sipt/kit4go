package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sipt/algorithm/structure/stack"
	"github.com/sipt/algorithm/util"
)

func TestStack(t *testing.T) {
	stack := stack.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.Equal(t, stack.Pop(), 3)
	assert.Equal(t, stack.Pop(), 2)
	stack.Push(4)
	assert.Equal(t, stack.Pop(), 4)
	assert.Equal(t, stack.Pop(), 1)
	assert.Equal(t, util.IsNil(stack.Pop()), true)
}
