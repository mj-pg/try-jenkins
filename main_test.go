package main

import (
	"flag"
	"testing"
)

var failTest = flag.Bool("fail", false, "Flag to fail test")

func TestIsEven(t *testing.T) {
	input := 2
	if *failTest {
		input = 3
	}

	if even := isEven(input); !even {
		t.Fatalf("Number %d is not even.", input)
	}
}
