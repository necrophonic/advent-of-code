package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestLanternfishAnswer(t *testing.T) {
	f := &day.Lanternfish{}
	part1, part2, err := f.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 349549, part1)
	assert.Equal(t, 1589590444365, part2)
}

func TestLanternfish_Solver(t *testing.T) {
	data := []int{3, 4, 3, 1, 2}

	f := &day.Lanternfish{}

	answer, err := f.Solver(18, data)
	if assert.NoError(t, err) {
		assert.Equal(t, 26, answer)
	}

	answer, err = f.Solver(256, data)
	if assert.NoError(t, err) {
		assert.Equal(t, 26984457539, answer)
	}
}
