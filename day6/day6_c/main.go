package main

import (
	"container/list"
	"fmt"
	"os"
	"slices"

	"ec-2024-1-koa/day6_c/utilities"
)

type Point struct {
	CurrentNode string
	Distance    int
	Path        string
}

func trim_tree(input *utilities.Puzzle) {
	nodes_to_keep := map[string]int{"RR": 1, "@": 1}

	for node, arr := range input.Branches {
		if node == "BUG" || node == "ANT" {
			continue
		}

		if slices.Contains(arr, "@") {
			nodes_to_keep[node] = 1
		}
	}

	for {
		starting_nodes := len(nodes_to_keep)

		for name, branches := range input.Branches {
			if name == "BUG" || name == "ANT" {
				continue
			}

			for node, _ := range nodes_to_keep {
				if slices.Contains(branches, node) {
					nodes_to_keep[name] = 1
				}
			}
		}

		if starting_nodes == len(nodes_to_keep) {
			break
		}
	}

	new_branches := map[string][]string{}

	for node, _ := range nodes_to_keep {
		leads_to := []string{}
		for _, dest := range input.Branches[node] {
			if _, ok := nodes_to_keep[dest]; ok {
				leads_to = append(leads_to, dest)
			}
		}
		new_branches[node] = leads_to
	}

	input.Branches = new_branches
}

func solution(input *utilities.Puzzle) string {
	// VV: This step is completely unecessary
	trim_tree(input)

	remaining := list.New()

	fruit_distances := map[int][]Point{}

	remaining.PushFront(Point{CurrentNode: "RR", Distance: 0, Path: "R"})

	cur := remaining.Front()

	total_fruits := 0

	for {
		if cur == nil {
			break
		}

		branch := cur.Value.(Point)

		if _, ok := input.Branches[branch.CurrentNode]; !ok {
			panic(branch.CurrentNode)
		}

		for _, n := range input.Branches[branch.CurrentNode] {
			if n == "BUG" || n == "ANT" {
				continue
			}
			path := fmt.Sprintf("%s%c", branch.Path, n[0])
			child := Point{
				Distance:    branch.Distance + 1,
				Path:        path,
				CurrentNode: n,
			}

			if n == "@" {
				total_fruits += 1
				if others, ok := fruit_distances[child.Distance]; ok {
					fruit_distances[child.Distance] = append(others, child)
				} else {
					fruit_distances[child.Distance] = []Point{child}
				}
			} else {
				remaining.PushBack(child)
			}

		}

		cur = cur.Next()
	}

	for _, fruits := range fruit_distances {
		if len(fruits) == 1 {
			return fruits[0].Path
		}
	}

	panic("Should never happen")
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
