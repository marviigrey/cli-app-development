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
func count(r io.Reader) int { // a function that reads
	//initializes a new Scanner with the provided io.Reader r
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	wc := 0

	for scanner.Scan() {
		wc++
	}
	return wc
}
