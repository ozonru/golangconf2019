package calc

import (
	"unicode"
)

func CalcRWords(txt string, letter rune) int {
	count := 0
	prevRune := ' '

	for _, r := range txt {
		if !unicode.IsLetter(prevRune) && unicode.ToLower(r) == letter {
			count++
		}

		prevRune = r
	}

	return count
}
