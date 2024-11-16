package main

import (
	"crypto/sha512"
	"encoding/hex"
	"os"
	"slices"
	"strconv"

	"ec-2024-1-koa/day5_c/utilities"
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

func hash(input *utilities.Puzzle, last_queue_index int) string {

	hasher := sha512.New()

	for _, queue := range input.Queues {
		for _, dancer := range queue {
			hasher.Write([]byte(strconv.Itoa(dancer)))
		}
	}
	hasher.Write([]byte(strconv.Itoa(last_queue_index)))

	return hex.EncodeToString(hasher.Sum(nil))
}

func solution(input *utilities.Puzzle) int {
	max := -1
	book := make(map[string]int, 0)

	// VV: Keep a record of all (Input, LastqueueIndex) tuples.
	// If you're about to dance the same way you've danced before, you know that
	// you're about to start repeating a dance. At that point just report the max dance number
	for round := 1; ; round++ {
		queue_index := (round - 1) % len(input.Queues)

		le_hash := hash(input, queue_index)

		if _, ok := book[le_hash]; ok {
			return max
		}

		number := dance_round(input, queue_index)

		if number > max {
			max = number
		}

		book[le_hash] = 1
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
