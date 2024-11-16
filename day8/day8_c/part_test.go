package main

import (
	"testing"

	"puzzle/utilities"
)

func TestSmall(t *testing.T) {
	small := `2`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input, 5, 160)
	const correct_answer = 2

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
