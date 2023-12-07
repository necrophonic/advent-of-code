package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestSmokeBasinAnswer(t *testing.T) {
	sb := &day.SmokeBasin{}
	part1, part2, err := sb.Answer()
	assert.NoError(t, err)
	assert.Equal(t, -1, part1)
	assert.Equal(t, -1, part2)
}

func TestGetPoint(t *testing.T) {
	data := []string{
		"1234",
		"5658",
		"9082",
	}

	sb := &day.SmokeBasin{Data: data}

	assert.Equal(t, 1, sb.GetPoint(0, 0))
	assert.Equal(t, 6, sb.GetPoint(1, 1))
	assert.Equal(t, 9, sb.GetPoint(-1, 0))
	assert.Equal(t, 9, sb.GetPoint(4, 0))
	assert.Equal(t, 9, sb.GetPoint(0, -1))
	assert.Equal(t, 4, sb.GetPoint(0, 3))
	assert.Equal(t, 9, sb.GetPoint(3, 0))
	assert.Equal(t, 5, sb.GetPoint(1, 2))
}

func TestLowPoint(t *testing.T) {
	data := []string{
		"1294",
		"5658",
		"9082",
	}

	sb := &day.SmokeBasin{Data: data}

	assert.Equal(t, 6, sb.LowPoint(1, 2))
	assert.Equal(t, 0, sb.LowPoint(1, 1))
	assert.Equal(t, 2, sb.LowPoint(0, 0))
}
