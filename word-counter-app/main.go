package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin))
}
func count(r io.Reader) int {
	//declare a variable to read text from a Reader using the bufio scanner package.
	scanner := bufio.NewScanner(r)

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