package main

import "fmt"

type node struct {
	CurMax int
	Val    int
}

type stack []node

func (s *stack) push(nd node) {
	(*s) = append((*s), nd)
}

func (s *stack) pop() node {
	n := len(*s)
	tmp := (*s)[n-1]
	*s = (*s)[0 : n-1]
	return tmp
}

func (s *stack) peek() node {
	return (*s)[len(*s)-1]
}

func main() {
	var total int
	fmt.Scanf("%d", &total)

	stk := &stack{}
	for i := 0; i < total; i++ {
		var cmd int
		fmt.Scanf("%d", &cmd)

		switch cmd {
		// Push
		case 1:
			var v int
			fmt.Scanf("%d", &v)

			if len(*stk) <= 0 {
				stk.push(node{Val: v, CurMax: v})
				continue
			}

			if curMax := stk.peek().CurMax; v > curMax {
				stk.push(node{Val: v, CurMax: v})
			} else {
				stk.push(node{Val: v, CurMax: curMax})
			}

		// Pop
		case 2:
			if len(*stk) > 0 {
				stk.pop()
			}

		// Print
		case 3:
			fmt.Println(stk.peek().CurMax)
		}
	}
}
