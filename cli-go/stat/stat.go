package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	word := flag.Bool("w", false, "Count word")
	byte := flag.Bool("b", false, "Count byte")
	filename := flag.String("file", "", "File to inspect")
	// Parsing the flags provided by the user
	flag.Parse()

	// If user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)

	}
	fmt.Println(count(*filename, *word, *byte))



}

func count(filename string, countWords bool, countBytes bool, ) int {
	input, err := os.Open(filename)

	r :=

	if err != nil {
		log.Fatal(err)
	}
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(input)

	// If the count words or bytes flag is set
	// the scanner will split type to words or bytes (default is split by lines)
	if countWords {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0

	// For every word or line scanned, add 1 to the counter
	for scanner.Scan() {
		wc++
	}
	// Return the total
	return wc
}
