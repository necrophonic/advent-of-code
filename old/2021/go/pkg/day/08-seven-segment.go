package day

import (
	"strings"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
	"github.com/pkg/errors"
)

type SevenSegment struct{}

func (*SevenSegment) Name() string {
	return "Seven Segment Search"
}

func (s *SevenSegment) Answer() (int, int, error) {
	data, err := advent.LoadData("seven-segment.txt")
	if err != nil {
		return -1, -1, err
	}

	ans1, err := s.Part1(data)
	if err != nil {
		return -1, -1, errors.WithMessage(err, "part 1 failed")
	}
	// ans2, err := s.Part2(data)
	// if err != nil {
	// 	return -1, -1, errors.WithMessage(err, "part 2 failed")
	// }
	return ans1, -1, nil
}

func (s *SevenSegment) Part1(data []string) (int, error) {
	count := 0
	for _, dataLine := range data {
		_, output := s.SplitData(dataLine)

		for _, outputDigit := range output {
			length := len(outputDigit)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}
	return count, nil
}

func (s *SevenSegment) Part2(data []string) (int, error) {
	total := 0

	// Map between what the segments should be and what they are
	segments := make(map[string]string)

	for _, dataLine := range data {
		input, _ := s.SplitData(dataLine)

		// Find the basic digits we can identify
		digits := make([]string, 10)
		for _, digit := range input {
			switch len(digit) {
			case 2: // one
				digits[1] = digit
			case 3: // seven
				digits[7] = digit
			case 4: // fout
				digits[4] = digit
			case 7: // eight
				digits[7] = digit
			}
		}

		// Diff 1 with 7 to get top line
		segments["a"] = s.DiffDigits(digits[1], digits[7])

		// Diff the five segment digits to get e
		fives := make([]string, 0, 3)
		for _, digit := range input {
			if len(digit) == 5 {
				fives = append(fives, digit)
			}
		}
		segments["e"] = s.DiffDigits(digits[4], s.DiffDigits(fives...))

		// spew.Dump(segments)
	}

	return total, nil
}

func (*SevenSegment) DiffDigits(digits ...string) string {
	intersect := make(map[rune]int)
	for i := 0; i < len(digits); i++ {
		for _, segment := range digits[i] {
			intersect[segment]++
		}
	}
	diff := ""
	for r, count := range intersect {
		if count == 1 {
			diff += string(r)
		}
	}
	return diff
}

func (*SevenSegment) SplitData(input string) ([]string, []string) {
	dataParts := strings.Split(input, " | ")
	display := strings.Split(dataParts[0], " ")
	output := strings.Split(dataParts[1], " ")
	return display, output
}
