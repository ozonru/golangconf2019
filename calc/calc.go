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
	prev_char_type := CHAR_OTHER
	word := ""
	for _, l := range txt {
		if unicode.IsLetter(l) {
			word += string(l)
			if prev_char_type != CHAR_LETTER {
				prev_char_type = CHAR_LETTER
			}
			continue
		}
		if prev_char_type == CHAR_OTHER {
			continue
		}

		// not letter and prev type - letter
		if strings.HasPrefix(strings.ToLower(word), strings.ToLower(letter)) {
			num_r_words += 1
		}

		prev_char_type = CHAR_OTHER
		word = ""
	}

	if word != "" && strings.HasPrefix(strings.ToLower(word), strings.ToLower(letter)) {
		num_r_words += 1
	}

	return num_r_words
}
