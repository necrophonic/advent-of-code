package computer_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2019/pkg/computer"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	computer.DataFile = "../../data/computer.txt"
	p1, p2, err := computer.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 5482655, p1, "Part 1")
	assert.Equal(t, 4967, p2, "Part 2")
}

func TestRun(t *testing.T) {
	type tc struct {
		program  []int
		expected int
	}

	tcs := []tc{
		{[]int{1, 0, 0, 0, 99}, 2},
		{[]int{2, 3, 0, 3, 99}, 2},
		{[]int{2, 4, 4, 5, 99, 0}, 2},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 30},
	}

	for _, test := range tcs {
		output, err := computer.Run(test.program)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, output)
	}
}
