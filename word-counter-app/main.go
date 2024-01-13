package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin, *lines))
	//using the Boolfunction under the flag package to count add an option to also count lines.
	lines := flag.Bool("l", false, "Count lines")
	//parsing the flags value provided above.
	flag.Parse()
}
func count(r io.Reader, countLines bool) int {
	//declare a variable to read text from a Reader using the bufio scanner package.
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	//define the scanner split type to words (default is split by lines.)
	scanner.Split(bufio.ScanWords)

	//defining a counter 
	wc := 0

	//for every word scanned increment the counter..

	for scanner.Scan() {
		 wc++
	}
	return wc
}