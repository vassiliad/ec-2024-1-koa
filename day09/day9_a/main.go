package main

import (
	"os"

	"puzzle/utilities"
)

var DIGITS = []int{1, 3, 5, 10}

func IntToBeetles(value int) []int {
	beetles := make([]int, len(DIGITS))

	for i := len(DIGITS) - 1; i > -1; i-- {
		if value >= DIGITS[i] {
			d := value / DIGITS[i]
			value -= d * DIGITS[i]

			beetles[i] = d
		}
	}

	return beetles
}

func BeetlesToInt(beetles []int) int {
	ret := 0

	for i, b := range beetles {
		ret += DIGITS[i] * b
	}

	return ret
}

func solution(input *utilities.Puzzle) int {
	sol := 0

	for _, b := range input.Brightnesses {
		beetles := IntToBeetles(b)

		for _, d := range beetles {
			sol += d
		}
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
