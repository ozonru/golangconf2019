// This script calculates the number of words in the text
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"

	"github.com/ozonru/golangconf2019/calc"
)

var (
	Filename = flag.String("f", "Exler.txt", "File target for char calculate")
	Letter   = flag.String("l", "р", "Char for calc")
	BulkSize = flag.Int("bs", 10, "Line size per one goroutine")
)

func main() {
	file, err := os.Open(*Filename)
	if err != nil {
		log.Fatalf("Failed to open text file: %s", err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var res int64
	var oneBulk string
	var wg sync.WaitGroup
	for scanner.Scan() {
		oneBulk += " " + scanner.Text()
		if len(oneBulk) > *BulkSize {
			go func(line string) {
				wg.Add(1)
				atomic.AddInt64(&res, calc.GetTargetLetterStartWords(line, *Letter))
				wg.Done()
			}(oneBulk)
			oneBulk = ""
		}
	}
	wg.Wait()
	fmt.Printf("The file %q contains %d words starting from letter '%s'\n", *Filename, res, *Letter)
}
