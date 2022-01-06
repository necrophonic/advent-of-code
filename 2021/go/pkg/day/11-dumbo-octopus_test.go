package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDumboOctopusAnswer(t *testing.T) {
	d := &day.DumboOctopus{}
	part1, part2, err := d.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 1713, part1)
	assert.Equal(t, 502, part2)
}

func TestFlash_100Step(t *testing.T) {
	d := &day.DumboOctopus{}

	data := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}

	total, _ := d.Flash(d.MakeGrid(data), 100)
	assert.Equal(t, 1656, total)
}

func TestFlash_allFlash(t *testing.T) {
	d := &day.DumboOctopus{}

	data := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}

	_, first := d.Flash(d.MakeGrid(data), 200)
	assert.Equal(t, 195, first)
}
