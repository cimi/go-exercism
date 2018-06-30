package robotname

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digits = "0123456789"
var registry = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomByte(input string) byte {
	return input[rand.Intn(len(input))]
}

const maxTries = 5

func newName() (string, error) {
	for tries := 0; tries < maxTries; tries++ {
		name := string([]byte{
			randomByte(letters),
			randomByte(letters),
			randomByte(digits),
			randomByte(digits),
			randomByte(digits)})
		if used := registry[name]; !used {
			registry[name] = true
			return name, nil
		}
	}
	return "", fmt.Errorf("Could not generate an unseen random name after %d tries", maxTries)
}

func (r *Robot) Name() string {
	if r.name == "" {
		name, err := newName()
		if err != nil {
			log.Fatal(err)
		}
		r.name = name
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}
