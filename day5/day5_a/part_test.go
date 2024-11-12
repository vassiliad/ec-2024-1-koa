package main

import (
	"testing"

	"ec-2024-1-koa/day5_a/utilities"
)

func TestSmall(t *testing.T) {
	small := `2 3 4 5
3 4 5 2
4 5 2 3
5 2 3 4`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = 2323

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
