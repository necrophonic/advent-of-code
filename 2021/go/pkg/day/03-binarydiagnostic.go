package day

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
)

type BinaryDiagnostic struct{}

func (b BinaryDiagnostic) Name() string {
	return "Binary Diagnostic"
}

func (b BinaryDiagnostic) Answer() (int, int, error) {
	f, err := advent.OpenFile("binary.txt")
	if err != nil {
		return -1, -1, nil
	}
	defer f.Close()

	part1Answer, common := b.part1(f)

	return part1Answer, b.part2(common), nil
}

func (BinaryDiagnostic) part1(file *os.File) (int, []string) {

	var counts []int
	totalLines := 0
	lineLength := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		totalLines++

		if counts == nil {
			lineLength = len(text)
			counts = make([]int, lineLength)
		}

		bits := strings.Split(text, "")
		for i, digit := range bits {
			if digit == "1" {
				counts[i]++
			}
		}
	}

	finalOnes := make([]string, lineLength)
	finalZeros := make([]string, lineLength)
	for i, onesCount := range counts {
		if onesCount >= totalLines/2 {
			finalOnes[i] = "1"
			finalZeros[i] = "0"
			continue
		}
		finalOnes[i] = "0"
		finalZeros[i] = "1"
	}

	ones, err := strconv.ParseInt(strings.Join(finalOnes, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	zeros, err := strconv.ParseInt(strings.Join(finalZeros, ""), 2, 64)
	if err != nil {
		panic(err)
	}

	return int(ones * zeros), finalOnes
}

func (BinaryDiagnostic) part2(common []string) int {

	return -1
}
