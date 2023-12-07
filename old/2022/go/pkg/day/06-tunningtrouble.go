package day

import (
	"github.com/necrophonic/advent-of-code/pkg/advent"
)

type Tuning struct{}

func (t Tuning) Name() string {
	return "Tuning Trouble"
}

func (t Tuning) Answer() (any, any, error) {
	data, err := advent.LoadData("tuningtrouble.txt")
	if err != nil {
		return -1, -1, err
	}

	signal := data[0]

	return t.Part1(signal), t.Part2(signal), nil
}

func (t Tuning) Part1(signal string) int {
	return t.findUniqueChunk(signal, 4)
}

func (t Tuning) Part2(signal string) int {
	return t.findUniqueChunk(signal, 14)
}

func (t Tuning) findUniqueChunk(signal string, length int) int {
	for start := 0; start < len(signal)-length; start++ {
		end := start + length
		if t.IsUnique(signal[start:end]) {
			return end
		}
	}
	return -1
}

func (Tuning) IsUnique(chunk string) bool {
	dedupe := make(map[rune]bool)
	for _, r := range chunk {
		dedupe[r] = true
	}
	return len(dedupe) == len(chunk)
}
