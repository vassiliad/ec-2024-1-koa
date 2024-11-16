package main

import (
	"os"

	"ec-2024-1-koa/day1_a/utilities"
)

func solution(input []utilities.Enemy) int {
	sol := 0

	for _, enemy := range input {
		sol += int(enemy)
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
