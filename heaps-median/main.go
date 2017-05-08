package main

import (
	"container/heap"
	"fmt"
)

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

// Seek TODO
func (h *MinHeap) Seek() interface{} {
	old := *h
	n := len(old)
	return old[n-1]
}

// An MaxHeap is a min-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

// Seek TODO
func (h *MaxHeap) Seek() interface{} {
	old := *h
	n := len(old)
	return old[n-1]
}

func main() {
	var total int
	fmt.Scanf("%d\n", &total)

	minH := &MinHeap{}
	maxH := &MaxHeap{}
	heap.Init(minH)
	heap.Init(maxH)
	for i := 0; i < total; i++ {
		var n int
		fmt.Scanf("%d\n", &n)

		if minH.Len() > 0 && n <= heap.Pop(minH).(int) {
			heap.Push(minH, n)
		} else {
			heap.Push(maxH, n)
		}

		for minH.Len() > maxH.Len() {
			heap.Push(maxH, heap.Pop(minH))
		}
		for (maxH.Len() - minH.Len()) > 1 {
			heap.Push(minH, heap.Pop(maxH))
		}

		var x float64
		if minH.Len() == maxH.Len() {
			fmt.Println("Heaps equal in len")
			x = float64(heap.Pop(minH).(int)+heap.Pop(maxH).(int)) / 2
		} else {
			// x = float64((*maxH)[maxH.Len()-1])
			fmt.Println("Heaps disequal in len")
			x = float64(heap.Pop(maxH).(int))
		}

		fmt.Printf("%.1f\n", x)
	}
}
