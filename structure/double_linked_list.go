package structure

import (
	"reflect"

	"github.com/sipt/algorithm/util"
)

// IDoubleCell 双链表单元接口
type IDoubleCell interface {
	Data() interface{}
	SetData(interface{})
	Next() IDoubleCell
	SetNext(IDoubleCell)
	Prev() IDoubleCell
	SetPrev(IDoubleCell)
}

// IDoubleLinkedList 双链表接口
type IDoubleLinkedList interface {
	Append(...interface{})
	Range(func(IDoubleCell) (breaked, removed bool))
	Clear()
	Head() IDoubleCell
	Tail() IDoubleCell
}

//NewDoubleCell 创建双链表单元
func NewDoubleCell(v interface{}) IDoubleCell {
	return &DoubleCell{
		data: v,
	}
}

//NewDoubleLinkedList 创建双链表
func NewDoubleLinkedList(f func(interface{}) IDoubleCell) IDoubleLinkedList {
	return &DoubleLinkedList{
		newCell: f,
	}
}

//DoubleCell 双链表单元
type DoubleCell struct {
	data interface{}
	next IDoubleCell
	prev IDoubleCell //Previous 上一个
}

//Data return data
func (d *DoubleCell) Data() interface{} {
	return d.data
}

//SetData set data
func (d *DoubleCell) SetData(v interface{}) {
	d.data = v
}

//Next return next cell
func (d *DoubleCell) Next() IDoubleCell {
	return d.next
}

//SetNext set next cell
func (d *DoubleCell) SetNext(next IDoubleCell) {
	if d != nil {
		d.next = next
	}
}

//Prev get prev cell
func (d *DoubleCell) Prev() IDoubleCell {
	return d.prev
}

//SetPrev set prev cell
func (d *DoubleCell) SetPrev(prev IDoubleCell) {
	if d != nil {
		d.prev = prev
	}
}

//DoubleLinkedList 双链表
type DoubleLinkedList struct {
	head, tail IDoubleCell
	newCell    func(interface{}) IDoubleCell
}

//Append 追加元素
func (d *DoubleLinkedList) Append(data ...interface{}) {
	for _, cell := range data {
		if util.IsNil(d.tail) {
			d.head = d.newCell(cell)
			d.tail = d.head
		} else {
			d.tail.SetNext(d.newCell(cell))
			d.tail = d.tail.Next()
		}
	}
}

//Range 遍历
func (d *DoubleLinkedList) Range(f func(IDoubleCell) (breaked, removed bool)) {
	for cursor := d.head; reflect.ValueOf(cursor).IsNil(); cursor = cursor.Next() {
		breaked, removed := f(cursor)
		if removed {
			if util.IsNil(d.tail.Prev()) {
				d.Clear()
			} else {
				d.tail.Prev().SetNext(d.tail.Next())
				if util.IsNil(d.tail.Next()) {
					d.tail = d.tail.Prev()
				} else {
					d.tail = d.tail.Next()
				}
			}
		}
		if breaked {
			break
		}
	}
}

//Clear 清空
func (d *DoubleLinkedList) Clear() {
	d.head, d.tail = nil, nil
}

//Head 头元素
func (d *DoubleLinkedList) Head() IDoubleCell {
	return d.head
}

//Tail 最后一个元素
func (d *DoubleLinkedList) Tail() IDoubleCell {
	return d.tail
}
