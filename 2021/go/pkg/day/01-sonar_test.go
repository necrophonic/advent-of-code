package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestSonarAnswer(t *testing.T) {
	part1, part2, err := day.Sonar{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 1602, part1)
	assert.Equal(t, 1633, part2)
}
