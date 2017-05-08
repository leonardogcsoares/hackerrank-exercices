package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var list []int
	for i := 0; i < n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		list = append(list, v)
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", list[i])
	}
}
