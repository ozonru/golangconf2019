package calc

import (
	"strings"
	"unicode"
)

func WordsBeginsByLetter(txt string, letter rune) int {
	cnt := 0
	letter = unicode.ToLower(letter)
	lastControl := true

	for _, r := range strings.ToLower(txt) {
		if lastControl && unicode.ToLower(r) == letter {
			cnt++
		}
		lastControl = !unicode.IsLetter(r)
	}

	return cnt
}
