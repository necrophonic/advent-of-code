package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestWhalesAnswer(t *testing.T) {
	w := &day.Whales{}
	part1, part2, err := w.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 352707, part1)
	assert.Equal(t, 95519693, part2)
}

func TestWhales_Part1(t *testing.T) {
	data := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	w := &day.Whales{}
	ans, err := w.Part1(data)
	if assert.NoError(t, err) {
		assert.Equal(t, 37, ans)
	}
}

func TestWhales_Part2(t *testing.T) {
	data := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	w := &day.Whales{}
	ans, err := w.Part2(data)
	if assert.NoError(t, err) {
		assert.Equal(t, 168, ans)
	}
}

func TestWhales_CalcFuelToDistance(t *testing.T) {
	assert.Equal(t, 5, day.CalcFuelToDistance(10, []int{15}))
	assert.Equal(t, 5, day.CalcFuelToDistance(15, []int{10}))
}

func TestWhales_CalcCrabFuelToDistance(t *testing.T) {
	assert.Equal(t, 15, day.CalcCrabFuelToDistance(10, []int{15}))
	assert.Equal(t, 30, day.CalcCrabFuelToDistance(10, []int{15, 15}))
	assert.Equal(t, 1, day.CalcCrabFuelToDistance(10, []int{11}))
	assert.Equal(t, 16, day.CalcCrabFuelToDistance(10, []int{11, 15}))
	assert.Equal(t, 168, day.CalcCrabFuelToDistance(5, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}))
	assert.Equal(t, 206, day.CalcCrabFuelToDistance(2, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}))
	assert.Equal(t, 15, day.CalcCrabFuelToDistance(15, []int{10}))
}
