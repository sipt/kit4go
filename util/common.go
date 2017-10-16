package util

import "reflect"

//IsNil 判断是否为空
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsNil()
}
