package main

import (
	"testing"

	"puzzle/utilities"
)

func TestSmall(t *testing.T) {
	small := `2
4
7
16`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 10

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestConvert(t *testing.T) {
	for value := range 9999999 {
		beetles := IntToBeetles(value)
		reconst := BeetlesToInt(beetles)

		if reconst != value {
			t.Fatal("Value", value, "converted to", beetles, "but reconstructed to", reconst)
		}
	}
}
