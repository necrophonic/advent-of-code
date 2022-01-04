package day

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
)

type Dive struct{}

func (d Dive) Name() string {
	return "Dive"
}

func (d Dive) Answer() (int, int, error) {
	f, err := advent.OpenFile("dive.txt")
	if err != nil {
		return -1, -1, err
	}
	defer f.Close()
	return d.part1(f), d.part2(f), nil
}

func (d Dive) part1(file *os.File) int {
	horizontal := 0
	depth := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		command, distance := parse(text)
		switch command {
		case "forward":
			horizontal += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}
	return horizontal * depth
}

func (d Dive) part2(file *os.File) int {
	file.Seek(0, 0)

	horizontal := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		command, distance := parse(text)
		switch command {
		case "forward":
			horizontal += distance
			depth += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	return horizontal * depth
}

func parse(text string) (string, int) {
	split := strings.Split(text, " ")
	position := split[0]
	distance, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return position, distance
}
