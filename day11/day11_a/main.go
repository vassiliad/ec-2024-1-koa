package main

import (
	"os"

	"puzzle/utilities"
)

func solution(input *utilities.Puzzle) int {
	acc := map[string]int{}
	for k, _ := range input.Recipe {
		acc[k] = 0
	}

	next_round := map[string]int{}

	acc["A"] = 1

	for _ = range 4 {

		for k, _ := range input.Recipe {
			next_round[k] = 0
		}

		for k, mutations := range input.Recipe {
			for _, n := range mutations {
				next_round[n] += acc[k]
			}

		}

		acc, next_round = next_round, acc
	}

	sum := 0
	for _, t := range acc {
		sum += t
	}
	return sum
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
