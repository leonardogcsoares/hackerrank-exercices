package main

import "fmt"

type stack struct {
	Items  []int
	Height int
}

func (s *stack) push(val int) {
	s.Items = append(s.Items, val)
	s.Height += val
}

func (s *stack) pop() int {
	n := len(s.Items)
	tmp := s.Items[0]
	s.Items = s.Items[1:n]
	s.Height -= tmp
	return tmp
}

// Last todo
func (s *stack) Last() int {
	n := len(s.Items)
	if n <= 0 {
		return 0
	}
	return s.Items[0]
}
func main() {
	var size1, size2, size3 int
	fmt.Scanf("%d %d %d\n", &size1, &size2, &size3)

	var stks []*stack
	stks = append(stks, readIntoStack(size1))
	fmt.Scanf("%d")
	stks = append(stks, readIntoStack(size2))
	fmt.Scanf("%d")
	stks = append(stks, readIntoStack(size3))

	for !(stks[0].Height == stks[1].Height && stks[1].Height == stks[2].Height) {
		// fmt.Println("Before nextTopPop")
		// printStacks(stks)
		stkToPop := nextToPop(stks)
		if len(stkToPop.Items) <= 0 {
			break
		}

		stkToPop.pop()
		height := stkToPop.Height
		// fmt.Println("Before rebalance, current height:", height)
		// printStacks(stks)
		rebalance(stks, height)
	}

	fmt.Println(stks[0].Height)
}

func readIntoStack(size int) *stack {
	stk := &stack{Items: []int{}}
	for i := 0; i < size; i++ {
		var val int
		fmt.Scanf("%d", &val)
		stk.push(val)
	}
	return stk
}

func nextToPop(stks []*stack) *stack {
	tmpStk := stks[0]
	for i := 1; i < 3; i++ {
		if stks[i].Height > tmpStk.Height {
			tmpStk = stks[i]
		} else if stks[i].Height == tmpStk.Height {
			if stks[i].Last() > tmpStk.Last() {
				tmpStk = stks[i]
			}
		}
	}

	return tmpStk
}

func rebalance(stks []*stack, height int) {
	for i := 0; i < len(stks); i++ {
		if len(stks[i].Items) <= 0 {
			continue
		}
		for stks[i].Height-stks[i].Last() > height && stks[i].Height > height {
			stks[i].pop()
		}
	}
}

func printStacks(stks []*stack) {
	for _, stk := range stks {
		fmt.Println(stk)
	}
}
