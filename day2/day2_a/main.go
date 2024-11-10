package main

import (
	"os"
	"strings"

	"ec-2024-1-koa/day2_a/utilities"
)

func solution(input *utilities.Puzzle) int {
	sol := 0

	for _, word := range input.Words {
		sol += strings.Count(input.Text, word)
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
