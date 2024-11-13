package main

import (
	"fmt"
	"os"
	"sort"

	"ec-2024-1-koa/day7_a/utilities"
)

type Score struct {
	Name        string
	Score       int
	Accumulated int
}

func solution(input *utilities.Puzzle) string {
	const segments = 10

	scores := []Score{}

	for _, plan := range input.Plans {
		scores = append(scores, Score{Name: plan.Name, Score: 0})
	}

	for round := range segments {

		for idx, plan := range input.Plans {
			idx_round := round % len(plan.Actions)

			switch action := plan.Actions[idx_round]; action {
			case "+":
				scores[idx].Score++
			case "-":
				scores[idx].Score--
			case "=":
			default:
				panic(fmt.Sprintf("unknown action %s for plan %s", action, plan.Name))
			}

			scores[idx].Accumulated += scores[idx].Score
		}
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

	// logger.Println("Input was", input)

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}

	sol := solution(input)

	logger.Println("Solution is", sol)
}
