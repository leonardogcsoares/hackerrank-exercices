package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var size, rotate int
	_, err := fmt.Scanf("%d %d\n", &size, &rotate)
	if err != nil {
		return
	}

	var nums []string
	for i := 0; i < size; i++ {
		var n string
		_, err := fmt.Scanf("%s", &n)
		if err != nil {
			return
		}
		nums = append(nums, n)
	}

	rotate = int(math.Mod(float64(rotate), float64(size)))
	newNums := append(nums[rotate:], nums[:rotate]...)
	fmt.Println(strings.Join(newNums, " "))
}
