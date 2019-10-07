package calc

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func CalcRWords(text, letter string) int {
	numRWords := 0

	search, size := utf8.DecodeRune([]byte(letter))
	if search == utf8.RuneError || size != len(letter) || !unicode.IsLetter(search) {
		panic("invalid search letter")
	}

	text = strings.TrimSpace(text)
	start := true
	for _, l := range text {
		if unicode.IsSpace(l) {
			start = true
			continue
		}

		if start {
			if l == search {
				numRWords++
			}
			start = false
			continue
		}
	}

	return numRWords
}
