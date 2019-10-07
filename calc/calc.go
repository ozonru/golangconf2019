package calc

import (
	"unicode"
)

type CharacterType int

const (
	CharNotLetter CharacterType = 0
	CharLetter    CharacterType = 1
)

func CalcRWords(txt string, letter rune) int {
	lower := letter
	upper := letter
	currentCharType := CharNotLetter
	if unicode.IsLower(letter) {
		upper = unicode.ToUpper(letter)
	} else {
		lower = unicode.ToLower(letter)
	}

	foundRWords := 0
	for _, l := range txt {
		if unicode.IsLetter(l) {
			if currentCharType != CharLetter {
				currentCharType = CharLetter
				if l == lower || l == upper {
					foundRWords += 1
				}
			}
		} else {
			currentCharType = CharNotLetter
		}
	}

	return foundRWords
}
