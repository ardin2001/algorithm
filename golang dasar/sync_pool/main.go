package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return "waiting"
		},
	}

	pool.Put("ardin")
	pool.Put("nugraha")
	pool.Put("anjaz")

	for i := 1; i <= 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
}
