package main

import (
	"reflect"
	"testing"

	"ec-2024-1-koa/day7_b/utilities"
)

func TestSmall(t *testing.T) {
	small := `A:+,-,=,=
B:+,=,-,+
C:=,-,+,+
D:=,=,=,+

S+===
-   +
=+=-+
`

	input, err := utilities.ReadString(small)

	// t.Logf("Input was %+v", input)

	expected_terrain := []string{"S", "+", "=", "=", "=", "+", "+", "-", "=", "+", "=", "-"}

	if !reflect.DeepEqual(input.Terrain, expected_terrain) {
		t.Fatal("Expected terrain to be", expected_terrain, "but it was", input.Terrain)
	}

	if err != nil {
		t.Fatal("Ran into problems while reading input. Problem", err)
	}

	solution := solution(input)

	const correct_answer = "DCBA"

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
