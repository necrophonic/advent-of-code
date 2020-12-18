package bus_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/bus"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	bus.DataFile = "../../data/bus.txt"
	p1, p2, err := bus.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 2095, p1, "Part 1")
	assert.Equal(t, -1, p2, "Part 2")
}

func TestEarliest(t *testing.T) {
	earliestDepart := 939
	schedule, _ := bus.ParseSchedule("7,13,x,x,59,x,31,19")
	assert.Equal(t, 295, bus.FindEarliest(earliestDepart, schedule))
}
