package main

import (
	"image"
	"os"
	"slices"
	"sort"

	"puzzle/utilities"
)

func distance(a, b image.Point) int {
	return utilities.Abs(a.X-b.X) + utilities.Abs(a.Y-b.Y)
}

func connect(stars map[int]image.Point, max_distance int) int {
	/*VV: This can be hella optimized but I can't be bothered to do that.

	Pretty much the same idea as part 1. Only it's a bit more annoying in that it only takes into account
	points whose distance is less than max_distance (i.e. 6). Once it cannot add any more points to the graph,
	it gives up and removes the points it consumed from @stars.

	Keep calling this method till it ends up consuming all the points in stars.
	*/
	connections := map[int]int{}

	connected := []int{}
	pending := map[int]int{}

	one := 0
	for i := range stars {
		one = i
		break
	}

	connected = append(connected, one)

	for i := range stars {
		if i != connected[0] {
			pending[i] = 1
		}
	}

	updated := true
	for len(pending) > 0 && updated {
		updated = false
		if len(connected) == 0 {

		}

		for k := range pending {
			pending[k] = int(^uint(0) >> 1)
		}

		for _, i := range connected {
			for j := range pending {
				d := distance(stars[i], stars[j])
				if pending[j] > d {
					connections[j] = i
					pending[j] = d
				}
			}
		}

		to_add := -1

		for j, d := range pending {
			if d < max_distance && (to_add == -1 || d < pending[to_add]) {
				to_add = j
			}
		}

		if to_add != -1 {
			updated = true
			connected = append(connected, to_add)
			delete(pending, to_add)
		}

	}

	sol := 0
	for i, j := range connections {
		if slices.Contains(connected, i) && slices.Contains(connected, j) {
			a := stars[i]
			b := stars[j]

			s := distance(a, b)
			sol += s
		}
	}

	for _, c := range connected {
		delete(stars, c)
	}

	return sol + len(connected)
}

func solution(input *utilities.Puzzle) int {
	/*VV: The idea here is to just brute force it without being any smart about it.
	Find all the possible clusters, sort them, and then return the product of the 3 largest ones.
	*/
	stars := map[int]image.Point{}

	for i, s := range input.Stars {
		stars[i] = s
	}

	sizes := []int{}

	for len(stars) > 0 {
		sizes = append(sizes, connect(stars, 6))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes[0] * sizes[1] * sizes[2]
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
