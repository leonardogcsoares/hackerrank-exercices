package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for h := 0; h < t; h++ {
		var num int
		fmt.Scanf("%d", &num)

		if num <= 1 {
			fmt.Println("Not prime")
			continue
		}
		if num == 2 {
			fmt.Println("Prime")
			continue
		}
		var isPrime = true
		for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
			if math.Mod(float64(num), float64(i)) == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}
