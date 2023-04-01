// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {

// 	var x = 0
// 	var mutex sync.RWMutex

// 	for i := 1; i <= 1000; i++ {
// 		go func() {
// 			for i := 1; i <= 100; i++ {
// 				mutex.Lock()
// 				x += 1
// 				mutex.Unlock()
// 			}
// 		}()
// 	}

// 	time.Sleep(3 * time.Second)
// 	mutex.RLock()
// 	fmt.Println("getX :", x)
// 	mutex.RUnlock()
// }
