package calc

import (
	"strings"
)

func CalcRWords(txt, letter string) int {
	count := 0
	words := strings.Fields(txt)

	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, "\" \n'!.,-"))

		if strings.HasPrefix(word, letter) {
			count += 1
		}
	}

	return count
}
