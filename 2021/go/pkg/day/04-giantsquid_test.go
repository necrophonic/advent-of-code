package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestGiantSquidAnswer(t *testing.T) {
	gs := &day.GiantSquid{}
	part1, part2, err := gs.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 69579, part1)
	assert.Equal(t, 14877, part2)
}

func TestGiantSquid_New(t *testing.T) {
	rows := []string{
		"10 11 12 13 14",
		"15 16 17 18 19",
		"20 21  2 23 24",
		"25 26 27 28 29",
		" 3 31 32 33 34",
	}

	bb, err := day.NewBingoBoard(rows)
	if assert.NoError(t, err) {
		assert.Equal(t, day.BingoBoard{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 2, 23, 24, 25, 26, 27, 28, 29, 3, 31, 32, 33, 34}, bb)

		// Get rows
		assert.Equal(t, day.BingoLine{15, 16, 17, 18, 19}, bb.Row(1))
		assert.Equal(t, day.BingoLine{3, 31, 32, 33, 34}, bb.Row(4))

		// Get columns
		assert.Equal(t, day.BingoLine{10, 15, 20, 25, 3}, bb.Col(0))
		assert.Equal(t, day.BingoLine{14, 19, 24, 29, 34}, bb.Col(4))
	}

}

func TestGiantSquid_Win(t *testing.T) {
	rows := []string{
		"10  0 12 13 14",
		" 0  0  0  0  0",
		"20  0  2 23 24",
		"25  0 27 28 29",
		" 3  0 32 33 34",
	}

	bb, err := day.NewBingoBoard(rows)
	if assert.NoError(t, err) {
		// Row 1 and col 1 should be winners
		assert.True(t, bb.Row(1).IsWinner())
		assert.True(t, bb.Col(1).IsWinner())

		// Row 3 and col 3 are not winners
		assert.False(t, bb.Row(3).IsWinner())
		assert.False(t, bb.Col(3).IsWinner())

		// Board should be a winner
		assert.True(t, bb.IsWinner())
	}
}

func TestGiantSquid_Part1(t *testing.T) {
	data := []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}

	gs := &day.GiantSquid{}
	numbers, boards, err := gs.Parse(data)
	if assert.NoError(t, err) {
		answer, err := gs.Part1(numbers, boards)
		if assert.NoError(t, err) {
			assert.Equal(t, 4512, answer)
		}
	}
}

func TestGiantSquid_Part2(t *testing.T) {
	data := []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}

	gs := &day.GiantSquid{}
	numbers, boards, err := gs.Parse(data)
	if assert.NoError(t, err) {
		answer, err := gs.Part2(numbers, boards)
		if assert.NoError(t, err) {
			assert.Equal(t, 1924, answer)
		}
	}
}
