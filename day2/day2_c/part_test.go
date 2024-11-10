package main

import (
	"testing"

	"ec-2024-1-koa/day2_c/utilities"
)

func TestSmall(t *testing.T) {
	small := `WORDS:THE,OWE,MES,ROD,RODEO

HELWORLT
ENIGWDXL
TRODEOAL`

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

func TestTrivial(t *testing.T) {
	small := `WORDS:THE

TA
HB
EC
HB
TA`

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
