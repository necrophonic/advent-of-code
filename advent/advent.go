package advent

import (
	"errors"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/necrohonic/advent-of-code/advent/data"
	"github.com/necrohonic/advent-of-code/advent/table"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const NoAnswer = "n/a"

type Solver interface {
	Init() error
	Part1() (any, error)
	Part2() (any, error)
	SetData(data.Data)
}

type Solutions map[int]Solver

func SolveFor(days Solutions) {
	daysSolved := make([]int, 0, len(days))
	for k := range days {
		daysSolved = append(daysSolved, k)
	}
	sort.Ints(daysSolved)

	tble := table.New()
	for _, day := range daysSolved {
		if err := days[day].Init(); err != nil {
			tble.Add(day, NoAnswer, NoAnswer, 0, err)
			continue
		}

		var e error
		start := time.Now()
		p1, err := days[day].Part1()
		if err != nil {
			e = errors.Join(e, err)
		}

		p2, err := days[day].Part2()
		if err != nil {
			e = errors.Join(e, err)
		}
		end := time.Since(start)

		tble.Add(day, p1, p2, end, e)
	}
	fmt.Println(tble)
}

type TestOpts struct {
	N                            string
	Day                          Solver
	InPart1, InPart2             data.Data
	ExpectedPart1, ExpectedPart2 any
}

// RunTest is a simple generic test runner for a day solver.
func RunTest(t *testing.T, tt *TestOpts) {
	t.Helper()

	if tt.InPart1 != nil {
		t.Run("Day "+tt.N+" Part 1", func(t *testing.T) {
			tt.Day.SetData(tt.InPart1)
			ans, err := tt.Day.Part1()
			require.NoError(t, err, "should not error")
			require.NotNil(t, ans, "should return a value")
			require.IsType(t, tt.ExpectedPart1, ans, "should return the expected type: %T", tt.ExpectedPart1)
			assert.Equal(t, tt.ExpectedPart1, ans, "should return the expected value")
		})
	}

	if tt.InPart2 != nil {
		t.Run("Day "+tt.N+" Part 2", func(t *testing.T) {
			tt.Day.SetData(tt.InPart2)
			ans, err := tt.Day.Part2()
			require.NoError(t, err, "should not error")
			require.NotNil(t, ans, "should return a value")
			require.IsType(t, tt.ExpectedPart2, ans, "should return the expected type: %T", tt.ExpectedPart2)
			assert.Equal(t, tt.ExpectedPart2, ans, "should return the expected value")
		})
	}
}

func RuneToDigit(r rune) int {
	r = r - 48
	if r >= 0 && r <= 9 {
		return int(r)
	}
	return -1
}
