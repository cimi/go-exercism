package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Regexp test")
	r := regexp.MustCompile("[A-Z]")
	fmt.Println(r.ReplaceAllString("A1 BC", ""))
}
