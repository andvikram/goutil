// Package heap provides heap implementations
// using Go's container/heap interface
package heap

import (
	"container/heap"
)

// An MinHeapInt provides min heap of elements with type int
// The priority is of minimum value
type MinHeapInt interface {
	// Push will add element to the heap
	Push(int)
	// Pop will remove element from the heap
	// and return it
	Pop() int
}

type intHeapMin []int

type intMinHeap struct {
	ih *intHeapMin
}

// NewMinHeapInt returns a new heap for operations on type int
// Can be provided intial values
func NewMinHeapInt(ints ...int) MinHeapInt {
	ih := intHeapMin{}
	if ints != nil {
		ih = append(ih, ints...)
	}
	heap.Init(&ih)
	return &intMinHeap{ih: &ih}
}

func (h *intMinHeap) Push(e int) {
	heap.Push(h.ih, e)
}

func (h *intMinHeap) Pop() int {
	e := heap.Pop(h.ih).(int)
	return e
}

func (h intHeapMin) Len() int {
	return len(h)
}
func (h intHeapMin) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h intHeapMin) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push will add element to the heap
func (h *intHeapMin) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop will remove element from the heap
// and return it
func (h *intHeapMin) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
