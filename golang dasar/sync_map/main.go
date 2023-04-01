package main

import (
	"fmt"
	"sync"
)

func addMap(data *sync.Map, group *sync.WaitGroup, val int) {
	defer group.Done()

	group.Add(1)
	data.Store(val, val)
}
func main() {
	data := sync.Map{}
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go addMap(&data, &group, i)
	}
	group.Wait()
	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
