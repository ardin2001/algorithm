package main

import (
	"fmt"
)

func BinarySearch(number int, arr []int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (right + left) / 2
		if arr[middle] == number {
			return true
		} else if arr[middle] < number {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	if arr[right] != number {
		return false
	}
	return true
}

func main() {
	fmt.Println(BinarySearch(9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
