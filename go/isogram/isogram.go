package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(input string) bool {
	seen := make(map[rune]bool)
	for _, r := range strings.ToUpper(input) {
		if seen[r] {
			return false
		}
		if unicode.IsLetter(r) {
			seen[r] = true
		}
	}
	return true
}
