package main

import "fmt"

func main() {

	var m [6][6]int

	// var maxM [3][3]string

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			var v int
			fmt.Scanf("%d", &v)
			m[i][j] = v
		}
	}

	max := -100
	for r := 1; r < 5; r++ {
		for c := 1; c < 5; c++ {
			sum := m[r-1][c-1] + m[r-1][c] + m[r-1][c+1] +
				m[r+1][c-1] + m[r+1][c] + m[r+1][c+1] +
				m[r][c]

			if sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)
}
