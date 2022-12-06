package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2022/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestRucksack_Answer(t *testing.T) {
	part1, part2, err := day.Rucksack{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 8139, part1)
	assert.Equal(t, 2668, part2)
}

func TestRucksack_Part1(t *testing.T) {
	r := day.Rucksack{}
	assert.Equal(t,
		157,
		r.Part1([]string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		}),
	)
}

func TestRucksack_Part2(t *testing.T) {
	r := day.Rucksack{}
	assert.Equal(t,
		70,
		r.Part2([]string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		}),
	)
}

func TestRucksack_SplitCompartments(t *testing.T) {
	r := day.Rucksack{}
	c1, c2 := r.SplitCompartments("vJrwpWtwJgWrhcsFMMfFFhFp")
	assert.Equal(t, c1, "vJrwpWtwJgWr")
	assert.Equal(t, c2, "hcsFMMfFFhFp")
}

func TestRucksack_FindDuplicates(t *testing.T) {
	r := day.Rucksack{}
	assert.Equal(t, "p", r.FindDuplicates(r.SplitCompartments("vJrwpWtwJgWrhcsFMMfFFhFp")))
	assert.Equal(t, "v", r.FindDuplicates(r.SplitCompartments("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn")))
	assert.Equal(t, "ab", r.FindDuplicates(r.SplitCompartments("abab")))

	assert.Equal(t, "ab", r.FindDuplicates("ab", "ab"))
	assert.Equal(t, "ab", r.FindDuplicates("ab", "abx"))
	assert.Equal(t, "a", r.FindDuplicates("ab", "abx", "fxa"))
}

func TestRucksack_ScoreDuplicates(t *testing.T) {
	r := day.Rucksack{}
	assert.Equal(t, 157, r.ScoreDuplicates("pLPvts"))
}

func TestRucksack_ScoreDupe(t *testing.T) {
	r := day.Rucksack{}
	assert.Equal(t, 1, r.ScoreDupe('a'))
	assert.Equal(t, 26, r.ScoreDupe('z'))
	assert.Equal(t, 27, r.ScoreDupe('A'))
	assert.Equal(t, 52, r.ScoreDupe('Z'))
}

func TestRucksack_DedupeCompartment(t *testing.T) {
	r := day.Rucksack{}
	assert.Equal(t, "abfcde", r.DedupeCompartment("aabbfcfffde"))
}
