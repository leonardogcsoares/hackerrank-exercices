package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var s1, s2, s3, s4 int

	scanner := bufio.NewScanner(os.Stdin)
	other := 0
	squares := 0
	rects := 0

	for scanner.Scan() {
		fmt.Fscanf(strings.NewReader(scanner.Text()), "%d %d %d %d", &s1, &s2, &s3, &s4)
		fmt.Println(s1, s2, s3, s4)
		if isOther(s1, s2, s3, s4) {
			other++
		} else if isSquare(s1, s2, s3, s4) {
			squares++
		} else if isRect(s1, s2, s3, s4) {
			rects++
		} else {
			other++
		}
	}

	fmt.Printf("%d %d %d", squares, rects, other)
}

func isOther(s1, s2, s3, s4 int) bool {
	if s1 <= 0 || s2 <= 0 || s3 <= 0 || s4 <= 0 {
		return true
	}

	return false
}

func isSquare(s1, s2, s3, s4 int) bool {
	if s1 == s2 && s2 == s3 && s3 == s4 {
		return true
	}

	return false
}

func isRect(s1, s2, s3, s4 int) bool {
	if s1 == s3 && s2 == s4 {
		return true
	}

	return false
}
