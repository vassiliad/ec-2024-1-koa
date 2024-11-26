package main

import (
	"os"
	"strings"

	"puzzle/utilities"
)

func solution(input *utilities.Puzzle) string {
	wheel_idx := make([]int, len(input.Wheels))

	for range 100 {
		for i, spin := range input.Spins {
			wheel_idx[i] = (wheel_idx[i] + spin) % len(input.Wheels[i])
		}
	}
	ret := []string{}

	for w, faces := range input.Wheels {
		ret = append(ret, string(faces[wheel_idx[w]]))
	}
	return strings.Join(ret, " ")
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
