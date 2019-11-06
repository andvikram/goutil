package binarysearch_test

import (
	bs "andvikram/goutil/search/binary_search"
	"math/rand"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7},
		{[]int{-3, -2, 4, 5, 8}, -2},
	}
	for _, tc := range testCases {
		el := tc.expected + 1
		got := bs.BinarySearch(tc.input, el)
		if tc.expected != got {
			t.Errorf("Expected %d, but got %d\n", tc.expected, got)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	collectionLength := b.N
	collection := make([]int, collectionLength)
	for i := 1; i <= collectionLength; i++ {
		collection[i-1] = i
	}
	el := rand.Intn(collectionLength)
	bs.BinarySearch(collection, el)
}
