package day_3

import "testing"
import "github.com/stretchr/testify/assert"

func Test_start_1x1(t *testing.T) {
	x, y := middle(grid(1, 1))
	assert.Equal(t, 0, y)
	assert.Equal(t, 0, x)
}

func Test_start_1x2(t *testing.T) {
	x, y := middle(grid(2, 1))
	assert.Equal(t, 0, y)
	assert.Equal(t, 0, x)
}

func Test_start_2x2(t *testing.T) {
	x, y := middle(grid(2, 2))
	assert.Equal(t, 1, y)
	assert.Equal(t, 0, x)
}

func Test_start_3x2(t *testing.T) {
	x, y := middle(grid(3, 2))
	assert.Equal(t, 1, y)
	assert.Equal(t, 1, x)
}

func Test_start_3x3(t *testing.T) {
	x, y := middle(grid(3, 3))
	assert.Equal(t, 1, y)
	assert.Equal(t, 1, x)
}

func Test_start_5x5(t *testing.T) {
	x, y := middle(grid(5, 5))
	assert.Equal(t,2, y)
	assert.Equal(t, 2, x)
}

func grid(x, y int) [][]int {
	grid := [][]int{}
	for i := 0; i < y; i++ {
		grid = append(grid, make([]int, x))
	}
	return grid
}