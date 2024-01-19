package main

import (
	"bytes"
	"testing"
)

func testCountWords(t *testing.T) {
	words := "i love you"
	b := bytes.NewBufferString(words)
	res := count(b, false, false)
	exp := 3
	if res != exp {
		t.Errorf("expected %v, got: %v", exp, res)
	}
}
