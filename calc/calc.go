package calc

import (
	"strings"
	"unicode"
)

const (
	CHAR_OTHER  = 0
	CHAR_LETTER = 1
)

func CalcRWords(txt, letter string) int {
	num_r_words := 0
	current_char_type := CHAR_OTHER
	word := ""
	for _, l := range txt {
		if unicode.IsLetter(l) {
			if current_char_type == CHAR_LETTER {
				word += string(l)
			} else {
				current_char_type = CHAR_LETTER
				word = string(l)
			}
		} else {
			if current_char_type == CHAR_LETTER {
				if strings.HasPrefix(strings.ToLower(word), strings.ToLower(letter)) {
					num_r_words += 1
				}
			}
			current_char_type = CHAR_OTHER
		}
	}

	return num_r_words
}

