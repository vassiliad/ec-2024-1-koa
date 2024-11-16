package main

import (
	"testing"

	"ec-2024-1-koa/day5_c/utilities"
)

func TestSmall(t *testing.T) {
	small := `2 3 4 5
6 7 8 9`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = 6584

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
