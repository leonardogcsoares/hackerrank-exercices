package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hotel struct {
	mentions int
	id       int
}

func main() {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}

	words := make(map[string]int)
	for _, w := range strings.Split(text, " ") {
		words[strings.ToLower(w)]++
	}

	var total int

	if scanner.Scan() {
		numStr := scanner.Text()
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return
		}

		total = num
	}

	hotels := make(map[int]hotel)
	for i := 0; i < total; i++ {
		var id int
		if scanner.Scan() {
			numStr := scanner.Text()
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return
			}

			id = num
		}

		if scanner.Scan() {
			text = scanner.Text()
		}
		text = removeDotsAndCommas(text)

		h := hotels[id]
		h.id = id
		setMentions(strings.Split(text, " "), words, &h)
		hotels[id] = h
	}

	mxH := &MaxHeap{}
	heap.Init(mxH)
	for _, v := range hotels {
		heap.Push(mxH, v)
	}

	for mxH.Len() > 0 {
		h := mxH.Pop().(hotel)
		fmt.Println(h)
		fmt.Printf("%d", h.id)
	}

}

// An MaxHeap is a min-heap of ints.
type MaxHeap []hotel

func (h *MaxHeap) Len() int { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool {
	if (*h)[i].mentions == (*h)[j].mentions {
		return (*h)[i].id < (*h)[j].id
	}

	return (*h)[i].mentions < (*h)[j].mentions
}
func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Push TODO
func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(hotel))
}

// Pop TODO
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func removeDotsAndCommas(text string) string {
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, ",", "", -1)
	return text
}

func setMentions(toCheck []string, words map[string]int, h *hotel) {
	for _, v := range toCheck {
		if _, ok := words[v]; ok {
			h.mentions++
		}
	}
}
