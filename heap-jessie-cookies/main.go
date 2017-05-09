package main

import (
	"container/heap"
	"fmt"
)

// IntHeap todo
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push TODO
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop todo
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var n, sweetness int
	fmt.Scanf("%d %d\n", &n, &sweetness)

	minHeap := &IntHeap{}
	heap.Init(minHeap)

	for i := 0; i < n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		heap.Push(minHeap, v)
	}

	fmt.Printf("%d", mixCookiesN(minHeap, sweetness))
}

func mixCookiesN(minHeap *IntHeap, sweetness int) int {
	var count int

	for minHeap.Len() >= 1 {
		v1 := heap.Pop(minHeap).(int)
		if v1 >= sweetness {
			return count
		}

		if minHeap.Len() <= 0 {
			break
		}

		v2 := heap.Pop(minHeap).(int)
		heap.Push(minHeap, v1+2*v2)
		count++
	}

	return -1
}
