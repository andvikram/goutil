//
//  This code implements the Counting Sort
//

package sort

// CountingSort sorts a collection of positive integers in O(n + k) time
func CountingSort(collection []int) []int {
	n := len(collection)
	if n <= 1 {
		return collection
	}

	sortedCol := make([]int, n)

	// First, find the maximum value in the collection
	k := getMax(collection)

	// Create a counting array with length as k + 1 more to accomodate 0 of the index
	carr := make([]int, k+1)

	// Store the frequencies of each distinct element of the collection,
	// by mapping its value as the index of counting array
	for _, v := range collection {
		carr[v]++
	}

	j := 0
	for i, v := range carr {
		if j >= n {
			break
		}
		// Counting array stores which element occurs how many times,
		// Add i in sortedCol[] according to the number of times i occured in the collection
		for v > 0 {
			sortedCol[j] = i
			j++
			v--
		}
	}

	return sortedCol
}

func getMax(collection []int) int {
	max := 0
	for _, v := range collection {
		if v > max {
			max = v
		}
	}

	return max
}
