package main

import (
	"os"
	"puzzle/utilities"
	"slices"
)

var DIGITS = []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30}

func solution(input *utilities.Puzzle) int {

	max_brightness := slices.Max(input.Brightnesses)

	dp := make([]int, max_brightness+1)

	for i := range dp {
		// VV: We can use N SingleDot stamps for any of the desired brightness levels
		dp[i] = i
	}

	for b := range max_brightness + 1 {
		for _, s := range DIGITS {
			if b >= s {
				dp[b] = min(dp[b], 1+dp[b-s])
			}
		}
	}

	sol := 0
	for _, b := range input.Brightnesses {
		sol += dp[b]
	}

	return sol
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
