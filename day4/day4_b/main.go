package main

import (
	"os"

	"ec-2024-1-koa/day4_b/utilities"
)

func solution(input *utilities.Puzzle) int {
	sol := 0

	min := input.Nails[0]

	for _, n := range input.Nails {
		if n < min {
			min = n
		}
	}

	for _, n := range input.Nails {
		sol += n - min
	}

	return sol
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
