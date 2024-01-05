package main

import (
	"bytes"
	"testing"
)

// a test function to test our word counter.
func TestCountWords(t *testing.T) {
	
	//create a variable to initialize a new buffer to read content of type string.

	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4 //created a varible for the number of words in the mentioned string

	res := count(b) //using the count fucntion to count the number of words inside b string.

	if res != exp {
		t.Errorf("Expected: %d got: %d", exp, res)
	}
}
