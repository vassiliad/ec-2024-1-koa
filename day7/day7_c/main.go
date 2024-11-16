package main

import (
	"fmt"
	"os"
	"strings"

	"ec-2024-1-koa/day7_c/utilities"
	"ec-2024-1-koa/day7_c/utilities/queue"
)

type Score struct {
	Name        string
	Score       int
	Accumulated int
}

func apply_terrain_to_action(action, terrain string) string {
	if terrain == "=" || terrain == "S" {
		return action
	}

	return terrain
}

func simulate(input *utilities.Puzzle) int {
	const loops = 2024

	scores := []Score{}

	for _, plan := range input.Plans {
		scores = append(scores, Score{Name: plan.Name, Score: 10})
	}

	idx_segment := 0

	// VV: The process repeats every LCM(track length, action plan) segments
	// there's no benefit in evaluating all 2024 loops - technically if
	// 2024 is not divisible by LCM *and* the 2 strategies are super close then there is a chance
	// that the extra few rounds make a difference. This worked for my input so I'm keeping it
	// as it speeds up the process by quite a bit :)
	max_steps := utilities.LCM(len(input.Terrain), len(input.Plans[0].Actions))

	for _ = range loops {
		for round := range input.Terrain {
			terrain := input.Terrain[round%len(input.Terrain)]

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

				scores[idx].Accumulated += scores[idx].Score
			}

			idx_segment++

			if idx_segment == max_steps {
				return scores[0].Accumulated
			}
		}
	}

	// VV: keeps compiler happy
	return scores[0].Accumulated
}

func digest_terrain(input *utilities.Puzzle) []int {
	stats := make([]int, 11)

	for i, c := range input.Terrain {
		d := 0

		if c == "=" || c == "S" {
			d = 1
		}
		stats[i%11] += d
	}

	return stats
}

func generate_plans() [][]string {
	ret := [][]string{}

	num_plus := 5
	num_minus := 3
	num_equal := 3

	type State struct {
		Plus, Minus, Equals int
		Value               []string
	}

	remaining := queue.New()

	remaining.Add(State{
		Plus: num_plus, Minus: num_minus, Equals: num_equal, Value: []string{},
	})

	for {
		cur := remaining.Pop()

		if cur == nil {
			break
		}

		val := cur.Value.(State)

		if val.Plus > 0 {
			n := State{Plus: val.Plus, Minus: val.Minus, Equals: val.Equals, Value: make([]string, len(val.Value))}
			copy(n.Value, val.Value)

			n.Plus--

			n.Value = append(n.Value, "+")

			remaining.Add(n)
		}

		if val.Minus > 0 {
			n := State{Plus: val.Plus, Minus: val.Minus, Equals: val.Equals, Value: make([]string, len(val.Value))}
			copy(n.Value, val.Value)

			n.Minus--

			n.Value = append(n.Value, "-")

			remaining.Add(n)
		}

		if val.Equals > 0 {
			n := State{Plus: val.Plus, Minus: val.Minus, Equals: val.Equals, Value: make([]string, len(val.Value))}
			copy(n.Value, val.Value)

			n.Equals--

			n.Value = append(n.Value, "=")
			remaining.Add(n)
		}

		if val.Equals+val.Minus+val.Plus == 0 {
			ret = append(ret, val.Value)
		}
	}

	return ret
}

func solution(input *utilities.Puzzle) int {

	all_plans := generate_plans()

	fmt.Println("Possible plans", len(all_plans))
	// fmt.Printf("%v\n", all_plans)

	// return 0
	to_beat := simulate(input)

	sol := 0

	for idx, candidate := range all_plans {

		input.Plans[0].Actions = candidate
		me := simulate(input)
		if idx%100 == 0 {
			strat := strings.Join(candidate, "")
			println("Evaluate", strat, 100*idx/len(all_plans), me, to_beat)
		}

		if me > to_beat {
			sol++
		}
	}

	return sol
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	// logger.Println(strings.Join(input.Terrain, ""))

	if err != nil {
		logger.Fatalln("Ran into problems while reading input. Problem", err)
	}

	sol := solution(input)

	logger.Println("Solution is", sol)
}
