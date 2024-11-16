package main

import (
	"os"
	"slices"

	"puzzle/utilities"
)

func measure_rune(input *utilities.Puzzle, dx, dy int) int {
	const size = 8
	ret := 0
	idx := 1
	for y := 2; y < size-2; y++ {
		for x := 2; x < size-2; x++ {
			if input.Grid[y+dy][x+dx] == rune('.') {
				return -1
			}
			ret += (int(input.Grid[y+dy][x+dx]) - int(rune('A')) + 1) * idx

			idx++
		}
	}

	return ret
}

func decode_rune(input *utilities.Puzzle, dx, dy int) int {
	const size = 8
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
				count := map[rune]int{}

				// VV: In some cases there can be a unique `?` symbol in both row and column,
				// track i here so that we can fill in its value
				var loc_question *Point = nil

				for i := 0; i < size; i++ {
					row = append(row, input.Grid[y+dy][i+dx])
					column = append(column, input.Grid[i+dy][x+dx])

					count[input.Grid[i+dy][x+dx]] += 1

					if input.Grid[i+dy][x+dx] == rune('?') {
						loc_question = &Point{y: i + dy, x: x + dx}
					}

					count[input.Grid[y+dy][i+dx]] += 1

					if input.Grid[y+dy][i+dx] == rune('?') {
						loc_question = &Point{y: y + dy, x: i + dx}
					}
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
					// VV: This is an unsolvable block
					return -2
				} else if count[rune('.')] == 2 {
					// VV: There might be a unique rune missing between the row and column arrays. This implies that there is
					// also 1 unique `?` symbol (the `.` symbol will appear twice)
					unique := []rune{}
					for k, v := range count {
						if v == 1 {
							unique = append(unique, k)
						}
					}

					if len(unique) == 2 {
						idx_question := -1
						idx_symbol := -1

						for i, what := range unique {
							if what == rune('?') {
								idx_question = i
							} else {
								idx_symbol = i
							}
						}

						if idx_question != -1 && idx_symbol != -1 {
							input.Grid[y+dy][x+dx] = unique[idx_symbol]
							input.Grid[loc_question.y][loc_question.x] = unique[idx_symbol]
						}
					}
				}
			}
		}

		if !updated {
			break
		}
	}

	return measure_rune(input, dx, dy)
}

type Point struct {
	x, y int
}

func solution(input *utilities.Puzzle) int {
	const block_size = 8
	ret := 0

	solved := map[Point]int{}

	for {
		updated := false

		for y := 0; y < (len(input.Grid)-2)/(block_size-2); y++ {
			for x := 0; x < (len(input.Grid[0])-2)/(block_size-2); x++ {
				if _, done := solved[Point{y: y, x: x}]; done {
					continue
				}

				block := decode_rune(input, x*(block_size-2), y*(block_size-2))

				if block > -1 {
					solved[Point{y: y, x: x}] = block
					updated = true
					ret += block
				} else if block == -2 {
					// VV: This is an unsolvable block, no need to re-visit it again
					solved[Point{y: y, x: x}] = -1
					updated = true
				}

			}
		}

		if !updated {
			break
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
