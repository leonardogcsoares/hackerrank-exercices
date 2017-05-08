package main

import "fmt"

func main() {
	var size, lines int
	fmt.Scanf("%d %d\n", &size, &lines)

	var list = make([]int, size)

	var max int
	for i := 0; i < lines; i++ {
		var start, end, toAdd int
		fmt.Scanf("%d %d %d\n", &start, &end, &toAdd)

		list[start-1] = list[start-1] + toAdd
		if end < size {
			list[end] = list[end] - toAdd
		}
		fmt.Println(list)
	}

	for i := 0; i < size; i++ {
		if max+list[i] > max {
			max = max + list[i]
		}
	}

	fmt.Println(max)
}

// 5 8
// 1 2 100
// 2 5 100
// 3 4 100
// 1 4 100
// 2 3 100
// 1 3 100
// 2 5 100
// 1 5 100
