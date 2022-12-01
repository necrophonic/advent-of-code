package day

import (
	"sort"

	"github.com/necrophonic/advent-of-code/pkg/advent"
)

type Calorie struct{}

func (c Calorie) Name() string {
	return "Calorie counting"
}

func (c Calorie) Answer() (int, int, error) {
	data, err := advent.LoadIntDataSliceOfSlices("calorie.txt")
	if err != nil {
		return -1, -1, err
	}

	// Build a sorted list of Elf calorie totals
	calorieTotals := make([]int, 0, len(data))
	for _, cals := range data {
		calorieTotals = append(calorieTotals, sumIntSlice(cals))
	}
	sort.Ints(calorieTotals)

	return c.part1(calorieTotals), c.part2(calorieTotals), nil
}

func (Calorie) part1(totals []int) int {
	return totals[len(totals)-1]
}

func (Calorie) part2(totals []int) int {
	return sumIntSlice(totals[len(totals)-3:])
}

// Helper for quickly summing a slice of ints
func sumIntSlice(s []int) (sum int) {
	for _, i := range s {
		sum += i
	}
	return
}
