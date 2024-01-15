package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"flag"
)

func main() {
	//declare a new variable for command line flag:
	lines := flag.Bool("l", false, "CountLines too")
	//parse the flags provided by the user:
	flag.Parse()
	//calling the count function to count words or lines recieved from the standard input and 
	//prints it out.
	fmt.Println(count(os.Stdin, *lines))
}
func count(r io.Reader, countLines bool) int { // a function that reads
	//initializes a new Scanner with the provided io.Reader r
	
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	//scanner.Split(bufio.ScanWords)

	wc := 0

	for scanner.Scan() {
		wc++
	}
	return wc
}
