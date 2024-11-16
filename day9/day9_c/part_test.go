package main

import (
	"puzzle/utilities"

	"testing"
)

func TestSmall(t *testing.T) {
	small := `156488
352486
546212`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 10449

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
