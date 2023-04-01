package main

import (
	"fmt"
	"sync"
	"time"
)

var group = sync.WaitGroup{}
var cond = sync.NewCond(&sync.Mutex{})

func WaitCondition(val int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("jalan", val)
	cond.L.Unlock()
}
func main() {
	for i := 1; i <= 10; i++ {
		go WaitCondition(i)
	}
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	group.Wait()
	fmt.Println("finish")
}
