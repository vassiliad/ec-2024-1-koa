package main

import (
	"os"

	"puzzle/utilities"
)

func prepare_face(wheel_idx []int, input *utilities.Puzzle) string {
	ret := []rune{}

	for i, idx := range wheel_idx {
		ret = append(ret, input.Wheels[i][idx]...)
		ret = append(ret, rune(' '))
	}

	return string(ret)
}

func count_score(wheel_idx []int, input *utilities.Puzzle) int {
	symbols := map[rune]int{}

	for i, idx := range wheel_idx {
		for j, c := range input.Wheels[i][idx] {
			if j != 1 {
				symbols[c] = symbols[c] + 1
			}

		}
	}

	coins := 0

	for _, count := range symbols {
		if count > 2 {
			coins += count - 2 // from 1 + (count - 3)
		}
	}

	return coins
}

func solution(input *utilities.Puzzle) int {
	// VV: The spins repeat every LCM of the wheel size, just need to calculate
	// the total number of repeats and then top it off with the remaining spins
	// to fill in the max_round spins after the last repeat
	wheel_idx := make([]int, len(input.Wheels))
	max_rounds := 202420242024

	total := 0
	round_score := map[int]int{}

	lcm := 1

	for _, wheels := range input.Wheels {
		lcm = utilities.LCM(lcm, len(wheels))
	}

	for round := range lcm {
		for i, spin := range input.Spins {
			wheel_idx[i] = (wheel_idx[i] + spin) % len(input.Wheels[i])
		}

		s := count_score(wheel_idx, input)
		total += s

		round_score[round] = total
	}

	top_up := round_score[max_rounds%lcm-1]

	// VV: (max_rounds/lcm) * 1 full "cycle" + top_up rounds past the last cycle
	return round_score[lcm-1]*(max_rounds/lcm) + top_up
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
