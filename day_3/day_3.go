package day_3

import (
	"fmt"
	"math"
)

func Part_2() {
	input := 347991
	//input := 23
	grid := allocateGrid(input)
	fmt.Println(populateGridStress(grid, input))
}

func Part_1() {
	input := 347991
	grid := allocateGrid(input)
	startX, startY := middle(grid)
	finalX, finalY := populateGrid(grid, input)
	fmt.Println(manhattanDistance(startX, startY, finalX, finalY))
}

func populateGridStress(grid [][]int, limit int) (val int) {
	x, y := middle(grid)
	grid[y][x] = 1
	dir := "u"
	x++
	for i := 2; i <= limit; i++ {
		grid[y][x] = sumOfNeighbours(grid, x, y)
		if grid[y][x] > limit {
			val = grid[y][x]
			break
		}
		if dir == "r" {
			x++
			if y-1 >= 0 && x < len(grid[y-1]) && grid[y-1][x] == 0 {
				dir = "u"
			}
		} else if dir == "u" {
			y--
			if x-1 >= 0 && x >= 0 && grid[y][x-1] == 0 {
				dir = "l"
			}
		} else if dir == "l" {
			x--
			if y+1 < len(grid) && x >= 0 && grid[y+1][x] == 0 {
				dir = "d"
			}
		} else if dir == "d" {
			y++
			if x+1 < len(grid[y]) && grid[y][x+1] == 0 {
				dir = "r"
			}
		}
	}
	return
}

func manhattanDistance(fromX, fromY, toX, toY int) float64 {
	return math.Abs(float64(toX-fromX)) + math.Abs(float64(toY-fromY))
}

func middle(grid [][]int) (x, y int) {
	y = len(grid) / 2
	x = int(math.Ceil(float64(len(grid[y]))/2 - 1))
	return
}

func populateGrid(grid [][]int, limit int) (finalX, finalY int) {
	x, y := middle(grid)
	dir := "r"

	for i := 1; i <= limit; i++ {
		grid[y][x] = i
		if i == limit {
			finalX = x
			finalY = y
			break
		}
		if dir == "r" {
			x++
			if y-1 >= 0 && x < len(grid[y-1]) && grid[y-1][x] == 0 {
				dir = "u"
			}
		} else if dir == "u" {
			y--
			if x-1 >= 0 && x >= 0 && grid[y][x-1] == 0 {
				dir = "l"
			}
		} else if dir == "l" {
			x--
			if y+1 < len(grid) && x >= 0 && grid[y+1][x] == 0 {
				dir = "d"
			}
		} else if dir == "d" {
			y++
			if x+1 < len(grid[y]) && grid[y][x+1] == 0 {
				dir = "r"
			}
		}
	}
	return
}

func allocateGrid(limit int) [][]int {
	x := int(math.Ceil(math.Sqrt(float64(limit))))
	y := int(math.Round(math.Sqrt(float64(limit))))
	out := [][]int{}
	for i := 0; i < y; i++ {
		out = append(out, make([]int, x))
	}
	return out
}

func sumOfNeighbours(grid [][]int, x, y int) int {
	sum := 0
	if x+1 < len(grid[y]) {
		sum += grid[y][x+1]
	}
	if x-1 >= 0 {
		sum += grid[y][x-1]
	}
	if y+1 < len(grid) {
		sum += grid[y+1][x]
	}
	if y-1 >= 0 {
		sum += grid[y-1][x]
	}
	if x+1 < len(grid[y]) && y+1 < len(grid) {
		sum += grid[y+1][x+1]
	}
	if x+1 < len(grid[y]) && y-1 >= 0 {
		sum += grid[y-1][x+1]
	}
	if x-1 >= 0 && y-1 >= 0 {
		sum += grid[y-1][x-1]
	}
	if x-1 >= 0 && y+1 < len(grid) {
		sum += grid[y+1][x-1]
	}
	return sum
}
