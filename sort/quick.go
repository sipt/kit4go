package sort

func QuickSort(arr []interface{}, compare func(x, y interface{}) bool) {
	if len(arr) <= 1 {
		return
	}
	i, j := 0, len(arr)-1
	flag := true
	for i < j {
		if !compare(arr[i], arr[j]) {
			arr[i], arr[j] = arr[j], arr[i]
			flag = !flag
		}
		if !flag {
			i++
		} else {
			j --
		}
	}
	if i > 1 {
		QuickSort(arr[:i], compare)
	}
	if i+1 < len(arr)-2 {
		QuickSort(arr[i+1:], compare)
	}
}
