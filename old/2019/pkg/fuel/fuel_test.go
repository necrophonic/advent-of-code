package fuel_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2019/pkg/fuel"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	fuel.DataFile = "../../data/fuel.txt"
	p1, p2, err := fuel.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 3434390, p1, "Part 1")
	assert.Equal(t, 5148724, p2, "Part 2")
}

func TestForMass(t *testing.T) {
	type tc struct {
		mass     int
		expected int
	}

	tcs := []tc{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, test := range tcs {
		assert.Equal(t, test.expected, fuel.ForMass(test.mass))
	}
}

func TestForFuelMass(t *testing.T) {
	type tc struct {
		mass     int
		expected int
	}

	tcs := []tc{
		{2, 2},
		{654, 966},
		{33583, 50346},
	}

	for _, test := range tcs {
		assert.Equal(t, test.expected, fuel.ForFuelMass(test.mass))
	}
}
