package main

import (
	"os"
	"slices"

	"ec-2024-1-koa/day5_a/utilities"
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

func solution(input *utilities.Puzzle) int {
	const rounds = 10

	for queue_index := range rounds {
		queue_index = queue_index % len(input.Queues)

		next_queue := (queue_index + 1) % len(input.Queues)

		// VV: Pop the first dancer from the queue
		dancer := input.Queues[queue_index][0]
		input.Queues[queue_index] = input.Queues[queue_index][1:len(input.Queues[queue_index])]

		dancer_index := dance_left(dancer, len(input.Queues[next_queue]))
		input.Queues[next_queue] = slices.Insert(input.Queues[next_queue], dancer_index, dancer)
	}

	sol := 0
	power := 1
	for idx := range input.Queues {
		sol += input.Queues[len(input.Queues)-idx-1][0] * power
		power *= 10
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
