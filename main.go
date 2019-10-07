package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ozonru/golangconf2019/calc"
)

/*
This script calculates the number of words in the text
*/

const FILENAME = "Exler.txt"
const LETTER = "р"

func main() {
	txt, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatalf("Failed to open text file: %s", err.Error())
	}
	num := calc.WordsStartsWithLetter(string(txt), LETTER)
	fmt.Printf("The text %s contains %d words starting from letter '%s'\n", FILENAME, num, LETTER)
}
