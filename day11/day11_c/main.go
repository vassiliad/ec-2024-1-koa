package main

import (
	"os"
	"slices"

	"puzzle/utilities"
)

func mutate(input *utilities.Puzzle, start string) uint64 {
	acc := map[string]uint64{}
	for k, _ := range input.Recipe {
		acc[k] = 0
	}

	next_round := map[string]uint64{}

	acc[start] = 1

	for _ = range 20 {

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

	var sum uint64 = 0
	for _, t := range acc {
		sum += t
	}
	return sum
}

func solution(input *utilities.Puzzle) uint64 {
	scenarios := []uint64{}

	for k := range input.Recipe {
		scenarios = append(scenarios, mutate(input, k))
	}

	min := slices.Min(scenarios)
	max := slices.Max(scenarios)

	return max - min
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
