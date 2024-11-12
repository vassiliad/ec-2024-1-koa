package main

import (
	"testing"

	"ec-2024-1-koa/day4_c/utilities"
)

func TestSmall(t *testing.T) {
	small := `2
4
5
6
8`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = 8

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
