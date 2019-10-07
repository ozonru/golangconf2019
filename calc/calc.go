package calc

import (
	"unicode"
)

func CountWordsStartingWithLetter(txt string, letter rune) int {
	count := 0
	prev := ' '
	for _, c := range txt {
		if !unicode.IsLetter(prev) && unicode.ToLower(c) == letter {
			count++
		}
		prev = c
	}

	return count
}
