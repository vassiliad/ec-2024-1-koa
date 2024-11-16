package main

import (
	"testing"

	"ec-2024-1-koa/day3_c/utilities"
)

func TestSmall(t *testing.T) {
	small := `..........
..###.##..
...####...
..######..
..######..
...####...
..........`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = 29

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
