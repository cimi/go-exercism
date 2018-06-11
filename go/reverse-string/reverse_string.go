package reverse

// String reverses (!?) a string - apparently this is what the tests want.
func String(input string) string {
	reverse := []rune(input)
	for lo, hi := 0, len(reverse)-1; hi >= lo; lo, hi = lo+1, hi-1 {
		reverse[lo], reverse[hi] = reverse[hi], reverse[lo]
	}
	return string(reverse)
}
