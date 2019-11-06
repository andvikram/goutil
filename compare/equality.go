// Package compare contains operations for comparison between data structures
package compare

// IntSliceEqual tells whether int slice a and b contain
// the same integers in the same order.
// A nil argument is equivalent to an empty slice.
func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// StringSliceEqual tells whether string slice a and b contain
// the same integers in the same order.
// A nil argument is equivalent to an empty slice.
func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
