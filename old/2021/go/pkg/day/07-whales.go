package day

import (
	"sort"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
)

type Whales struct{}

func (*Whales) Name() string {
	return "Whales"
}

func (w *Whales) Answer() (int, int, error) {
	data, err := advent.LoadIntSliceData("whales.txt")
	if err != nil {
		return -1, -1, nil
	}

	ans1, err := w.Part1(data)
	if err != nil {
		return -1, -1, err
	}

	ans2, err := w.Part2(data)
	if err != nil {
		return -1, -1, err
	}

	return ans1, ans2, nil
}

func (w *Whales) Part1(data []int) (int, error) {
	lowestFuel := 0
	for _, distance := range data {
		if fuel := CalcFuelToDistance(distance, data); lowestFuel == 0 || fuel < lowestFuel {
			lowestFuel = fuel
		}
	}
	return lowestFuel, nil
}

func CalcFuelToDistance(target int, distances []int) int {
	fuel := 0
	for _, distance := range distances {
		thisFuel := distance - target
		if thisFuel < 0 {
			thisFuel *= -1
		}
		fuel += thisFuel
	}
	return fuel
}

func (w *Whales) Part2(data []int) (int, error) {
	sort.Ints(data)
	lowestFuel := 0
	for target := 0; target < data[len(data)-1]; target++ {
		fuel := CalcCrabFuelToDistance(target, data)
		if lowestFuel == 0 || fuel < lowestFuel {
			lowestFuel = fuel
		}
	}
	return lowestFuel, nil
}

func CalcCrabFuelToDistance(target int, distances []int) int {
	fuel := 0
	for _, distance := range distances {
		diff := distance - target
		if diff < 0 {
			diff *= -1
		}
		for i := 1; i <= diff; i++ {
			fuel += i
		}
	}
	return fuel
}
