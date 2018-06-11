package scrabble

import "strings"

func AllScores() map[rune]int {
	scores := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10}
	output := make(map[rune]int)
	for runes, score := range scores {
		for _, r := range runes {
			output[r] = score
		}
	}
	return output
}

var scores map[rune]int = AllScores()

func Score(input string) int {
	score := 0
	for _, r := range strings.ToUpper(input) {
		score += scores[r]
	}
	return score
}
