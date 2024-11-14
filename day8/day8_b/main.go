package main

import (
	"os"

	"puzzle/utilities"
)

func solution(input *utilities.Puzzle, acolytes, total_marble int) int {
	priests := input.Priests

	thickness := 1
	next_layer := 1
	width := 1
	total := 1

	for total = 1; total <= total_marble; {
		width += 2

		thickness = (priests * thickness) % acolytes
		next_layer = width * thickness
		total += next_layer
		// println("width", width, "total", total, "thickness", thickness, "next_layer", next_layer)
	}

	missing := total - total_marble

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

	const TOTAL_MARBLE = 20240000
	const ACOLYTES = 1111

	sol := solution(input, ACOLYTES, TOTAL_MARBLE)

	logger.Println("Solution is", sol)
}
