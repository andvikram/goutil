package queuestring_test

import (
	"andvikram/goutil/compare"
	queuestring "andvikram/goutil/queue/string"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestStringQueue(t *testing.T) {
	collectionLength := 4
	stringQueue := queuestring.NewQueue(collectionLength)
	expected := []string{"a", "b", "c", "d"}
	got := make([]string, 0)
	if !stringQueue.Empty() {
		t.Errorf("Expected queue to be empty, but wasn't")
	}
	for _, e := range expected {
		stringQueue.Enqueue(e)
	}
	if !stringQueue.Full() {
		t.Errorf("Expected queue to be full, but wasn't")
	}
	if stringQueue.Size() != collectionLength {
		t.Errorf("Expected queue size to be %d, got %d", collectionLength, stringQueue.Size())
	}
	if stringQueue.Peek() != "a" {
		t.Errorf("Expected first item to be %s, got %s", "a", stringQueue.Peek())
	}
	for i := 0; i < collectionLength; i++ {
		got = append(got, stringQueue.Dequeue())
	}
	if !compare.StringSliceEqual(got, expected) {
		t.Errorf("Expected popped items to be %v, got %v", expected, got)
	}
}

func BenchmarkStringQueue(b *testing.B) {
	start := time.Now()
	collectionLength := b.N
	fmt.Println("collectionLength:", collectionLength)
	collection := make([]string, collectionLength)
	stringQueue := queuestring.NewQueue(collectionLength)
	for i := 0; i < collectionLength; i++ {
		collection[i] = fmt.Sprintf("%d", rand.Intn(collectionLength))
	}
	fmt.Println("Setup time:", time.Since(start))
	start = time.Now()
	b.ResetTimer()
	for _, e := range collection {
		stringQueue.Enqueue(e)
	}
	fmt.Println("Enqueue time:", time.Since(start))
	start = time.Now()
	for i := 0; i < collectionLength; i++ {
		stringQueue.Dequeue()
	}
	fmt.Println("Dequeue time:", time.Since(start))
}
