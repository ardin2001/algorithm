package main

import (
	"fmt"
	"sync"
	"time"
)

func OnlyOnce() {
	fmt.Println("only sync success")
}
func main() {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			// sync.once di dalam functionnya tidak menerima parameter
			once.Do(OnlyOnce)
			//OnlyOnce()
		}()
	}
	time.Sleep(2 * time.Second)
}
