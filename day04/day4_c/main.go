package main

import (
	"os"

	"ec-2024-1-koa/day4_c/utilities"
)

func bang_hammer_many_times(input *utilities.Puzzle, desired_height int) int {
	sol := 0

	for _, n := range input.Nails {
		t := n - desired_height
		if t < 0 {
			t = -t
		}
		sol += t
	}

	return sol
}

func solution(input *utilities.Puzzle) int {
	min := bang_hammer_many_times(input, input.Nails[0])

	for _, n := range input.Nails[1:] {
		t := bang_hammer_many_times(input, n)
		if t < min {
			min = t
		}
	}

	return min
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	// logger.Println("Input was", input)

	if err != nil {
		logger.Fatalln("Ran into problems while reading input. Problem", err)
	}

	sol := solution(input)

	logger.Println("Solution is", sol)
}
