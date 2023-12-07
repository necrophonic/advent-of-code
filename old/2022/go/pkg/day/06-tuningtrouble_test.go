package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2022/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestTuning_Answer(t *testing.T) {
	part1, part2, err := day.Tuning{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 1892, part1)
	assert.Equal(t, 2313, part2)
}

func TestTuning_Part1(t *testing.T) {
	d := day.Tuning{}
	assert.Equal(t, 5, d.Part1("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 6, d.Part1("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 10, d.Part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 11, d.Part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestTuning_Part2(t *testing.T) {
	d := day.Tuning{}
	assert.Equal(t, 19, d.Part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 23, d.Part2("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 23, d.Part2("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 29, d.Part2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 26, d.Part2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestTuning_IsUnique(t *testing.T) {
	d := day.Tuning{}
	assert.True(t, d.IsUnique("abcd"), "abcd = true")
	assert.False(t, d.IsUnique("abcb"), "abcb = false")
}
