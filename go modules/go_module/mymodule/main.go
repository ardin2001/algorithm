package main

import (
	"fmt"

	"mymodule/mypackage"

	sortingArray "github.com/ardin2001/pkg_sort_go/v2"
)

func main() {
	fmt.Println("Hello, Modules!")
	mypackage.PrintHello()
	sortingArray.InsectionSort(99999999, 7, 5, 8, 1, 9, 3, 4, 2)
}
