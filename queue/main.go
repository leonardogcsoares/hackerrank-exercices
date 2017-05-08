package main

import "fmt"

type queue []int

func (q *queue) Enqueue(x int) {
	*q = append(*q, x)
}

func (q *queue) Dequeue() {
	if len(*q) > 0 {
		*q = (*q)[1:]
	}
}

func (q *queue) Print() int {
	if len(*q) > 0 {
		return (*q)[0]
	}

	return -1
}

func main() {

	var q queue
	var cmds int
	fmt.Scanf("%d\n", &cmds)

	var toPrint []int
	for i := 0; i < cmds; i++ {
		var cmd int
		fmt.Scanf("%d", &cmd)

		switch cmd {
		case 1:
			var v int
			fmt.Scanf("%d", &v)
			q.Enqueue(v)

		case 2:
			q.Dequeue()

		case 3:
			if v := q.Print(); v > 0 {
				toPrint = append(toPrint, v)
			}
		}
	}

	for _, v := range toPrint {
		fmt.Println(v)
	}

}
