package day_2

import (
	"advent-of-code-2017/input"
	"fmt"
	"strconv"
	"strings"
)

func findDvisable(v []int) (div int) {
	for i, x := range v {
		for j := i + 1; j < len(v); j++ {
			y := v[j]
			if x%y == 0 {
				div = x / y
				return
			} else if y%x == 0 {
				div = y / x
				return
			}
		}
	}
	return 0
}

func Part2() {
	matrix := getMatrix()
	sum := 0
	for _, v := range matrix {
		sum += findDvisable(v)
	}
	fmt.Println(sum)
}

func Part1() {
	matrix := getMatrix()
	sum := 0
	for _, v := range matrix {
		min := v[0]
		max := v[0]
		for _, d := range v {
			if d < min {
				min = d
			}
			if d > max {
				max = d
			}
		}
		sum += max - min
	}
	fmt.Println(sum)
}

func getMatrix() [][]int {
	lines := input.Lines("day_2/in.txt")
	matrix := [][]int{}
	for _, l := range lines {
		fields := strings.Fields(l)
		values := []int{}
		for _, i := range fields {
			v, _ := strconv.Atoi(i)
			values = append(values, v)
		}
		matrix = append(matrix, values)
	}
	return matrix
}
