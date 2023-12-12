package trebuchet_test

import (
	"testing"

	trebuchet "github.com/necrohonic/advent-of-code/2023/go/01-trebuchet"
	"github.com/necrohonic/advent-of-code/advent"
	"github.com/necrohonic/advent-of-code/advent/data"
)

func TestDay(t *testing.T) {
	advent.RunTest(t,
		&advent.TestOpts{
			N:   "2",
			Day: &trebuchet.Trebuchet{},
			InPart1: data.Data{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			ExpectedPart1: 142,
			InPart2: data.Data{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			ExpectedPart2: 281,
		},
	)
}
