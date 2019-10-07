package main

import (
	"fmt"
	"github.com/ozonru/golangconf2019/calc"
	"io/ioutil"
	"log"
)

/*
This script calculates the number of words in the text
*/

const FILENAME = "Exler.txt"
const LETTER = 'р'

func main() {
	txt, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatalf("Failed to open text file: %s", err.Error())
	}
	num := calc.WordsBeginsByLetter(string(txt), LETTER)
	fmt.Printf("The text %s contains %d words starting from letter %q\n", FILENAME, num, LETTER)
}
