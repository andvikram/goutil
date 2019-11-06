// Package heap provides heap implementations
// using Go's container/heap interface
package heap

import (
	"container/heap"
)

// PriorityQ provides priority queue
type PriorityQ interface {
	// Push will add element to the heap
	Push(string, int) int
	// Pop will remove element from the heap
	// and return it
	Pop() (string, int)
	// Update will update values
	// for an element in the queue
	Update(int, string, int)
}

type item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type priorityQueue []*item

type pqWrap struct {
	pq priorityQueue
}

// NewPriorityQ returns a new priority queue
func NewPriorityQ() PriorityQ {
	pw := new(pqWrap)
	heap.Init(&(pw.pq))
	return pw
}

func (pw *pqWrap) Push(v string, p int) int {
	i := new(item)
	i.value = v
	i.priority = p
	heap.Push(&(pw.pq), i)
	li := pw.pq[len(pw.pq)-1]
	return li.index
}

func (pw *pqWrap) Pop() (string, int) {
	i := heap.Pop(&(pw.pq)).(*item)
	return i.value, i.priority
}

func (pw *pqWrap) Update(indx int, v string, p int) {
	i := pw.pq[indx]
	i.value = v
	i.priority = p
	heap.Fix(&(pw.pq), indx)
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest priority
	// so we use greater than here
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	i := x.(*item)
	i.index = n
	*pq = append(*pq, i)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	i := old[n-1]
	old[n-1] = nil // avoid memory leak
	i.index = -1   // for safety
	*pq = old[0 : n-1]
	return i
}
