package stack

func ReverseStack(stack IStack) {
	if stack.IsEmpty() {
		return
	}
	temp := stack.Pop()
	ReverseStack(stack)
	if stack.IsEmpty() {
		stack.Push(temp)
		return
	}
	temp2 := stack.Pop()
	ReverseStack(stack)
	stack.Push(temp)
	ReverseStack(stack)
	stack.Push(temp2)
}
