package main

import "fmt"

func SelectionSort(arr ...int) []int {
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		index := 0
		for j := i; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				index = j
			}
		}
		if arr[i] > min {
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
	return arr
}

func main() {
	fmt.Println(SelectionSort(7, 3, 6, 2, 4, 9, 8))
}
