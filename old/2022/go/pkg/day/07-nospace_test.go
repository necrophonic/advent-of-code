package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2022/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestNoSpace_Answer(t *testing.T) {
	part1, part2, err := day.NoSpace{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, -999, part1)
	assert.Equal(t, -999, part2)
}

func TestNoSpace_Part1(t *testing.T) {
	// d := day.NoSpace{}
}

func TestNoSpace_Part2(t *testing.T) {
	// d := day.NoSpace{}
}
