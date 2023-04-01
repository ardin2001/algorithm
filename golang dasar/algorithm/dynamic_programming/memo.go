package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	fmt.Println(n)
	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

var memo []int

func main() {
	memo = make([]int, 100)
	memo[1] = 1
	fmt.Println(fibonacci(6))
	fmt.Println(memo)
}
