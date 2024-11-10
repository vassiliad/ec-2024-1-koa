package main

import (
	"os"

	"ec-2024-1-koa/day1_b/utilities"
)

func potions_for_enemy(enemy *utilities.Enemy, other *utilities.Enemy) int {
	if *enemy == utilities.X {
		return 0
	}

	ret := int(*enemy)

	if *other != utilities.X {
		ret += 1
	}

	return ret
}

func solution(input []utilities.Enemy) int {
	sol := 0

	for i := 0; i < len(input); i += 2 {
		sol += potions_for_enemy(&input[i], &input[i+1]) + potions_for_enemy(&input[i+1], &input[i])
	}

	return sol
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	sol := solution(input)

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}

	logger.Println("Solution is", sol)
}
