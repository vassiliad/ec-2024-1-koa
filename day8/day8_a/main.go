package main

import (
	"os"

	"puzzle/utilities"
)

func solution(input *utilities.Puzzle) int {

	next_layer := 3
	width := 1
	total := 1

	for total = 1; total <= input.AvailableBlocks; {

		total += next_layer
		next_layer += 2
		width += 2
	}

	missing := total - input.AvailableBlocks

	return missing * width
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	// logger.Println("Input was", input)

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}

	sol := solution(input)

	logger.Println("Solution is", sol)
}
