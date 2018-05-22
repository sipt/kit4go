package str

import (
	"testing"
)

func TestManacher(t *testing.T) {
	str := "abba"
	maxLen := Manacher(str)
	answer := 4
	if maxLen != answer {
		t.Errorf("str:[%s] reply:[%d], answer:[%d]", str, maxLen, answer)
	}

	str = "abcba"
	maxLen = Manacher(str)
	answer = 5
	if maxLen != answer {
		t.Errorf("str:[%s] reply:[%d], answer:[%d]", str, maxLen, answer)
	}

	str = "1abba"
	maxLen = Manacher(str)
	answer = 4
	if maxLen != answer {
		t.Errorf("str:[%s] reply:[%d], answer:[%d]", str, maxLen, answer)
	}

	str = "abba1"
	maxLen = Manacher(str)
	answer = 4
	if maxLen != answer {
		t.Errorf("str:[%s] reply:[%d], answer:[%d]", str, maxLen, answer)
	}

	str = "abbaacxcaa"
	maxLen = Manacher(str)
	answer = 7
	if maxLen != answer {
		t.Errorf("str:[%s] reply:[%d], answer:[%d]", str, maxLen, answer)
	}

	str = "asdfasdfasdfxfdsa123"
	maxLen = Manacher(str)
	answer = 9
	if maxLen != answer {
		t.Errorf("str:[%s] reply:[%d], answer:[%d]", str, maxLen, answer)
	}
}
