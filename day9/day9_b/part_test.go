package main

import (
	"puzzle/utilities"

	"testing"
)

func TestSmall(t *testing.T) {
	small := `33
41
55
99`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 10

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
