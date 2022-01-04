package day

import (
	"sort"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
	"github.com/pkg/errors"
)

type Lanternfish struct{}

func (*Lanternfish) Name() string {
	return "Lanternfish"
}

func (f *Lanternfish) Answer() (int, int, error) {
	data, err := advent.LoadIntSliceData("lanternfish.txt")
	if err != nil {
		return -1, -1, err
	}
	ans1, err := f.Solver(80, data)
	if err != nil {
		return -1, -1, errors.WithMessage(err, "part 1 failed")
	}
	ans2, err := f.Solver(256, data)
	if err != nil {
		return -1, -1, errors.WithMessage(err, "part 2 failed")
	}

	return ans1, ans2, nil
}

func (f *Lanternfish) Solver(days int, fish []int) (int, error) {
	sort.Ints(fish)

	// Groups 0 .. 8
	groups := make([]int, 9)

	for _, f := range fish {
		groups[f]++
	}

	for d := 0; d < days; d++ {
		spawn := 0
		// Anything 0, becomes 6 (so 7 before we minus)
		groups[7] += groups[0]
		// We also want to spawn xGroup0 at the end of the day
		spawn = groups[0]

		// Now decay
		groups = append(groups[1:], 0)

		// Add new spawns
		groups[8] = spawn
	}

	count := 0
	for _, f := range groups {
		count += f
	}
	return count, nil
}
