package math

//Gcd 最大公约数
func Gcd(x, y int) int {
	if x == y {
		return x
	} else if x < y {
		return Gcd(y, x)
	}
	if x&1 == 0 && y&1 == 0 {
		return Gcd(x>>1, y>>1) << 1
	} else if x&1 == 0 && y&1 != 0 {
		return Gcd(x>>1, y)
	} else if x&1 != 0 && y&1 == 0 {
		return Gcd(x, y>>1)
	}
	return Gcd(x-y, y)
}
