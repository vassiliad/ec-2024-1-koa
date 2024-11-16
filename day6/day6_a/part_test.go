package main

import (
	"testing"

	"ec-2024-1-koa/day6_a/utilities"
)

func TestSmall(t *testing.T) {
	small := `RR:A,B,C
A:D,E
B:F,@
C:G,H
D:@
E:@
F:@
G:@
H:@`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = "RRB@"

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
