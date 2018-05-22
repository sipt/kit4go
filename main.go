package main

import (
	"fmt"
	// "strings"
)

func main() {
	str1 := "acacacacacabacacabacacacababab"
	str2 := "acacabacacabacacac"
	fmt.Println(IndexOf(str1, str2))
}



func test1() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(len(arr), cap(arr))
	arr = append(arr, 1, 2)
	fmt.Println(len(arr), cap(arr))
	arr2 := arr[:len(arr)/2]
	arr2[1] = 999
	fmt.Println("arr:", arr, "\n", "arr2:", arr2)
	arr2 = append(arr2, 1000)
	fmt.Println("arr:", arr, "\n", "arr2:", arr2)
}
