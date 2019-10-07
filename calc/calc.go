package calc

import (
	"unicode"
)

func CalcRWords(txt string, letter rune) (numRWords int) {
	f := true
	for _, l := range []rune(txt) {
		if !unicode.IsLetter(l) {
			f = true
			continue
		}
		if l == letter && f {
			numRWords++
		}
		f = false
	}
	return numRWords
}
