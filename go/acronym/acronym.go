package acronym

import (
	"strings"
)

func AbbreviateWord(word string) string {
	parts := strings.Split(word, "-")
	if len(parts) > 1 {
		abbrv := ""
		for _, part := range parts {
			abbrv += AbbreviateWord(part)
		}
		return abbrv
	} else {
		return strings.ToUpper(string(word[0]))
	}
}

// Abbreviate returns the acronym of the given string.
func Abbreviate(s string) string {
	words := strings.Split(s, " ")
	abbrv := ""
	for _, word := range words {
		abbrv += AbbreviateWord(word)
	}
	return abbrv
}
