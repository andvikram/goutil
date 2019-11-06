package stackstring_test

import (
	"andvikram/goutil/compare"
	stackstring "andvikram/goutil/stack/string"
	"fmt"
	"math/rand"
	"testing"
)

func TestStringStack(t *testing.T) {
	collectionLength := 4
	stringStack := stackstring.NewStack(collectionLength)
	collection := []string{"a", "b", "c", "d"}
	expected := []string{"d", "c", "b", "a"}
	got := make([]string, 0)
	if !stringStack.Empty() {
		t.Errorf("Expected stack to be empty, but wasn't")
	}
	for _, e := range collection {
		stringStack.Push(e)
	}
	if !stringStack.Full() {
		t.Errorf("Expected stack to be full, but wasn't")
	}
	if stringStack.Size() != collectionLength {
		t.Errorf("Expected stack size to be %d, got %d", collectionLength, stringStack.Size())
	}
	if stringStack.Peek() != "d" {
		t.Errorf("Expected first item to be %s, got %s", "d", stringStack.Peek())
	}
	for i := 0; i < len(collection); i++ {
		got = append(got, stringStack.Pop())
	}
	if !compare.StringSliceEqual(got, expected) {
		t.Errorf("Expected popped items to be %v, got %v", expected, got)
	}
}

func BenchmarkStringStack(b *testing.B) {
	collectionLength := b.N
	collection := make([]string, collectionLength)
	stringStack := stackstring.NewStack(collectionLength)
	for i := 0; i < collectionLength; i++ {
		collection[i] = fmt.Sprintf("%d", rand.Intn(collectionLength))
	}
	b.ResetTimer()
	for _, e := range collection {
		stringStack.Push(e)
	}
	for i := 0; i < collectionLength; i++ {
		stringStack.Pop()
	}
}
