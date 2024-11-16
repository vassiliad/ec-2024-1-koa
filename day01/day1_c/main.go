package main

import (
	"os"

	"ec-2024-1-koa/day1_c/utilities"
)

func potions_for_enemy(enemy_index int, group []utilities.Enemy) int {
	ret := 0

	for idx, enemy := range group {
		if idx == enemy_index {
			if enemy == utilities.X {
				return 0
			} else {
				ret += int(enemy)
			}

		} else {
			if enemy != utilities.X {
				ret += 1
			}
		}

	}

	return ret
}

func solution(input []utilities.Enemy) int {
	sol := 0
	const group_size = 3

	for i := 0; i < len(input); i += group_size {

		t := 0
		for enemy_idx := range 3 {
			t += potions_for_enemy(enemy_idx, input[i:i+3])
		}

		sol += t
	}

	return sol
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	sol := solution(input)

	if err != nil {
		logger.Fatalln("Ran into problems while reading input. Problem", err)
	}

	logger.Println("Solution is", sol)
}
