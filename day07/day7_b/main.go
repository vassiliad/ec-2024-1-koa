package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"ec-2024-1-koa/day7_b/utilities"
)

type Score struct {
	Name        string
	Score       int
	Accumulated int
	H           []int
}

func apply_terrain_to_action(action, terrain string) string {
	if terrain == "=" || terrain == "S" {
		return action
	}

	return terrain
}

func solution(input *utilities.Puzzle) string {
	const loops = 10

	scores := []Score{}

	for _, plan := range input.Plans {
		scores = append(scores, Score{Name: plan.Name, Score: 10, H: []int{}})
	}

	idx_segment := 0

	for _ = range loops {
		for round := range input.Terrain {
			terrain := input.Terrain[(round+1)%len(input.Terrain)]

			for idx, plan := range input.Plans {
				idx_round := idx_segment % len(plan.Actions)
				action := apply_terrain_to_action(plan.Actions[idx_round], terrain)

				switch action {
				case "+":
					scores[idx].Score++
				case "-":
					scores[idx].Score--
				case "=":
				default:
					panic(fmt.Sprintf("unknown action %s for plan %s", action, plan.Name))
				}

				scores[idx].H = append(scores[idx].H, scores[idx].Score)
				scores[idx].Accumulated += scores[idx].Score
			}

			idx_segment++
		}
	}

	logger := utilities.SetupLogger()

	for _, s := range scores {
		// logger.Printf("%s: %d - %+v\n", s.Name, s.Accumulated, s.H)
		logger.Printf("%s: %d\n", s.Name, s.Accumulated)
	}

	sort.SliceStable(scores, func(i int, j int) bool {
		return scores[i].Accumulated > scores[j].Accumulated
	})

	sol := ""

	for _, s := range scores {
		sol += s.Name
	}

	return sol
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	logger.Println(strings.Join(input.Terrain, ""))

	if err != nil {
		logger.Fatalln("Ran into problems while reading input. Problem", err)
	}

	sol := solution(input)

	logger.Println("Solution is", sol)
}
