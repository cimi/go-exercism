package robotname

import (
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digits = "0123456789"
var registry map[string]bool = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomByte(input string) byte {
	return input[rand.Intn(len(input))]
}

func newName() string {
	for {
		name := string([]byte{
			randomByte(letters),
			randomByte(letters),
			randomByte(digits),
			randomByte(digits),
			randomByte(digits)})
		if used := registry[name]; !used {
			registry[name] = true
			return name
		}
	}
}

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = newName()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}
