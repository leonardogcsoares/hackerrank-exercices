package main

import (
	"fmt"
	"math"
)

func main() {
	var size, rotate int
	// var nums []string
	var numStr string

	fmt.Scanf("%d %d\n", &size, &rotate)
	rotate = int(math.Mod(float64(rotate), float64(size)))

	fmt.Scanln(&numStr)
	// nums = strings.Split(numStr, " ")
	fmt.Println(numStr)
	// newNums := append(nums[rotate:], nums[:rotate]...)
	// fmt.Println(strings.Join(newNums, " "))
}
