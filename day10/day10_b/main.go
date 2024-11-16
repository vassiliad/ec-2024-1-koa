package main

import (
	"fmt"
	"os"
	"slices"

	"puzzle/utilities"
)

func decode_rune(input *utilities.Puzzle, dx, dy int) int {
	const size = 8
	ret := 0

	for {
		updated := false

		for y := 2; y < size-2; y++ {
			for x := 2; x < size-2; x++ {
				if input.Grid[y+dy][x+dx] != rune('.') {
					continue
				}

				common := []rune{}
				row := []rune{}
				column := []rune{}

				for i := 0; i < size; i++ {
					row = append(row, input.Grid[y+dy][i+dx])
					column = append(column, input.Grid[i+dy][x+dx])
				}
				for i := 0; i < size; i++ {
					if row[i] != rune('.') && row[i] != rune('*') && slices.Contains(column, row[i]) {
						common = append(common, row[i])
					}
				}

				if len(common) == 1 {
					input.Grid[y+dy][x+dx] = common[0]
					updated = true
				} else if len(common) > 0 {
					for l := 0; l < len(input.Grid); l++ {
						println(string(input.Grid[l]))
					}
					fmt.Printf("Too many common %s at %d, %d from row %s and column %s at delta %d %d\n", string(common), y, x, string(row), string(column), dy/9, dx/9)
					panic(common)
				}
			}
		}

		if !updated {
			break
		}
	}

	ret = 0
	idx := 1
	for y := 2; y < size-2; y++ {
		for x := 2; x < size-2; x++ {
			ret += (int(input.Grid[y+dy][x+dx]) - int(rune('A')) + 1) * idx
			idx++
		}
	}

	return ret
}

func solution(input *utilities.Puzzle) int {
	const block_size = 8 + 1
	ret := 0
	for dy := 0; dy < len(input.Grid); dy += block_size {
		for dx := 0; dx < len(input.Grid[0]); dx += block_size {
			ret += decode_rune(input, dx, dy)

		}
	}

	for l := 0; l < len(input.Grid); l++ {
		println(string(input.Grid[l]))
	}

	return ret
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
