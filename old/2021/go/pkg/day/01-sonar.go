package day

import (
	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
)

type Sonar struct{}

func (s Sonar) Name() string {
	return "Sonar Sweep"
}

func (s Sonar) Answer() (int, int, error) {
	data, err := advent.LoadIntData("sonar-sweep.txt")
	if err != nil {
		return -1, -1, err
	}
	return s.part1(data), s.part2(data), nil
}

func (Sonar) part1(data []int) int {
	count := 0
	current := data[0]
	for _, next := range data[1:] {
		if next > current {
			count++
		}
		current = next
	}
	return count
}

func (s Sonar) part2(data []int) int {
	count := 0
	l := len(data)
	for i := 0; i < l-3; i++ {
		window1 := s.sumWindow(data[i : i+3])
		window2 := s.sumWindow(data[i+1 : i+4])
		if window2 > window1 {
			count++
		}
	}
	return count
}

func (Sonar) sumWindow(window []int) (sum int) {
	for _, window := range window {
		sum += window
	}
	return
}
