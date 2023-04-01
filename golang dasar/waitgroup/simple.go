package main

import (
	"fmt"
	"sync"
	"time"
)

func runAsyncronous(group *sync.WaitGroup, val int) {
	defer group.Done()

	group.Add(1)
	fmt.Println("hallo :", val)
	time.Sleep(1 * time.Second)
}
func main() {
	group := &sync.WaitGroup{}
	counter := 0

	for i := 1; i <= 100; i++ {
		counter++
		go runAsyncronous(group, i)
	}

	group.Wait()
	fmt.Println("complete", counter)
}
