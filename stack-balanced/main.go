package main

import "fmt"

var opps = map[string]string{
	"{": "}",
	"[": "]",
	"(": ")",
}

func main() {
	var total int
	fmt.Scanf("%d\n", &total)

	for i := 0; i < total; i++ {
		var line string
		fmt.Scanln(&line)

		stack := []string{}
		for _, r := range line {
			stack = append(stack, string(r))
		}

		var closing []string
		var isNo bool
		for _, char := range stack {
			if val, ok := opps[char]; ok {
				closing = append(closing, val)
			} else if len(closing) == 0 || char != closing[len(closing)-1] {
				isNo = true
				break
			} else {
				closing = closing[:len(closing)-1]
			}
		}

		if isNo {
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
	}
}
