//
//  This code implements the Quick Sort
//

package sort

// QuickSort sorts an collection of integers in O(n * logn) average time
func QuickSort(collection []int) []int {
	if len(collection) <= 1 {
		return collection
	}
	lo, hi := 0, len(collection)-1

	return qsort(collection, lo, hi)
}

func qsort(collection []int, lo, hi int) []int {
	if lo < hi {
		p := partition(collection, lo, hi)
		qsort(collection, lo, p)
		qsort(collection, p+1, hi)
	}

	return collection
}

func partition(collection []int, lo, hi int) int {
	i, j := lo, hi
	pivot := collection[(lo+hi)/2] // better way is the "median-of-three" method

	for {
		for collection[i] < pivot {
			i++
		}

		for collection[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		// swap
		collection[i], collection[j] = collection[j], collection[i]

		// contingency for duplicates
		if collection[i] == pivot {
			i++
		}
		if collection[j] == pivot {
			j--
		}
	}
}
