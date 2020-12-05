package boarding_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/boarding"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	boarding.DataFile = "../../data/boarding.txt"
	p1, p2, err := boarding.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 996, p1, "Part 1")
	assert.Equal(t, 671, p2, "Part 2")
}

func TestFindSeat(t *testing.T) {
	type tc struct {
		directions string
		row        int
		col        int
		seat       int
	}

	tcs := []tc{
		{"FBFBBFFRLR", 44, 5, 357},
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	for _, test := range tcs {
		r, c, s := boarding.FindSeat(test.directions)
		assert.Equal(t, test.col, c, "Column")
		assert.Equal(t, test.row, r, "Row")
		assert.Equal(t, test.seat, s, "Seat")
	}
}
