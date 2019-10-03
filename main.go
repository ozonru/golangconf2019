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

func main() {
	txt, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatalf("Failed to open text file: %s", err.Error())
	}
	num := calc.CalcRWords(string(txt))
	fmt.Printf("The text contains %d words starting from russian 'р'\n", num)
}
