package main

import (
	"ec-2024-1-koa/day2_b/utilities"
	"os"
	"strings"
)

// Modified func from strings.Count
func mark_runic_symbol_indices(s, substr string, hit_indices []int) {

	n := 0
	for {
		i := strings.Index(s, substr)

		if i == -1 {
			return
		}

		for j := range len(substr) {
			hit_indices[n+i+j] = 1
		}

		// VV: This is to account for words that overlap with themselves
		n += i + 1
		s = s[i+1:]
	}
}

func solution(input *utilities.Puzzle) int {
	sol := 0

	for _, text := range input.Text {
		hit_indices := make([]int, len(text))

		for _, word := range input.Words {
			mark_runic_symbol_indices(text, word, hit_indices)
		}

		score := 0
		for _, hit := range hit_indices {
			score += hit
		}

		sol += score
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
