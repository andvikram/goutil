//
//  This code implements the Binary Search
//

package search

// BinarySearch can search an element in a sorted and unique collection of integers in O(n) time
// It returns the index of the element
func BinarySearch(collection []int, el int) int {
	n := len(collection)
	if n == 1 {
		return 0
	}
	if n < 1 || el < collection[0] || el > collection[n-1] {
		return -1
	}
	l := 0
	r := n
	m := n / 2
	return bSearch(collection, el, l, r-1, m)
}

func bSearch(collection []int, el, l, r, m int) int {
	for {
		if el == collection[m] {
			return m
		}
		if el < collection[m] {
			r = m - 1
			m = (l + r) / 2
		}
		if el > collection[m] {
			l = m + 1
			m = (l + r) / 2
		}
	}

	// recursively
	// if el == collection[m] {
	// 	return m
	// }
	// if el < collection[m] {
	// 	r = m - 1
	// 	m = (l + r) / 2
	// 	m = bSearch(collection, el, l, r, m)
	// }
	// if el > collection[m] {
	// 	l = m + 1
	// 	m = (l + r) / 2
	// 	m = bSearch(collection, el, l, r, m)
	// }
	// return m
}
