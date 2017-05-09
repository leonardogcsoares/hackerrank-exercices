package main

import "fmt"

var opps = map[string]string{
	"{": "}",
	"[": "]",
	"(": ")",
}
var closing = map[string]struct{}{
	"}": struct{}{},
	"]": struct{}{},
	")": struct{}{},
}

type stack []string

func (s *stack) push(c string) {
	(*s) = append((*s), c)
}

func (s *stack) pop() string {
	n := len(*s)
	tmp := (*s)[n-1]
	*s = (*s)[0 : n-1]
	return tmp
}

func main() {
	var total int
	fmt.Scanf("%d\n", &total)

	for i := 0; i < total; i++ {
		var line string
		fmt.Scanln(&line)
		stk := &stack{}

		isNo := false
		for _, r := range line {
			c := string(r)
			if isClosing(c) {
				if len(*stk) > 0 {
					if !(stk.pop() == c) {
						isNo = true
						break
					}
				} else {
					isNo = true
					break
				}
				// if len(*stk) > 0 {
				// 	if !(stk.pop() == c) {
				// 		isNo = true
				// 		break
				// 	}
				// 	} else {
				// 		isNo = true
				// 		break
				// 	}
			} else {
				stk.push(getOpposite(c))
			}
		}

		if !isNo && len(*stk) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isClosing(c string) bool {
	_, ok := closing[c]
	return ok
}

func getOpposite(c string) string {
	opp := opps[c]
	return opp
}
