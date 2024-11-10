package main

import (
	"testing"

	"ec-2024-1-koa/day2_a/utilities"
)

func TestSmall(t *testing.T) {
	small := `WORDS:THE,OWE,MES,ROD,HER

	AWAKEN THE POWER ADORNED WITH THE FLAMES BRIGHT IRE`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = 4

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
