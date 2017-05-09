package main

import (
	"container/heap"
	"fmt"
)

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Push TODO
func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop TODO
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An MaxHeap is a min-heap of ints.
type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Push TODO
func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop TODO
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var total int
	fmt.Scanf("%d\n", &total)

	uppers := &MinHeap{}
	lowers := &MaxHeap{}
	heap.Init(uppers)
	heap.Init(lowers)
	for i := 0; i < total; i++ {
		var n int
		fmt.Scanf("%d\n", &n)

		if lowers.Len() == 0 || getPeekLower(lowers) > n {
			heap.Push(lowers, n)
		} else {
			heap.Push(uppers, n)
		}
		rebalance(lowers, uppers)

		x := getMedian(lowers, uppers)

		fmt.Printf("%.1f\n", x)
	}
}

func rebalance(lowers *MaxHeap, uppers *MinHeap) {
	for lowers.Len()-uppers.Len() > 1 {
		l := heap.Pop(lowers).(int)
		heap.Push(uppers, l)
	}
	for uppers.Len()-lowers.Len() > 0 {
		u := heap.Pop(uppers).(int)
		heap.Push(lowers, u)
	}
}

func getMedian(lowers *MaxHeap, uppers *MinHeap) float64 {
	var median float64
	if lowers.Len() == uppers.Len() {
		l := getPeekLower(lowers)
		u := getPeekUpper(uppers)
		median = float64(l+u) / 2
	} else {
		l := getPeekLower(lowers)
		median = float64(l)
	}

	return median
}

func getPeekLower(lowers *MaxHeap) int {
	v := heap.Pop(lowers).(int)
	heap.Push(lowers, v)
	return v
}
func getPeekUpper(uppers *MinHeap) int {
	v := heap.Pop(uppers).(int)
	heap.Push(uppers, v)
	return v
}
