package main

import (
	"ec-2024-1-koa/day2_c/utilities"
	"os"
	"strings"
)

func index(line string, substr string, start int) int {
	for i := start; i <= len(line); i++ {
		match := true

		for j := 0; j < len(substr); j++ {
			if line[(i+j+len(line))%len(line)] != substr[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}

func mark_runic_symbol_indices_horizontal(s, substr string, hit_indices []int) {

	true_len := len(s)

	// VV: Lines wrap horizontally. Here, I manually wrap the first N-1 characters of a line
	// where N is the length of the word we're searching. There's no need to go further
	// than that. The least amount of characters we need to check from the beginning of the line
	// is 1. The most, is N-1 i.e. we use the last character of the line and the remaining ones
	// from its beginning.
	expanded := strings.Join([]string{s, s[0 : len(substr)-1]}, "")
	s = expanded

	n := 0
	for {
		i := strings.Index(s, substr)

		if i == -1 {
			return
		}

		for j := range len(substr) {
			hit_indices[(n+i+j)%true_len] = 1
		}

		// VV: This is to account for words that overlap with themselves
		n += i + 1
		s = s[i+1:]
	}
}

func vertical_index(lines []string, substr string, column int) int {
	for i := 0; i <= len(lines)-len(substr); i++ {
		match := true

		for j := 0; j < len(substr); j++ {
			if lines[i+j][column] != substr[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}

func mark_runic_symbol_indices_vertical(lines []string, substr string, hit_indices [][]int, column int) {
	n := 0

	for {
		i := vertical_index(lines[n:], substr, column)

		if i == -1 {
			return
		}

		for j := range len(substr) {
			hit_indices[n+i+j][column] = 1
		}

		// VV: This is to account for words that overlap with themselves
		n += i + 1
	}
}

// Modified func from strings.Count
func mark_runic_symbol_indices(lines []string, substr string, hit_indices [][]int) {
	for column := range len(lines[0]) {
		mark_runic_symbol_indices_vertical(lines, substr, hit_indices, column)
	}

	for line_idx, s := range lines {
		// VV: Need to expand each line so that we the last character comes right before the 1st character
		mark_runic_symbol_indices_horizontal(s, substr, hit_indices[line_idx])
	}
}

func solution(input *utilities.Puzzle) int {
	sol := 0
	hit_indices := make([][]int, len(input.Text))
	for line_idx, text := range input.Text {
		hit_indices[line_idx] = make([]int, len(text))
	}

	for _, word := range input.Words {
		mark_runic_symbol_indices(input.Text, word, hit_indices)

		score := 0
		for _, hit_indices_line := range hit_indices {
			for _, hit := range hit_indices_line {
				score += hit
			}

		}
	}

	score := 0
	for _, hit_indices_line := range hit_indices {
		for _, hit := range hit_indices_line {
			score += hit
		}

	}

	sol += score

	// for row, text := range input.Text {
	// 	for column, ch := range text {
	// 		hit := hit_indices[row][column]
	// 		if hit == 1 {
	// 			fmt.Printf("%c", ch)
	// 		} else {
	// 			fmt.Printf("%c", unicode.ToLower(ch))
	// 		}
	// 	}
	// 	fmt.Println()
	// }

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
