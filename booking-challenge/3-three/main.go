package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const escapeToken = "-128"

func main() {
	var numStr string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		// fmt.Fscanf(strings.NewReader(scanner.Text()), "%d %d %d %d", &s1, &s2, &s3, &s4)
		// fmt.Println(s1, s2, s3, s4)
		numStr = scanner.Text()
	}

	var nums []int
	for _, v := range strings.Split(numStr, " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	if len(nums) == 0 {
		return
	}

	var toPrint []string
	previous := nums[0]
	if previous < 127 || previous > 127 {
		toPrint = append(toPrint, strconv.Itoa(previous))
	} else {
		toPrint = append(toPrint, strconv.Itoa(previous))
	}

	var deltas []int
	for i := 1; i < len(nums); i++ {
		fmt.Println(nums[i], previous, nums[i]-previous)
		delta := nums[i] - previous
		deltas = append(deltas, delta)
		previous = nums[i]
	}

	for _, d := range deltas {
		if d > 127 || d < -127 {
			toPrint = append(toPrint, escapeToken)
			toPrint = append(toPrint, strconv.Itoa(d))
			continue
		}

		toPrint = append(toPrint, strconv.Itoa(d))
	}

	fmt.Println(strings.Join(toPrint, " "))

}
