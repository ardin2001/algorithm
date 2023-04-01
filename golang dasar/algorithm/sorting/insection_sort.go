package main

import (
	"fmt"
)

func InsectionSort(numbers ...int) {
	for i := 1; i < len(numbers); i++ {
		for j := i - 1; j >= 0; j-- {
			if numbers[j+1] < numbers[j] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	fmt.Println(numbers)

}
func main() {
	InsectionSort(4, 2, 7, 3, 9, 5, 8, 1, 5)
	InsectionSort(4, 2, 7, 1, 56, 6, 3, 5, 8)
}
