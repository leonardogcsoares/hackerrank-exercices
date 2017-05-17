package main

import (
	"fmt"
	"math"
)

/*
  The idea is given an array:
  [1 1 2 3 4 2 3 6 7]

  output all the duplicates considering
  1 <= x <= lenOfArray

  One way of solving is doing N*N loops through the array and appending to a HashMap
  any duplicate value which would lead to:
  O(n) time complexity

  Another solution would be to loop through the array appending to a map[int]int
  and counting how many times each item in the array occurs, then looping through
  the map checking for count > 1.
  Time: O(2n) = O(n)
  Space: O(n)

  A more complicated solution that would lead to Time: O(n) and Space: O(1) is:
  for each possible value index from:
  0 <= index <= lenArray - 1

  Each value would correspond to an index in the array, and we flip (invert the value) of any
  duplicated values, then just check the array for negative values.
*/

func main() {
	array := []int{1, 2, 1, 2, 3, 4, 5, 2}

	fmt.Println(getDuplicates(array))
}

/*
array = [1 2 1 2]

index = 0
results = []
array = [-1 2 1 2]

index = 2-1 = 1
array = [-1 -2 1 2]
results = 0

index = 1-1 = 0
results = [1]
array = [-1 -2 1 2]

index = 2-1 = 1
results [1 2]
array = [-1 -2 1 2]
*/

func getDuplicates(array []int) []int {
	duplicates := make(map[int]struct{})
	var dupArr []int

	for i := 0; i < len(array); i++ {
		index := int(math.Abs(float64(array[i]))) - 1

		if array[index] < 0 {
			duplicates[int(math.Abs(float64(array[i])))] = struct{}{}
		} else {
			array[index] = -array[index]
		}
	}

	for k := range duplicates {
		dupArr = append(dupArr, k)
	}

	return dupArr
}
