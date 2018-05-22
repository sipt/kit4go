package str

func IndexOf(str, sub string) int {
	s1, s2 := []byte(str), []byte(sub)
	next := getNext(sub)
	i, j := 0, 0
	for i < len(s1) {
		for i < len(s1) {
			if s2[j] == s1[i] {
				j++
				i++
				if j >= len(s2) {
					return i - len(s2)
				}
			} else if j > 0 {
				j = next[j-1]
			} else {
				break
			}
		}
		if j == 0 {
			i ++
		}
	}
	if i-len(s1) > 0 {
		return i - len(s1)
	} else {
		return -1
	}
}

func getNext(sub string) []int {
	bytes := []byte(sub)
	next := make([]int, len(bytes))
	next[0] = 0
	for i, j := 1, 0; i < len(bytes); {
		for i < len(bytes) {
			if bytes[j] == bytes[i] {
				next[i] = j + 1
				j++
				i++
			} else if j > 0 {
				j = next[j-1]
			} else {
				break
			}
		}
		if j == 0 {
			next[i] = 0
			i ++
		}
	}
	return next
}
