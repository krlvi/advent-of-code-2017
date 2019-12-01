package day_5

import (
	"advent-of-code-2017/input"
	"fmt"
)

func Part1() {
	jumps := input.ToInts(input.Lines("day_5/in.txt"))
	steps := 0
	pos := 0
	for {
		if pos < 0 || pos >= len(jumps) {
			break
		}
		step := jumps[pos]
		jumps[pos]++
		pos = pos + step
		steps++
	}
	fmt.Println(steps)
}

func Part2() {
	jumps := input.ToInts(input.Lines("day_5/in.txt"))
	steps := 0
	pos := 0
	for {
		if pos < 0 || pos >= len(jumps) {
			break
		}
		step := jumps[pos]
		if step >= 3 {
			jumps[pos]--
		} else {
			jumps[pos]++
		}
		pos = pos + step
		steps++
	}
	fmt.Println(steps)
}
