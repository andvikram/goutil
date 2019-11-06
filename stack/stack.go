// Package stack provides stack implementations safe for concurrent operations
//
// The package uses generic type. To generate code for specific types follow the examples below.
//
// Example:
// To generate a stack with string values:
// 	genny -in stack.go -out string/stack_string.go -pkg stackstring gen "Item=string"
package stack

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the binary search tree
type Item generic.Type

// Stack provides methods for stack operations
type Stack interface {
	// Push allows pushing of Item element onto the stack
	Push(Item)
	// Pop will return the last element from the stack and remove it
	Pop() Item
	// Peek returns the an from the stack but does not removes it
	Peek() Item
	// Empty checks if the stack is empty
	Empty() bool
	// Full checks if the stack is full
	Full() bool
	// Size returns the current size of the stack
	Size() int
	// Print will print the stack elements
	Print()
}

const (
	defaultSize = 100
	upperLimit  = 10000000
)

type ssStruct struct {
	stackR []Item
	indx   int
	l      sync.RWMutex
}

var ln int

// NewStack returns a new stack for Item elements,
// with default size of 100,
// which can be overwritten by an
// optional positive integer parameter.
//
// Paramter must be less than 10000000,
// else it panics!
func NewStack(s ...int) Stack {
	ln = defaultSize
	if s != nil && s[0] > 0 {
		ln = s[0]
	}
	if ln > upperLimit {
		panic("Stack size exceeds upper limit")
	}
	return &ssStruct{
		stackR: make([]Item, ln),
		indx:   0,
	}
}

// Push allows pushing of Item element onto the stack
func (s *ssStruct) Push(el Item) {
	s.l.Lock()
	defer s.l.Unlock()
	if s.indx < len(s.stackR) {
		s.stackR[s.indx] = el
		s.indx++
	}
}

// Pop will get the last element from the stack, and delete it
func (s *ssStruct) Pop() Item {
	s.l.Lock()
	defer s.l.Unlock()
	if s.indx < 1 {
		return ""
	}
	s.indx--
	el := s.stackR[s.indx]
	s.stackR[s.indx] = ""
	return el
}

// Peek returns the an from the stack but does not removes it
func (s *ssStruct) Peek() Item {
	s.l.RLock()
	defer s.l.RUnlock()
	if s.indx < 1 {
		return ""
	}
	e := s.stackR[s.indx-1]
	return e
}

// Empty checks if the stack is empty
func (s *ssStruct) Empty() bool {
	s.l.RLock()
	defer s.l.RUnlock()
	truthy := s.indx == 0
	return truthy
}

// Full checks if the stack is full
func (s *ssStruct) Full() bool {
	s.l.RLock()
	defer s.l.RUnlock()
	truthy := s.indx >= ln
	return truthy
}

// Size returns the current size of the stack
func (s *ssStruct) Size() int {
	s.l.RLock()
	defer s.l.RUnlock()
	l := s.indx
	return l
}

// Print will print the stack elements
func (s *ssStruct) Print() {
	s.l.RLock()
	defer s.l.RUnlock()
	for i := 0; i <= s.indx; i++ {
		fmt.Printf("%s ", s.stackR[i])
	}
	fmt.Println()
}
