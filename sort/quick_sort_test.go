package sort_test

import (
	"andvikram/goutil/compare"
	"andvikram/goutil/sort"
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 6, 2, 7, 1, 4, 9, 5, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{4, -3, 8, 5, -2}, []int{-3, -2, 4, 5, 8}},
	}
	for _, tc := range testCases {
		got := sort.QuickSort(tc.input)
		if !compare.IntSliceEqual(tc.expected, got) {
			t.Errorf("Expected:\n%v\nGot:\n%v", tc.expected, got)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	collectionLength := b.N
	collection := make([]int, collectionLength)
	expected := make([]int, collectionLength)
	for i := 0; i < collectionLength; i++ {
		expected[i] = i + 1
	}
	copy(collection, expected)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(collectionLength,
		func(i, j int) {
			collection[i], collection[j] = collection[j], collection[i]
		},
	)
	b.ResetTimer()
	sort.QuickSort(collection)
}
