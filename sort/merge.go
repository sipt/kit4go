package sort

func MergeSort(arr []interface{}, compare func(x, y interface{}) bool) {
	//Divide
	l := len(arr)
	mid := l / 2
	if mid > 1 {
		MergeSort(arr[:mid], compare)
	}
	if mid < len(arr)-1 {
		MergeSort(arr[mid:], compare)
	}

	//Merge
	i, j := 0, mid
	c := make(chan interface{}, l)
	for i < mid || j < l {
		if i < mid && j < l {
			if compare(arr[i], arr[j]) {
				c <- arr[i]
				i++
			} else {
				c <- arr[j]
				j++
			}
		} else if i < mid {
			c <- arr[i]
			i++
		} else {
			c <- arr[j]
			j++
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = <-c
	}
}
