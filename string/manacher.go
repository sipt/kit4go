package str

import (
	"strings"
	"fmt"
)

func Manacher(str string) int {
	str = "#" + strings.Join(strings.Split(str, ""), "#") + "#"
	fmt.Println(strings.Split(str, ""))
	maxLen := -1 // maxLen:最长回文长度
	mx := 0      // mx:已知回文的最右界
	id := 0      // id:已经回文的中心
	p := make([]int, len(str))
	for i := 1; i < len(str)-1; i++ {
		if i < mx {
			p[i] = min(p[2*id-i], mx-i) // 利用已知回文减少运算
		} else {
			p[i] = 1 // 无可利用，暴力运算
		}
		for i-p[i] >= 0 && i+p[i] < len(str) && str[i-p[i]] == str[i+p[i]] { // 不需边界判断，因为左有'$',右有'%', 两者不相等
			p[i]++
		}
		// 我们每走一步 i，都要和 mx 比较，我们希望 mx 尽可能的远，这样才能更有机会执行 if (i < mx)这句代码，从而提高效率
		if mx < i+p[i] {
			id = i
			mx = i + p[i]
		}
		maxLen = max(maxLen, p[i]-1)
	}
	return maxLen
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
