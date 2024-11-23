package main

import (
	"testing"

	"puzzle/utilities"
)

func TestSmall(t *testing.T) {
	small := `##########.##########
#...................#
#.###.##.###.##.#.#.#
#..A#.#..~~~....#A#.#
#.#...#.~~~~~...#.#.#
#.#.#.#.~~~~~.#.#.#.#
#...#.#.B~~~B.#.#...#
#...#....BBB..#....##
#Z............#....Z#
#####################`

	println(small)

	input, err := utilities.ReadString(small)
	t.Logf("Input is %+v\n", input)

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)
	const correct_answer = 38

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}