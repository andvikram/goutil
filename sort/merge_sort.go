//
//  This code implements the Merge Sort
//

package sort

// MergeSort sorts an collection of integers in O(n * logn) time
func MergeSort(collection []int) []int {
	if len(collection) <= 1 {
		return collection
	}

	m := len(collection) / 2
	l := collection[:m]
	r := collection[m:]

	return merge(MergeSort(l), MergeSort(r))
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)
	indexLeft, indexRight := 0, 0

	for (indexLeft < len(left)) && (indexRight < len(right)) {
		if left[indexLeft] < right[indexRight] {
			result = append(result, left[indexLeft])
			indexLeft++
		} else {
			result = append(result, right[indexRight])
			indexRight++
		}
	}

	// append the left overs as indexLeft and indexRight can grow disproportionately
	if indexLeft < len(left) {
		result = append(result, left[indexLeft:]...)
	}
	if indexRight < len(right) {
		result = append(result, right[indexRight:]...)
	}

	return result
}
