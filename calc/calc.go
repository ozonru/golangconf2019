package calc

import (
	"strings"
	"unicode"
)

// WordsStartsWithLetter returns amount of words beginning with incoming letter
func WordsStartsWithLetter(txt, incomingLetter string) int {
	wordsCount := 0
	txtAsRunes := append([]rune{}, []rune(txt)...)
	incomingLetter = strings.ToLower(incomingLetter)

	for i := 1; i < len(txtAsRunes); i++ {
		letter := strings.ToLower(string(txtAsRunes[i]))
		if unicode.IsLetter(txtAsRunes[i]) && !unicode.IsLetter(txtAsRunes[i-1]) {
			if letter == incomingLetter {
				wordsCount++
			}
		}
	}

	return wordsCount
}
