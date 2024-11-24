package main

import (
	"os"

	"puzzle/utilities"
)

func simulate(y, power int) []int {
	// VV: Contains all Y points starting from time 0 (BEFORE THE SHOT IS MADE)
	ret := []int{y}

	for range power {
		y++
		ret = append(ret, y)
	}

	for range power {
		ret = append(ret, y)
	}

	y--

	for ; y >= 0; y-- {
		ret = append(ret, y)
	}

	return ret
}

func solution(input *utilities.Puzzle) int {
	/*VV: Find all possible positions of projectiles, then simulate meteors falling down.
	For each meteor, start the simulation from time = 0, and then waiting for 1 second in each subsequent simulation.

	Report the cost with the minimum score for each meteor, aggregate these numbers and that's your solution.
	*/
	sol := 0
	max_power := max(input.Height, input.Width)

	// VV: Format is : {segmentId: {powerLevel: [y at t =0, y at t=1, ...., 0 (i.e. projectile collides with the ground)]}}
	cache := map[int]map[int][]int{}

	for i := 1; i < 4; i++ {
		seg := input.Segments[i]

		cache[i] = make(map[int][]int)

		for power := range max_power {
			cache[i][power] = simulate(seg.Y, power)
		}
	}

	for _, target := range input.Targets {
		cost := int(^uint(0) >> 1)
		// VV: We might need to wait for a bit

		for wait := 0; wait < target.X/2; wait++ {
			for i := 1; i < 4; i++ {
				t := target

				t.X -= wait
				t.Y -= wait

				if t.X%2 != 0 {
					// VV: The projectile would collide with the meteor at a non discrete time
					continue
				}

				// VV: TODO This is the bit that can be further optimized. For example there's no way
				// for a projectile with power 5 to collide with a meteor that spawns at 1000, 1000.
				// Also I'm sure there's an analytical way to find this score ...
				for power := 1; power < max_power; power++ {
					for col_t, col_y := range cache[i][power] {
						mx := t.X - col_t
						my := t.Y - col_t

						if mx < col_t {
							// VV: The projectile went past the meteor!
							break
						}

						if col_t == mx && col_y == my {
							cost = min(cost, i*power)
							break
						}
					}
				}

				if cost != int(^uint(0)>>1) {
					break
				}

			}

			if cost != int(^uint(0)>>1) {
				break
			}
		}

		sol += cost
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
