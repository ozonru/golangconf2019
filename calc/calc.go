package calc

import (
	"strings"
	"unicode"
)

const (
	CHAR_OTHER  = 0
	CHAR_LETTER = 1
)

func CalcRWords(txt string) int {
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
				if strings.HasPrefix(word, "р") || strings.HasPrefix(word, "Р") {
					num_r_words += 1
				}
			}
			current_char_type = CHAR_OTHER
		}
	}

	//last word
	if current_char_type == CHAR_LETTER {
		if strings.HasPrefix(word, "р") || strings.HasPrefix(word, "Р") {
			num_r_words += 1
		}
	}

	return num_r_words
}

