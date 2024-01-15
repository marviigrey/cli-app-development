package main

import (
	"testing"
)
// a simple reminder to understand testing in golang modules.
func tetEven(t *testing.T){
	b := 4
	res := evenTest(b)
	expected := "yes"
	if expected != res {
		t.Errorf("expected: %v, got: %v", expected, res)
	}
}