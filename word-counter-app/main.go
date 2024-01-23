package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//declare a new variable for command line flag:
	lines := flag.Bool("l", false, "CountLines too")
	bytes := flag.Bool("b", false, "CountBytes")
	//parse the flags provided by the user:
	flag.Parse()
	//calling the count function to count words or lines recieved from the standard input and
	//prints it out.
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
func count(r io.Reader, countLines bool, countBytes bool) int { // a function that reads
	//initializes a new Scanner with the provided io.Reader r

	scanner := bufio.NewScanner(r)
	
	//initial value of countline is set to false, negating the value using the logical operator !
	//which means countline is true: thus, executing the code block inside the if statement.
	if !countLines && !countBytes {
		
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	}


	// repeat same logic for counting bytes.

	wc := 0

	for scanner.Scan() {
		wc++
	}
	return wc
}
