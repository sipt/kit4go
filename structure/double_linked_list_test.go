package structure_test

import (
	"testing"

	"github.com/sipt/algorithm/structure"
	"github.com/stretchr/testify/assert"
)

func TestDoubleLinedList(t *testing.T) {
	var testData = []int{1, 2, 3, 4, 7}
	var testData1 = []int{2, 3, 4, 7}
	var index = 0
	list := structure.NewDoubleLinkedList(structure.NewDoubleCell)
	for _, cell := range testData {
		list.Append(cell)
	}
	list.Range(func(cell structure.IDoubleCell) (bool, bool) {
		data := cell.Data()
		if index == 3 {
			return false, true // test break
		} else if index == 4 {
			return true, false // test remove
		}
		assert.Equal(t, data, testData[index], "value not equal, %d != %d", data, testData[index])
		cell.SetData(testData1[index]) // test update
		index++
		return false, false
	})
	index = 0
	list.Range(func(cell structure.IDoubleCell) (bool, bool) {
		assert.Equal(t, cell.Data(), testData1[index])
		index++
		return false, false
	})
}
