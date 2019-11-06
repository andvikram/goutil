// Package queue provides generic queue implementations safe for concurrent operations
//
// The package uses generic type. To generate code for specific types follow the examples below.
//
// Example:
// To generate a queue with string values:
// 	genny -in queue.go -out string/queue_string.go -pkg queuestring gen "Item=string"
package queue

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the binary search tree
type Item generic.Type

// Queue provides methods for queue operations
type Queue interface {
	// Enqueue adds an item to the queue
	Enqueue(Item)
	// Dequeue removes an item from the queue and returns it
	Dequeue() Item
	// Peek returns the an from the queue but does not removes it
	Peek() Item
	// Empty checks if the queue is empty
	Empty() bool
	// Full checks if the queue is full
	Full() bool
	// Size returns the size of the queue
	Size() int
	// Print will print the queue elements
	Print()
}

const (
	defaultSize = 100
	upperLimit  = 10000000
)

type sqStruct struct {
	qR   []Item
	indx int
	len  int
	l    sync.RWMutex
}

var ln int

// NewQueue returns a new queue of Item elements,
// with default size of 100,
// which can be overwritten by an
// optional positive integer parameter.
//
// Paramter must be less than 10000000,
// else it panics!
func NewQueue(s ...int) Queue {
	ln = defaultSize
	if s != nil && s[0] > 0 {
		ln = s[0]
	}
	if ln > upperLimit {
		panic("Queue size exceeds upper limit")
	}
	return &sqStruct{
		qR:   make([]Item, ln),
		indx: ln - 1,
		len:  0,
	}
}

func (s *sqStruct) Enqueue(e Item) {
	s.l.Lock()
	defer s.l.Unlock()
	if s.indx >= 0 {
		s.qR[s.indx] = e
		s.len++
		s.indx--
	}
}

func (s *sqStruct) Dequeue() Item {
	s.l.Lock()
	defer s.l.Unlock()
	if s.len < 1 {
		return ""
	}
	indx := s.indx + s.len
	e := s.qR[indx]
	s.qR[indx] = ""
	s.len--
	if s.len == 0 && s.indx == 0 {
		s.reset()
	}
	return e
}

func (s *sqStruct) Peek() Item {
	s.l.RLock()
	defer s.l.RUnlock()
	e := s.qR[s.indx+s.len]
	return e
}

func (s *sqStruct) Empty() bool {
	s.l.RLock()
	defer s.l.RUnlock()
	truthy := s.len == 0
	return truthy
}

func (s *sqStruct) Full() bool {
	s.l.RLock()
	defer s.l.RUnlock()
	truthy := s.len >= ln-1
	return truthy
}

func (s *sqStruct) Size() int {
	s.l.RLock()
	defer s.l.RUnlock()
	l := s.len
	return l
}

func (s *sqStruct) Print() {
	s.l.RLock()
	defer s.l.RUnlock()
	i := s.indx + s.len
	for i >= 0 {
		if s.qR[i] == "" {
			break
		}
		fmt.Printf("%s ", s.qR[i])
		i--
	}
	fmt.Println()
}

func (s *sqStruct) reset() {
	s.l.Lock()
	defer s.l.Unlock()
	s.indx = ln - 1
	s.len = 0
}
