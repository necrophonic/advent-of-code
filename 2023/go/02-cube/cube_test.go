package cube_test

import (
	"testing"

	cube "github.com/necrohonic/advent-of-code/2023/go/02-cube"
	"github.com/necrohonic/advent-of-code/advent"
	"github.com/necrohonic/advent-of-code/advent/data"
)

func TestDay(t *testing.T) {
	// Same test data for both parts today
	d := data.Data{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	advent.RunTest(t,
		&advent.TestOpts{
			N:             "2",
			Day:           &cube.Cube{},
			InPart1:       d,
			ExpectedPart1: 8,
			InPart2:       d,
			ExpectedPart2: 2286,
		},
	)
}
