//
//  This code implements the Quick Sort
//

package sort

// QuickSort sorts an array of integers in O(n * logn) average time
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	lo, hi := 0, len(arr)-1

	return qsort(arr, lo, hi)
}

func qsort(arr []int, lo, hi int) []int {
	if lo < hi {
		p := partition(arr, lo, hi)
		qsort(arr, lo, p)
		qsort(arr, p+1, hi)
	}

	return arr
}

func partition(arr []int, lo, hi int) int {
	i, j := lo, hi
	pivot := arr[(lo+hi)/2]

	for {
		for arr[i] < pivot {
			i++
		}

		for arr[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		swap(arr, i, j)

		// contingency for duplicates
		if arr[i] == pivot {
			i++
		}
		if arr[j] == pivot {
			j--
		}
	}
}

func swap(arr []int, a, b int) {
	c := arr[b]
	arr[b] = arr[a]
	arr[a] = c
}
