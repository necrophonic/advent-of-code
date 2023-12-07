package advent

import (
	"errors"
	"fmt"
	"sort"

	"github.com/necrohonic/advent-of-code/advent/table"
)

const NoAnswer = "n/a"

type Solver interface {
	Init() error
	Part1() (any, error)
	Part2() (any, error)
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
			tble.Add(day, NoAnswer, NoAnswer, err)
			continue
		}

		var e error
		p1, err := days[day].Part1()
		if err != nil {
			e = errors.Join(e, err)
		}

		p2, err := days[day].Part2()
		if err != nil {
			e = errors.Join(e, err)
		}

		tble.Add(day, p1, p2, e)
	}
	fmt.Println(tble)
}
