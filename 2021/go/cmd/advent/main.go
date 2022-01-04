package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
)

type Day interface {
	Answer() (int, int, error)
	Name() string
}

func main() {
	run(1, &day.Sonar{})
	run(2, &day.Dive{})
	run(3, &day.BinaryDiagnostic{})
	run(4, &day.GiantSquid{})
	run(5, &day.Hyper{})
	run(6, &day.Lanternfish{})
	run(7, &day.Whales{})
	run(8, &day.SevenSegment{})
}

func run(number int8, day Day) {
	part1, part2, err := day.Answer()
	if err != nil {
		fmt.Printf(color.RedString("- [%2d] %30s errored: %v\n"), number, day.Name(), err)
		return
	}

	ans1 := color.YellowString("<no answer>")
	if part1 > -1 {
		ans1 = color.GreenString(strconv.Itoa(part1))
	}
	ans2 := color.YellowString("<no answer>")
	if part2 > -1 {
		ans2 = color.GreenString(strconv.Itoa(part2))
	}

	fmt.Printf("- [%2d] %30s part 1: %-20s part 2: %-20s\n", number, color.BlueString(day.Name()), ans1, ans2)
}
