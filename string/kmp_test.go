package str

import "testing"

func TestGetNext(t *testing.T) {
	//str1 := "acacacacacabacacabacacacababab"
	str2 := "acacabacacabacacac"
	next := getNext(str2)
	reply := []int{0, 0, 1, 2, 3, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 4}
	for i, v := range next {
		if reply[i] != v {
			t.Errorf("reply not match:%v,%v", next, reply)
		}
	}
}

func TestIndexOf(t *testing.T) {
	str1 := "acacacacacabacacabacacacababab"
	str2 := "acacabacacabacacac"
	i := IndexOf(str1, str2)
	if i != 6 {
		t.Errorf("index failed: return :%v, answer:%v", i, 6)
	}
	str1 = "asdfasdf"
	str2 = "123123"
	i = IndexOf(str1, str2)
	if i != -1 {
		t.Errorf("index failed: return :%v, answer:%v", i, -1)
	}
	str1 = "asdfasdf"
	str2 = "asdfa"
	i = IndexOf(str1, str2)
	if i != 0 {
		t.Errorf("index failed: return :%v, answer:%v", i, 0)
	}
	str1 = "asdfasdfg"
	str2 = "sdfg"
	i = IndexOf(str1, str2)
	if i != 5 {
		t.Errorf("index failed: return :%v, answer:%v", i, 5)
	}
}
