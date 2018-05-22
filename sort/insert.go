package sort

func InsertSort(arr []interface{}, compare func(x, y interface{}) bool) {
	for i, j := 1, 0; i < len(arr); i++ {
		temp := arr[i]
		for j = i; j > 0 && !compare(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
