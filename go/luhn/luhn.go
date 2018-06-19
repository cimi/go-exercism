package luhn

import (
	"regexp"
	"strings"
)

var onlyDigits = regexp.MustCompile("^\\d+$")

func UpdateDigit(b byte) byte {
	double := 2 * b
	if double > 9 {
		return double - 9
	}
	return double
}

func Luhn(input string) bool {
	var sum byte
	shouldUpdate := false
	for i := len(input) - 1; i >= 0; i-- {
		digit := input[i] - '0'
		if shouldUpdate {
			sum += UpdateDigit(digit)
		} else {
			sum += digit
		}
		shouldUpdate = !shouldUpdate
	}
	return sum%10 == 0
}

func Valid(input string) bool {
	digits := strings.Replace(input, " ", "", -1)
	if len(digits) < 2 {
		return false
	}
	if !onlyDigits.MatchString(digits) {
		return false
	}
	return Luhn(digits)
}
