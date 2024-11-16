package main

import (
	"testing"

	"ec-2024-1-koa/day2_b/utilities"
)

func TestSmall(t *testing.T) {
	small := `WORDS:THE,OWE,MES,ROD,HER

AWAKEN THE POWE ADORNED WITH THE FLAMES BRIGHT IRE
THE FLAME SHIELDED THE HEART OF THE KINGS
POWE PO WER P OWE R
THERE IS THE END`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = 37

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
