package main

import "fmt"

func merge(a, b []int) []int {
	arr := make([]int, 0)
	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			arr = append(arr, a[0])
			a = a[1:]
		} else {
			arr = append(arr, b[0])
			b = b[1:]
		}
	}
	arr = append(arr, a...)
	arr = append(arr, b...)
	return arr
}
func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] < arr[1] {
			return arr
		}
		return []int{arr[1], arr[0]}
	}
	mid := len(arr) / 2
	left := arr[0:mid]
	right := arr[mid:]
	sortLeft := sort(left)
	sortRight := sort(right)
	return merge(sortLeft, sortRight)
}
func main() {
	arr := []int{8, 3, 4, 1, 9, 1, 55, 76, 34, 89}
	fmt.Println(sort(arr))
}
