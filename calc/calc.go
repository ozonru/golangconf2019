package calc

func GetTargetLetterStartWords(txt, letter string) int64 {
	var numTargetWords int64
	lastIsSpace := true
	for _, currentChar := range txt {
		currentCharAsString := string(currentChar)
		if lastIsSpace && currentCharAsString == letter {
			numTargetWords++
			lastIsSpace = false
			continue
		}
		lastIsSpace = currentCharAsString == " " || currentCharAsString == "\n"
	}

	return numTargetWords
}
