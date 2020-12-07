package wires_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2019/pkg/wires"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	wires.DataFile = "../../data/wires.txt"
	p1, p2, err := wires.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 1264, p1, "Part 1")
	assert.Equal(t, -1, p2, "Part 2")
}

func TestFindIntersection(t *testing.T) {
	type tc struct {
		w1       string
		w2       string
		expected int
	}

	tcs := []tc{
		{
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			6,
		},
		{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			159,
		},
		{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			135,
		},
	}

	for _, test := range tcs {
		p, _ := wires.FindIntersection(test.w1, test.w2)
		assert.Equal(t, test.expected, p.Distance())
	}
}
