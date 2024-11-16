package main

import (
	"os"
	"puzzle/utilities"
	"slices"
)

var DIGITS = []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101}

func solution(input *utilities.Puzzle) int {
	// VV: We don't really care about the maximum brightness, just the brightness up
	// to 50 more than the half of the maximum brightness. This is because we're
	// effectively splitting a sparkball into 2 and making sure the 2 halfs differ
	// by at most 100
	max_brightness := slices.Max(input.Brightnesses)/2 + 100

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

		t := int(^uint(0) >> 1)

		// VV: Try all points within distance of 50 around the desired brightness which:
		// 1. when added up they amount to the desired brightness, and
		// 2. they differ by at most 100

		middle := b/2 - 50

		for one := range 101 {
			for two := range 101 {
				high := middle + two
				low := middle + one

				if high-low >= 100 || low-high >= 100 || high+low != b {
					continue
				}

				t = min(t, dp[low]+dp[high])
			}

		}

		sol += t
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
