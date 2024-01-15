package main

import(
	"bytes"
	"testing"
)
func TestCountWords(t *testing.T) {
	// Declared a variable to help read an existing string using the newBufferString function.
	b := bytes.NewBufferString("word1 word2 word3 word4")

	//a variable to define the result of the count(b) function.
	res := count(b, false)
	expected := 4 
	//if we fail to get then number declared in the "expected" it will return an error.
	if expected != res {
		t.Errorf("expected: %v, got: %v", expected, res)
	}
}

func testCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\nline2 \nline1word3 word4")
	res := count(b, true)
	exp := 3

	if exp != res {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

}