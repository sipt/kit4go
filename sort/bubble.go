package sort

func BubbleSort(arr []interface{}, compare func(x, y interface{}) bool) {
	var needBreak bool
loop:
	for i := 0; i < len(arr)-1; i++ {
		needBreak = true
		for j := 0; j < len(arr)-1-i; j++ {
			if !compare(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				needBreak = false
			}
			if needBreak {
				break loop
			}
		}
	}
}
