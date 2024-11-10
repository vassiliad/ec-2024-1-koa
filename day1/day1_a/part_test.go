package main

import (
	"testing"

	"ec-2024-1-koa/day1_a/utilities"
)

func TestSmall(t *testing.T) {
	small := `ABBAC`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = 5

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
