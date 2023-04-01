package main

import (
	"fmt"
	"runtime"
)

func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			fmt.Println("wswqexwex : ", counter)
			counter++
		}
	}()

	fmt.Println("done createcounter", destination)
	return destination
}
func main() {
	fmt.Println("hello")
	fmt.Println("Total goroutine : ", runtime.NumGoroutine())

	destination := CreateCounter()

	for n := range destination {
		fmt.Println("COUNTER n :", n)
		if n == 10 {
			break
		}
	}
	fmt.Println("Total goroutine : ", runtime.NumGoroutine())
}
