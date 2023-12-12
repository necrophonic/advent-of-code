package gear_test

import (
	"testing"

	gear "github.com/necrohonic/advent-of-code/2023/go/03-gear"
	"github.com/necrohonic/advent-of-code/advent"
	"github.com/necrohonic/advent-of-code/advent/data"
)

func TestDay(t *testing.T) {
	// Same test data for both parts today
	d := data.Data{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	advent.RunTest(t,
		&advent.TestOpts{
			N:             "3",
			Day:           &gear.Gear{},
			InPart1:       d,
			ExpectedPart1: 4361,
			InPart2:       d,
			ExpectedPart2: 467835,
		},
	)
}
