package main

import (
	"os"

	"puzzle/utilities"
)

func solution(input *utilities.Puzzle) int {
	sol := 0

	for r := 0; len(input.Targets) > 0; r++ {
		max_height := input.Targets[0].Y

		for _, t := range input.Targets {
			if t.Y > max_height {
				max_height = t.Y
			}
		}

		idx := r % len(input.Segments)
		idx = len(input.Segments) - idx

		s := input.Segments[idx]

		for power := range input.Width {
			y := power + s.Y
			x := 2*power + s.X

			for i, t := range input.Targets {
				dx := t.X - x
				dy := y - t.Y

				if dx > 0 && dx == dy && t.Y == max_height {

					sol += idx * power

					if len(input.Targets) == 1 {
						return sol
					}

					input.Targets = append(input.Targets[:i], input.Targets[i+1:]...)
					break
				}
			}
		}

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
