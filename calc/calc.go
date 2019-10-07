package calc

import (
	"unicode"
)

func CalcRWords(txt string, letter rune) int {
	letter = unicode.ToLower(letter)
	prev := ' '
	count := 0
	for _, c := range txt {
		switch !unicode.IsLetter(prev) {
		case true:
			switch unicode.ToLower(c) == letter {
			case true:
				count += 1
			case false:
			}
			fallthrough
		case false:
			prev = c
		}
	}
	return count
}
