package main

import (
	"testing"

	"ec-2024-1-koa/day7_a/utilities"
)

func TestSmall(t *testing.T) {
	small := `A:+,-,=,=
B:+,=,-,+
C:=,-,+,+
D:=,=,=,+`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = "BDCA"

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
