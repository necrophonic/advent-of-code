package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDiveAnswer(t *testing.T) {
	part1, part2, err := day.Dive{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 1660158, part1)
	assert.Equal(t, 1604592846, part2)
}
