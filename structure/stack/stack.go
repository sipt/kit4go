package stack

import (
	"github.com/sipt/algorithm/structure/list"
	"github.com/sipt/algorithm/util"
)

//IStack 栈接口
type IStack interface {
	Pop() interface{}
	Push(interface{})
	IsEmpty() bool
	Print()
	Clear()
}

//NewStack 创建Stack
func NewStack() IStack {
	return &Stack{
		linkedList: list.NewDoubleLinkedList(list.NewDoubleCell),
	}
}

//Stack 栈
type Stack struct {
	linkedList list.IDoubleLinkedList
}

//Pop pop
func (s *Stack) Pop() interface{} {
	cell := s.linkedList.PopTail()
	if util.IsNil(cell) {
		return nil
	}
	return cell.Data()
}

//Push push
func (s *Stack) Push(v interface{}) {
	s.linkedList.Append(v)
}

//Clear clear
func (s *Stack) Clear() {
	s.linkedList.Clear()
}

//IsEmpty
func (s *Stack) IsEmpty() bool {
	return s.linkedList.IsEmpty()
}

func (s *Stack) Print() {
	s.linkedList.Print()
}
