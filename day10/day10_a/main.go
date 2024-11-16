package main

import (
	"os"
	"slices"

	"puzzle/utilities"
)

func solution(input *utilities.Puzzle) string {
	const size = 8

	for {
		updated := false

		for y := 2; y < size-2; y++ {
			for x := 2; x < size-2; x++ {
				if input.Grid[y][x] != rune('.') {
					continue
				}

				common := []rune{}
				row := []rune{}
				column := []rune{}

				for i := 0; i < size; i++ {
					row = append(row, input.Grid[y][i])
					column = append(column, input.Grid[i][x])
				}
				for i := 0; i < size; i++ {
					if row[i] != rune('.') && row[i] != rune('*') && slices.Contains(column, row[i]) {
						common = append(common, row[i])
					}
				}

				if len(common) == 1 {
					input.Grid[y][x] = common[0]
					updated = true
				} else if len(common) > 0 {
					panic(common)
				}
			}
		}

		if !updated {
			break
		}
	}

	ret := []rune{}

	for y := 2; y < size-2; y++ {
		for x := 2; x < size-2; x++ {
			ret = append(ret, input.Grid[y][x])
		}
	}

	return string(ret)
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
