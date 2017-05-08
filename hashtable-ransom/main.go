package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanf("%d %d\n", &m, &n)

	var magazine = make(map[string]int)
	var ransom []string

	for i := 0; i < m; i++ {
		var word string
		fmt.Scanf("%s", &word)
		magazine[word] = magazine[word] + 1
	}
	for i := 0; i < n; i++ {
		var word string
		fmt.Scanf("%s", &word)
		ransom = append(ransom, word)
	}

	for _, word := range ransom {
		n, ok := magazine[word]
		if !ok {
			fmt.Println("No")
			return
		}
		if n == 0 {
			fmt.Println("No")
			return
		}

		magazine[word] = n - 1
	}

	fmt.Println("Yes")
}
