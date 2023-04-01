package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// timer := time.NewTimer(5 * time.Second)
	// fmt.Println(time.Now())

	// time := <-timer.C
	// fmt.Println(time)

	// timer := time.After(5 * time.Second)
	// fmt.Println(time.Now())

	// time := <-timer
	// fmt.Println(time)

	group := sync.WaitGroup{}
	group.Add(1)
	fmt.Println("start")
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Execute after 3 second")
		group.Done()
	})
	group.Wait()
	fmt.Println("finish")

}
