package structure

//ISinglyCell 单链表单元interface
type ISinglyCell interface {
	Data() interface{}
	SetData(interface{})
	Next() ISinglyCell
	SetNext(ISinglyCell)
}

// ISinglyLinkedList 单链表interface
type ISinglyLinkedList interface {
	Range(func(ISinglyCell) (breaked, removed bool))
	Append(...interface{})
	Clear()
	Head() ISinglyCell
}

// NewSinglyCell 新建一个SinglyCell
func NewSinglyCell(v interface{}) ISinglyCell {
	return &SinglyCell{
		data: v,
	}
}

// NewSinglyLinkedList 创建SinglyLinkedList
func NewSinglyLinkedList(f func(interface{}) ISinglyCell) ISinglyLinkedList {
	return &SinglyLinkedList{
		newCell: f,
	}
}

// SinglyCell 单链表单元
type SinglyCell struct {
	next ISinglyCell
	data interface{}
}

//Next 下一个元素
func (s *SinglyCell) Next() ISinglyCell {
	return s.next
}

//SetNext 设置下一个元素
func (s *SinglyCell) SetNext(next ISinglyCell) {
	if s != nil {
		s.next = next
	}
}

//Data 返回数据
func (s *SinglyCell) Data() interface{} {
	return s.data
}

//SetData 设置数据
func (s *SinglyCell) SetData(v interface{}) {
	s.data = v
}

//SinglyLinkedList 单链表
type SinglyLinkedList struct {
	head, tail ISinglyCell
	newCell    func(interface{}) ISinglyCell
}

//Head 单链表头元素
func (s *SinglyLinkedList) Head() ISinglyCell {
	return s.head
}

//Range 遍历链表
func (s *SinglyLinkedList) Range(f func(ISinglyCell) (breaked, removed bool)) {
	var (
		cursor, before ISinglyCell
	)
	for cursor, before = s.head, nil; cursor != nil; cursor, before = cursor.Next(), cursor {
		breaked, removed := f(cursor)
		if removed {
			if cursor.Next() == nil {
				s.tail = before
				before.SetNext(nil)
				break
			} else {
				before.SetNext(cursor.Next())
				cursor.SetNext(cursor.Next())
			}
		}
		if breaked {
			break
		}
	}
}

//Append 添加元素
func (s *SinglyLinkedList) Append(data ...interface{}) {
	for _, cell := range data {
		if s.head == nil {
			s.head = s.newCell(cell)
			s.tail = s.head
		} else {
			s.tail.SetNext(s.newCell(cell))
			s.tail = s.tail.Next()
		}
	}
}

//Clear 清空链表
func (s *SinglyLinkedList) Clear() {
	s.head, s.tail = nil, nil
}
