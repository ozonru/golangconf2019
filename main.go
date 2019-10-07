package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ozonru/golangconf2019/calc"
	"github.com/uber-go/multierr"
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
This script calculates the number of words in the text
*/

const (
	defaultFilename = "Exler.txt"
	defaultLetter   = "р"
)

func main() {
	// \todo hack russian letter as an argument
	// \todo graceful shutdown
	filename := flag.String("file", defaultFilename, "file to parse")
	flag.Parse()

	letter := defaultLetter
	num, err := calculate(*filename)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("Text %s contains %d words starting from letter '%s'\n", *filename, num,
		letter)
}

func calculate(filename string) (int64, error) {
	fd, err := os.Open(filename)
	defer func() {
		if closeErr := fd.Close(); closeErr != nil {
			err = multierr.Append(err, closeErr)
		}
	}()

	workersNum := runtime.NumCPU()
	linesChan := make(chan string, workersNum)
	var wg sync.WaitGroup
	wg.Add(workersNum)
	var num int64
	for i := 0; i < workersNum; i++ {
		go func() {
			defer wg.Done()
			worker(linesChan, defaultLetter, &num)
		}()
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		linesChan <- scanner.Text()
	}
	close(linesChan)
	wg.Wait()

	return num, nil
}

func worker(linesChan chan string, letter string, num *int64) {
	var delta int64
	for line := range linesChan {
		n := int64(calc.CalcRWords(line, letter))
		delta += n
	}
	atomic.AddInt64(num, delta)
}
