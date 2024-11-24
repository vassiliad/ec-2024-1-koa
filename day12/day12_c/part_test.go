package main

import (
	"puzzle/utilities"
	"testing"
)

func TestSmall(t *testing.T) {
	small := `6 5
6 7
10 5`

	input, err := utilities.ReadString(small)
	t.Logf("Input is %v\n", input)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 11

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestTiny(t *testing.T) {
	small := `6 7`

	input, err := utilities.ReadString(small)
	t.Logf("Input is %v\n", input)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 6

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestTiny2(t *testing.T) {
	small := `6 7`

	input, err := utilities.ReadString(small)
	t.Logf("Input is %v\n", input)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 6

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestTiny3(t *testing.T) {
	small := `10 10`

	input, err := utilities.ReadString(small)
	t.Logf("Input is %v\n", input)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 5

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestTiny4(t *testing.T) {
	small := `14 7`

	input, err := utilities.ReadString(small)
	t.Logf("Input is %v\n", input)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 4

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
