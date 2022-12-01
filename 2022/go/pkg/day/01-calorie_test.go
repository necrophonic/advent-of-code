package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2022/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestCalorie_Answer(t *testing.T) {
	part1, part2, err := day.Calorie{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 74711, part1)
	assert.Equal(t, 209481, part2)
}
