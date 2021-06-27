package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// Defining boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")

	// Defining boolean flag -b to count byte
	countByte := flag.Bool("b", false, "Count Bytes")

	// Parsing the flags provided by the user
	flag.Parse()

	// calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *countByte))
}

func count(reader io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from the Reader (such as files)
	scanner := bufio.NewScanner(reader)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines && !countBytes {
		// Define the scanner split type to words (default is split by lines)
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Define a counter
	wc := 0

	// For every word scanned increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
