package main

import (
	"fmt"
	"os"
	"slices"

	"puzzle/utilities"
)

func solution(input *utilities.Puzzle) int {
	/* VV: I actually solved this using a calculator so this may not work for all puzzle inputs.
	My input had a couple columns which only contained `+` and `.` symbols. This means that you can get to those
	columns and just keep gliding towards the south till you safely land.
	*/

	// VV: For each column record a positive number if it only contains '.' (or 'S') and '+' symbols
	// the key is the column index, the value is the total cost when visiting all the points in the column
	// actually, we can include columns with '-' symbols too. The problematic symbols are '#' cause they
	// force us to move out of the column

	columns := []int{}

	for x := range len(input.Board[0]) {
		cost := 0
		for _, row := range input.Board {
			if row[x] == '#' {
				cost = -1
				break
			}
			// VV: Note that we're recording the cost here so, a '+' incures "negative" cost
			switch row[x] {
			case '.', 'S':
				cost++
			case '-':
				cost += 2
			case '+':
				cost -= 1
			}
		}

		columns = append(columns, cost)
	}

	start_x := slices.Index(input.Board[0], 'S')

	if slices.Contains(input.Board[0][1:len(input.Board)-1], '#') {
		panic("Solution assumes that there's a path to the beginning of any column")
	}

	max_distance := -1
	const starting_height = 384400

	for x, cost := range columns {
		if cost == -1 || x == 0 || x == len(columns)-1 {
			continue
		}

		// VV: We need to get right before the start of the column
		dx := max(x-start_x, start_x-x) - 1

		// VV: First move left/right to the start of the column
		height := (starting_height - dx)

		blocks := height / cost
		// VV: subtract 1 to account for the first row in the very first block
		distance := blocks*len(input.Board) - 1

		height -= blocks * cost

		// VV: just simulate the last few steps in the last (and incomplete) block
		rem_distance := 0
		for y := 0; height > 0; y++ {
			rem_distance++

			switch input.Board[y][x] {
			case '.', 'S':
				height--
			case '+':
				height++
			case '-':
				height -= 2
			default:
				panic(fmt.Sprintf("Unexpected position %d,%d = %s", x, y, string(input.Board[y][x])))
			}
		}

		max_distance = max(max_distance, distance+rem_distance)
	}

	if max_distance < 0 {
		panic("Oh no")
	}

	return max_distance
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
