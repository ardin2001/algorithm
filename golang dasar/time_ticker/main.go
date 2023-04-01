package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(1 * time.Second)

	for times := range ticker {
		fmt.Println(times)
	}
}
