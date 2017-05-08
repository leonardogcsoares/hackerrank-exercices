package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var an1, an2 string
	fmt.Scanf("%s\n", &an1)
	fmt.Scanf("%s\n", &an2)

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	var s1, s2 int

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, letter := range an1 {
			v := m1[letter]
			m1[letter] = v + 1
			s1++
		}
	}()

	for _, letter := range an2 {
		v := m2[letter]
		m2[letter] = v + 1
		s2++
	}

	wg.Wait()

	count := 0
	for k, v := range m1 {
		abs := m2[k] - v
		m1[k], m2[k] = 0, 0
		count = count + int(math.Abs(float64(abs)))
	}
	for k, v := range m2 {
		abs := m1[k] - v
		count = count + int(math.Abs(float64(abs)))
	}

	fmt.Println(count)
}
