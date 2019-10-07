package calc

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	CHAR_OTHER  = 0
	CHAR_LETTER = 1
)

// CalcRWords returns count of words that starts with $letter.
func CalcRWords(txt, letter string) int {
	letterRune, _ := utf8.DecodeRuneInString(letter)
	letterUpperRune := unicode.ToUpper(letterRune)
	letterLowerRune := unicode.ToLower(letterRune)
	var numRWords int
	words := strings.Fields(txt)
	for _, word := range words {
		s := []byte(word)
		for utf8.RuneCount(s) >= 1 {
			r, size := utf8.DecodeRune(s)
			if !unicode.IsLetter(r) {
				s = s[size:]
				continue
			}
			if r == letterLowerRune || r == letterUpperRune {
				numRWords++
			}
			break
		}
	}

	return numRWords
}
