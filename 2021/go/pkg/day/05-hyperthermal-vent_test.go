package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestHyerthermalVentAnswer(t *testing.T) {
	h := &day.Hyper{}
	part1, part2, err := h.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 7468, part1)
	assert.Equal(t, 22364, part2)
}

var thermalTestData = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
}

func TestHyerthermalVent_Part1(t *testing.T) {
	h := &day.Hyper{}
	vents, err := h.ParseVents(thermalTestData)
	if assert.NoError(t, err) {
		ans, err := h.Part1(vents)
		if assert.NoError(t, err) {
			assert.Equal(t, 5, ans)
		}
	}
}

func TestHyerthermalVent_Part2(t *testing.T) {
	h := &day.Hyper{}
	vents, err := h.ParseVents(thermalTestData)
	if assert.NoError(t, err) {
		ans, err := h.Part2(vents)
		if assert.NoError(t, err) {
			assert.Equal(t, 12, ans)
		}
	}
}
