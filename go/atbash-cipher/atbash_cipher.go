package atbash

import (
	"strings"
	"unicode"
)

type Cipher struct {
	b        strings.Builder
	alphabet []rune
	numChars int
}

func NewCipher(alphabet []rune) *Cipher {
	return &Cipher{alphabet: alphabet}
}

func (c *Cipher) Add(r rune) {
	c.numChars += 1
	c.b.WriteRune(r)
	if c.numChars%5 == 0 {
		c.b.WriteString(" ")
	}
}

func (c *Cipher) String() string {
	return c.b.String()
}

func (c *Cipher) Encode(r rune) bool {
	for pos := range c.alphabet {
		if c.alphabet[pos] == unicode.ToLower(r) {
			c.Add(c.alphabet[len(c.alphabet)-pos-1])
			return true
		}
	}
	if unicode.IsDigit(r) {
		c.Add(r)
		return true
	}
	return false
}

func Atbash(input string) string {
	cipher := NewCipher([]rune("abcdefghijklmnopqrstuvwxyz"))
	for _, r := range input {
		cipher.Encode(r)
	}
	return strings.TrimSpace(cipher.String())
}
