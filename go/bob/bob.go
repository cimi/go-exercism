// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Bob answers 'Sure.' if you ask him a question.
func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

var r = regexp.MustCompile("[A-Z]")

// He answers 'Whoa, chill out!' if you yell at him.
func isShouting(remark string) bool {
	return remark == strings.ToUpper(remark) &&
		r.ReplaceAllString(remark, "") != remark
}

// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

// He says 'Fine. Be that way!' if you address him without actually saying
// anything.
func isSilent(remark string) bool {
	return remark == ""
}

// He answers 'Whatever.' to anything else.

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isSilent(remark) {
		return "Fine. Be that way!"
	} else if isQuestion(remark) && isShouting(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		return "Sure."
	} else if isShouting(remark) {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}
