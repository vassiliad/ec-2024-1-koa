package main

import (
	"os"

	"puzzle/utilities"
)

func solution(input *utilities.Puzzle, acolytes, total_marble int) int {
	priests := input.Priests

	thickness := 1
	next_layer := 0000 // uninitialized
	width := 1
	total := 1
	height := 1
	layers := []int{thickness}

	// VV: Build the tower and then carve out its internal blocks
	for total = 1; total < total_marble; {
		width += 2

		thickness = ((thickness * priests) % acolytes) + acolytes
		next_layer = width * thickness
		total += next_layer

		layers = append(layers, thickness)
		height += thickness
	}

	// VV: The very last layer adds a single column whose height is the same
	// as the entire height of the tower. The code below assumes that each layer
	// adds 2 columns so here we adjust for this edge case
	correction := (priests * width * height) % acolytes

	empty := 0

	// VV: Skip the outtermost columns
	for idx := range len(layers) - 1 {
		thickness = layers[idx]

		// VV: Every new layer introduces 2 columns whose height/thickness is the same as the
		// height of the entire tower when the builders finished these 2 columns, which at the
		// time were external ones
		carve_out := (priests * width * height) % acolytes
		empty += 2 * carve_out

		height -= thickness
	}

	missing := (total + correction - empty) - total_marble

	// VV: sanity check
	if next_layer <= missing {
		panic("Oh no, you need to add one more level to your pyramid before carving out blocks")
	}

	return missing
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	// logger.Println("Input was", input)

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}

	const TOTAL_MARBLE = 202400000
	const ACOLYTES = 10

	sol := solution(input, ACOLYTES, TOTAL_MARBLE)

	logger.Println("Solution is", sol)
}
