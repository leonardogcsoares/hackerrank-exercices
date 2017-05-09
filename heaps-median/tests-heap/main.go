package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Peek TODO
func (h *MinHeap) Peek() interface{} {
	v := heap.Pop(h).(int)
	heap.Push(h, v)
	return v
}

// Push TODO
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

//Pop TODO
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Peek TODO
func (h *MaxHeap) Peek() interface{} {
	v := heap.Pop(h).(int)
	heap.Push(h, v)
	return v
}

// Push TODO
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

//Pop TODO
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	uppers := &MaxHeap{6, 3, 2}
	lowers := &MaxHeap{3, 2, 6}

	heap.Init(uppers)
	heap.Init(lowers)

	fmt.Println("Uppers: ", uppers)
	v := heap.Pop(uppers).(int)
	fmt.Println("Pop:", v)
	heap.Push(uppers, v)
	fmt.Println("After Push:", heap.Pop(uppers).(int))
	// fmt.Println("Lowers: ", lowers)
	// fmt.Println(heap.Pop(lowers).(int))
	// fmt.Println(lowers.Peek().(int))

}
