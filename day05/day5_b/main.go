package main

import (
	"os"
	"slices"

	"ec-2024-1-koa/day5_b/utilities"
)

// Returns the index in the queue for a dancer with @value, where the dancer
// starts dancing on the left side of the queue
func dance_left(value, queue_length int) int {
	loops := (value - 1) / queue_length
	end_up_left := (loops % 2) == 0

	if end_up_left {
		return (value - 1) % queue_length
	} else {
		delta_from_end := (value - 1) % queue_length
		return queue_length - delta_from_end
	}
}

func dance_round(input *utilities.Puzzle, queue_index int) int {
	next_queue := (queue_index + 1) % len(input.Queues)

	dancer := input.Queues[queue_index][0]
	input.Queues[queue_index] = input.Queues[queue_index][1:len(input.Queues[queue_index])]

	dancer_index := dance_left(dancer, len(input.Queues[next_queue]))

	input.Queues[next_queue] = slices.Insert(input.Queues[next_queue], dancer_index, dancer)

	number := 0
	power := 1
	for idx := range input.Queues {
		// VV: Support multi-digit dancers
		for remaining := input.Queues[len(input.Queues)-idx-1][0]; remaining > 0; remaining = remaining / 10 {
			number += (remaining % 10) * power
			power *= 10
		}

	}

	return number
}

func solution(input *utilities.Puzzle) int {

	const REPEATS = 2024
	book := make(map[int]int)

	for round := 1; ; round++ {
		queue_index := (round - 1) % len(input.Queues)

		number := dance_round(input, queue_index)

		if repeated, ok := book[number]; ok {
			if repeated == REPEATS-1 {
				return number * round
			}

			book[number] = repeated + 1
		} else {
			book[number] = 1
		}

	}

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
