package main

import (
	"os"

	"ec-2024-1-koa/day3_c/utilities"
)

func neighbour(input *utilities.Puzzle, x, y, dx, dy int) int {
	if x+dx < 0 || x+dx >= input.Width || y+dy < 0 || y+dy >= input.Height {
		return -1
	}

	return input.Blocks[x+dx+(y+dy)*input.Width]
}

func solution(input *utilities.Puzzle) int {
	sol := 0

	for _, level := range input.Blocks {
		if level != -1 {
			sol += 1
		}
	}

	frame := make([]int, len(input.Blocks))

	deltas := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

	for {
		copy(frame, input.Blocks)

		dug := 0
		for y := 0; y < input.Height; y++ {
			for x := 0; x < input.Width; x++ {
				me := input.Blocks[x+y*input.Width]
				if me == -1 {
					continue
				}
				can_dig := true

				me += 1

				for _, d := range deltas {
					dl := me - neighbour(input, x, y, d[0], d[1])

					if dl < -1 || dl > 1 {
						can_dig = false
						break
					}
				}

				if can_dig {
					frame[x+y*input.Width] += 1
					dug += 1
				}
			}
		}

		frame, input.Blocks = input.Blocks, frame

		sol += dug

		// println("Dug", dug)

		// for y := 0; y < input.Height; y++ {
		// 	for x := 0; x < input.Width; x++ {
		// 		l := input.Blocks[x+y*input.Width]
		// 		if l == -1 {
		// 			print(".")
		// 		} else {
		// 			print(l)
		// 		}
		// 	}
		// 	println()
		// }

		if dug == 0 {
			break
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
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}

	sol := solution(input)

	logger.Println("Solution is", sol)
}
