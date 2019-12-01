package day_1

import (
	"advent-of-code-2017/input"
	"fmt"
	"strconv"
)

func Part2() {
	lines := input.Lines("day_1/in.txt")
	digits := []int{}
	for _, r := range lines[0] {
		d, _ := strconv.Atoi(string(r))
		digits = append(digits, d)
	}

	sum := 0
	for i, d := range digits {
		half := len(digits)/2
		var nextVal int
		if i + half < len(digits) {
			nextVal = digits[i+half]
		} else {
			nextVal = digits[i + half - len(digits)]
		}

		if d == nextVal {
			sum += d
		}
	}
	fmt.Println(sum)
}

func Part1() {
	lines := input.Lines("day_1/in.txt")
	digits := []int{}
	for _, r := range lines[0] {
		d, _ := strconv.Atoi(string(r))
		digits = append(digits, d)
	}
	sum := 0
	for i, d := range digits {
		var nextVal int
		if i+1 < len(digits) {
			nextVal = digits[i+1]
		} else {
			nextVal = digits[i]
		}
		if d == nextVal {
			sum += d
		}
	}
	fmt.Println(sum)
}
