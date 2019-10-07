package calc

import (
	"strings"
	"unicode"
)

func CalcWordsStartedWithLetter(txt, letter string) (q int) {
	words := strings.FieldsFunc(txt, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	for _, word := range words {
		if strings.ToLower(firstLetter(word)) == strings.ToLower(letter) {
			q++
		}
	}
	return q
}

func firstLetter(word string) string {
	if word == "" {
		return ""
	}
	return string([]rune(word)[0])
}
