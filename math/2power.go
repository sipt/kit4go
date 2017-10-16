package math

//Is2Power 判断是否是2的乘方
func Is2Power(value int) bool {
	return value&(value-1) == 0
}
