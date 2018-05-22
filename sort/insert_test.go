package sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	result := []int{76, 64, 23, 9, 5, 4, 2, 1}
	arr := []interface{}{9, 2, 4, 5, 1, 23, 64, 76}
	InsertSort(arr, func(x, y interface{}) bool {
		return x.(int) > y.(int)
	})
	if len(result) != len(arr) {
		t.Errorf("result len : %d != %d", len(arr), len(result))
	} else {
		for k, v := range result {
			if v != arr[k].(int) {
				t.Errorf("result: %v, answer: %v", arr, result)
				return
			}
		}
	}
}
