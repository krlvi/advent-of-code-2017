package day_6

import (
	"advent-of-code-2017/input"
	"fmt"
	"strings"
)

func Part2() {
	in := in()
	seen := map[string]bool{}
	cycles := 0

	count := false
	counter := 0
	state := ""
	for {
		str := fmt.Sprint(redist(in))
		cycles++
		if count {
			counter++
		}
		if seen[str] && count == false {
			state = str
			count = true
		} else if str == state && count == true {
			break
		}
		seen[str] = true
	}
	fmt.Println(counter)
}

func Part1() {
	in := in()
	seen := map[string]bool{}
	cycles := 0

	for {
		str := fmt.Sprint(redist(in))
		cycles++
		if seen[str] {
			break
		}
		seen[str] = true
	}
	fmt.Println(cycles)
}

func redist(in []int) []int {
	i, valToRedist := max(in)
	in[i] = 0
	for {
		if valToRedist == 0 {
			break
		}
		if i+1 >= len(in) {
			i = 0
		} else {
			i++
		}
		in[i]++
		valToRedist--
	}
	return in
}

func max(in []int) (idx, val int) {
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] >= val {
			val = in[i]
			idx = i
		}
	}
	return
}

func in() []int {
	lines := input.Lines("day_6/in.txt")
	return input.ToInts(strings.Fields(lines[0]))
}
