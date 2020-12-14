package ferry_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/ferry"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	ferry.DataFile = "../../data/ferry.txt"
	p1, p2, err := ferry.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 521, p1, "Part 1")
	assert.Equal(t, -1, p2, "Part 2")
}

func TestMove(t *testing.T) {
	f := &ferry.Ferry{
		Facing: ferry.E,
		X:      10,
		Y:      10,
	}

	ferry.Move(f, ferry.N, 10)
	assert.Equal(t, &ferry.Ferry{Facing: ferry.E, X: 10, Y: 0}, f)

	ferry.Move(f, ferry.R, 90)
	assert.Equal(t, &ferry.Ferry{Facing: ferry.S, X: 10, Y: 0}, f)

	ferry.Move(f, ferry.F, 10)
	assert.Equal(t, &ferry.Ferry{Facing: ferry.S, X: 10, Y: 10}, f)
}

func TestMoveWaypoint(t *testing.T) {
	f := &ferry.Ferry{
		Facing:   ferry.E,
		Waypoint: ferry.Waypoint{10, -1},
	}

	ferry.MoveWaypoint(f, ferry.F, 10)
	assert.Equal(t, &ferry.Ferry{
		Facing:   ferry.E,
		X:        100,
		Y:        -10,
		Waypoint: ferry.Waypoint{10, -1},
	}, f)

	ferry.MoveWaypoint(f, ferry.N, 3)
	assert.Equal(t, &ferry.Ferry{
		Facing:   ferry.E,
		X:        100,
		Y:        -10,
		Waypoint: ferry.Waypoint{10, -4},
	}, f)

	ferry.MoveWaypoint(f, ferry.F, 7)
	assert.Equal(t, &ferry.Ferry{
		Facing:   ferry.E,
		X:        170,
		Y:        -38,
		Waypoint: ferry.Waypoint{10, -4},
	}, f)

	ferry.MoveWaypoint(f, ferry.R, 90)
	assert.Equal(t, &ferry.Ferry{
		Facing:   ferry.E,
		X:        170,
		Y:        -38,
		Waypoint: ferry.Waypoint{4, 10},
	}, f)

	ferry.MoveWaypoint(f, ferry.F, 11)
	assert.Equal(t, &ferry.Ferry{
		Facing:   ferry.E,
		X:        214,
		Y:        72,
		Waypoint: ferry.Waypoint{4, 10},
	}, f)

	assert.Equal(t, 286, f.Distance())
}

func TestChangeHeading(t *testing.T) {
	f := &ferry.Ferry{Facing: ferry.E}

	steps := []struct {
		direction int
		distance  int
		expected  int
	}{
		{ferry.R, 90, ferry.S},
		{ferry.R, 90, ferry.W},
		{ferry.R, 90, ferry.N},
		{ferry.L, 180, ferry.S},
		{ferry.R, 90, ferry.W},
		{ferry.R, 270, ferry.S},
	}

	for _, test := range steps {
		f.ChangeHeading(test.direction, test.distance)
		assert.Equal(t, &ferry.Ferry{Facing: test.expected}, f)
	}
}
