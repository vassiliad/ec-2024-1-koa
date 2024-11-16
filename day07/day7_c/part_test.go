package main

import (
	"reflect"
	"strings"
	"testing"

	"ec-2024-1-koa/day7_c/utilities"
	"ec-2024-1-koa/day7_c/utilities/queue"
)

func TestSmall(t *testing.T) {
	small := `A:+,-,=,=

S+===
-   +
=+=-+
`

	input, _ := utilities.ReadString(small)

	// t.Logf("Input was %+v", input)

	expected_terrain := []string{"+", "=", "=", "=", "+", "+", "-", "=", "+", "=", "-", "S"}

	if !reflect.DeepEqual(input.Terrain, expected_terrain) {
		t.Fatal("Expected terrain to be", expected_terrain, "but it was", input.Terrain)
	}
}

func TestGCD(t *testing.T) {
	x := utilities.GCD(1071, 462)
	if x != 21 {
		t.Fatal("GCD(1071, 462)=21 but got", x)
	}

	x = utilities.GCD(48, 18)
	if x != 6 {
		t.Fatal("GCD(48, 18)=6 but got", x)
	}

	x = utilities.GCD(462, 1071)
	if x != 21 {
		t.Fatal("GCD(462, 1071)=21 but got", x)
	}

	x = utilities.GCD(18, 48)
	if x != 6 {
		t.Fatal("GCD(18, 48)=6 but got", x)
	}

	x = utilities.GCD(340, 11)

	if x != 1 {
		t.Fatal("GCD(340, 11)=1 but got", x)
	}
}

func TestLCM(t *testing.T) {
	x := utilities.LCM(6, 7)
	if x != 42 {
		t.Fatal("LCM(6, 7)=42 but got", x)
	}

	x = utilities.LCM(340, 11)

	if x != 3740 {
		t.Fatal("LCM(340, 11)=3740 but got", x)
	}
}

func TestCheckInput(t *testing.T) {
	input, _ := utilities.ReadInputFile("input.txt")

	stats := make([]int, 11)

	for i, c := range input.Terrain {
		d := 0

		if c == "=" || c == "S" {
			d = 1
		}
		stats[i%11] += d
	}

	// fmt.Printf("%+v\n", stats)
}

func TestGeneratePlans(t *testing.T) {
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

	cur := remaining.Front()

	generated := map[string]int{}

	for {
		cur = remaining.Pop()

		if cur == nil {
			break
		}

		val := cur.Value.(State)

		action := strings.Join(val.Value, "")
		// fmt.Printf("%+v\n", val)
		if _, ok := generated[action]; ok {
			cur = remaining.Front()
			continue
		}

		// fmt.Printf("%+v\n", val)

		if val.Plus > 0 {
			n := State(val)

			n.Plus--

			n.Value = append(n.Value, "+")

			action := strings.Join(n.Value, "")

			if _, ok := generated[action]; !ok {
				remaining.Add(n)
			}
		}

		if val.Minus > 0 {
			n := State(val)

			n.Minus--

			n.Value = append(n.Value, "-")

			action := strings.Join(n.Value, "")

			if _, ok := generated[action]; !ok {
				remaining.Add(n)
			}
		}

		if val.Equals > 0 {
			n := State(val)

			n.Equals--

			n.Value = append(n.Value, "=")

			action := strings.Join(n.Value, "")

			if _, ok := generated[action]; !ok {
				remaining.Add(n)
			}
		}

		if val.Equals+val.Minus+val.Plus == 0 {
			action := strings.Join(val.Value, "")

			if _, ok := generated[action]; !ok {
				// println(action)
			}
		}

		generated[action] = 1
	}
}
