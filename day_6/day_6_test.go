package day_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_redist_1(t *testing.T) {
	assert.EqualValues(t, []int{2, 4, 1, 2}, redist([]int{0, 2, 7, 0}))
}

func Test_redist_2(t *testing.T) {
	assert.EqualValues(t, []int{3, 1, 2, 3}, redist([]int{2, 4, 1, 2}))
}

func Test_redist_3(t *testing.T) {
	assert.EqualValues(t, []int{0, 2, 3, 4}, redist([]int{3, 1, 2, 3}))
}

func Test_redist_4(t *testing.T) {
	assert.EqualValues(t, []int{1, 3, 4, 1}, redist([]int{0, 2, 3, 4}))
}

func Test_redist_5(t *testing.T) {
	assert.EqualValues(t, []int{2, 4, 1, 2}, redist([]int{1, 3, 4, 1}))
}

func Test_max_when_two(t *testing.T) {
	idx, val := max([]int{3, 1, 2, 3})
	assert.Equal(t, 0, idx)
	assert.Equal(t, 3, val)
}
