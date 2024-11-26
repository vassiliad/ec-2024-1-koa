package main

import (
	"container/list"
	"fmt"
	"os"

	"puzzle/utilities"
)

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

type State struct {
	times_right int
	offset      int
	score       int
}

func (p State) spin_with_offset(offset int, input *utilities.Puzzle, wheel_idx []int) State {
	p.offset += offset
	p.times_right++

	for i, spin := range input.Spins {
		// VV: offset % N can be negative (min number is -(N-1)) so add N
		offset := (p.offset % len(input.Wheels[i])) + len(input.Wheels[i])
		wheel_idx[i] = (offset + spin*p.times_right) % len(input.Wheels[i])
	}

	s := count_score(wheel_idx, input)

	p.score += s
	return p
}

func solution(input *utilities.Puzzle) string {
	/*VV: basically brute force it while only visiting the same state once

	Actually, the fact that this works implies that pulling the lever and then pushing it
	is the same as not touching the lever at all.

	We can definitely use T to calculate T+1.

	That would significantly reduce the execution time of this solution which is not particularly great :)
	*/
	pending := list.New()
	pending.PushBack(
		State{
			times_right: 0, offset: 0, score: 0,
		},
	)

	seen := map[State]int{}
	max_score := -1
	min_score := int(^uint(0) >> 1)
	wheel_idx := make([]int, len(input.Wheels))

	max_rounds := 256
	steps := 0
	for pending.Len() > 0 {
		cur := pending.Remove(pending.Front()).(State)

		steps++

		// VV: slow enough to warrant periodic prints!
		if steps%100000 == 0 {
			fmt.Printf("%d) %+v\n", steps, cur)
		}

		/*
			VV: Three ways to do a spin:
			1. Do not mess with the left lever (no offset)
			2. Pull the left lever (i.e. move wheels forward) (offset +1)
			3. Push the left lever (i.e. move wheels backward) (offset -1)

			Perhaps we can use knowledge about move at time T to make a decision for move at time T+1
		*/

		for offset := -1; offset < 2; offset++ {
			next := cur.spin_with_offset(offset, input, wheel_idx)

			if _, ok := seen[next]; ok {
				continue
			}
			seen[next] = 1

			if next.times_right == max_rounds {
				max_score = max(max_score, next.score)
				min_score = min(min_score, next.score)

				continue
			}

			pending.PushBack(next)
		}

	}

	return fmt.Sprintf("%d %d", max_score, min_score)
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
