package main

import "fmt"

func add(m map[string]int, name string) {
	m[string(name[0])]++
	for i := 1; i < len(name) && i < 21; i++ {
		m[name[:i+1]]++
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	partialM := make(map[string]int)

	for i := 0; i < n; i++ {
		var cmd, str string
		fmt.Scanf("%s %s\n", &cmd, &str)

		switch cmd {
		case "add":
			add(partialM, str)
			// fmt.Println(names)
		case "find":
			fmt.Println(partialM[str])
		}
	}
}
