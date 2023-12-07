package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2022/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestCampCleanup_Answer(t *testing.T) {
	part1, part2, err := day.CampCleanup{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 485, part1)
	assert.Equal(t, 857, part2)
}

func TestCampArea_Contains(t *testing.T) {
	cc := day.CampCleanup{}
	cr1, err := cc.NewCampArea("2-8")
	assert.NoError(t, err)
	cr2, err := cc.NewCampArea("3-7")
	assert.NoError(t, err)

	assert.True(t, cr1.Contains(cr2))
	assert.False(t, cr2.Contains(cr1))
}

func TestCampArea_Overlaps(t *testing.T) {
	cc := day.CampCleanup{}
	cr1, err := cc.NewCampArea("2-8")
	assert.NoError(t, err)
	cr2, err := cc.NewCampArea("3-10")
	assert.NoError(t, err)

	assert.True(t, cr1.Overlaps(cr2))
	assert.True(t, cr2.Overlaps(cr1))
}

func TestCampCleanup_Part1(t *testing.T) {
	cc := day.CampCleanup{}
	ans, err := cc.Part1([]string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, ans)
}

func TestCampCleanup_Part2(t *testing.T) {
	cc := day.CampCleanup{}
	ans, err := cc.Part2([]string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	})
	assert.NoError(t, err)
	assert.Equal(t, 4, ans)
}
