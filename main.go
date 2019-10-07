package main

import (
	"fmt"
	"github.com/ozonru/golangconf2019/calc"
	"io/ioutil"
	"log"
	"strings"
)

/*
This script calculates the number of words in the text
*/

const FILENAME = "Exler.txt"
const LETTER = rune('р')

func main() {
	txt, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatalf("Failed to open text file: %s", err.Error())
	}
	num := calc.CalcRWords(strings.ToLower(string(txt)), LETTER)
	fmt.Printf("The text %s contains %d words starting from letter '%s'\n", FILENAME, num, string(LETTER))
}
