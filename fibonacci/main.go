package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	recur := fibRecur(8)
	fmt.Println("Fib Recursion: ", recur, time.Since(start))

	start = time.Now()
	loop := fibIter(8)
	fmt.Println("Fib Loop: ", loop, time.Since(start))
}

// Example of Fibonacci sequence
// 1 1 2 3 5 8 13 21 ...

func fibRecur(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return fibRecur(n-1) + fibRecur(n-2)
}

func fibIter(n int) int {
	start, end := 0, 1

	for index := 0; index < n; index++ {
		tmp := start + end
		start = end
		end = tmp
	}

	return start
}
