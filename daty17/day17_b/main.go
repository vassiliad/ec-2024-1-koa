package main

import (
	"image"
	"os"

	"puzzle/utilities"
)

func distance(a, b image.Point) int {
	return utilities.Abs(a.X-b.X) + utilities.Abs(a.Y-b.Y)
}

func connect(input *utilities.Puzzle, connections map[int]int) {
	/*Here we're basically building a graph that spans across all points.

	I could have looked at some well known method to do this but this was more fun.

	So, starting from a random star (e.g. 1st one) build a graph by adding one star at a time.
	Always add the star that's closest to the graph. Whenever you add a star don't consider it again.

	Eventually the graph will contain all nodes
	*/
	connected := []int{0}
	pending := map[int]int{}

	for i := 1; i < len(input.Stars); i++ {
		pending[i] = 1
	}

	for len(pending) > 0 {
		for k := range pending {
			pending[k] = int(^uint(0) >> 1)
		}

		for _, i := range connected {
			for j := range pending {
				d := distance(input.Stars[i], input.Stars[j])
				if pending[j] > d {
					connections[j] = i
					pending[j] = d
				}
			}
		}

		to_add := -1

		for j, d := range pending {
			if to_add == -1 || d < pending[to_add] {
				to_add = j
			}
		}

		connected = append(connected, to_add)
		delete(pending, to_add)
	}
}

func solution(input *utilities.Puzzle) int {
	connections := map[int]int{}

	connect(input, connections)

	sol := len(input.Stars)

	for i, j := range connections {
		sol += distance(input.Stars[i], input.Stars[j])
	}

	return sol
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
