package day

import "github.com/necrophonic/advent-of-code/2021/go/pkg/advent"

type SmokeBasin struct {
	Data []string
}

func (*SmokeBasin) Name() string {
	return "Smoke Basin"
}

func (sb *SmokeBasin) Answer() (int, int, error) {
	data, err := advent.LoadData("smoke-basin.txt")
	if err != nil {
		return -1, -1, nil
	}
	sb.Data = data

	part1, err := sb.Part1()
	if err != nil {
		return -1, -1, err
	}

	return part1, -1, nil
}

func (sb *SmokeBasin) Part1() (int, error) {
	lowSum := 0

	for r := 0; r < len(sb.Data); r++ {
		for c := 0; c < len(sb.Data[r]); c++ {
			lowSum += sb.LowPoint(r, c)
		}
	}

	return lowSum, nil
}

// LowPoint returns the value of the point +1 if it is a low point in the height
// map. Otherwise will return 0.
func (sb *SmokeBasin) LowPoint(row, col int) int {
	point := sb.GetPoint(row, col)

	// Left
	if point >= sb.GetPoint(row-1, col) {
		return 0
	}
	// Right
	if point >= sb.GetPoint(row+1, col) {
		return 0
	}
	// Top
	if point >= sb.GetPoint(row, col-1) {
		return 0
	}
	// Bottom
	if point >= sb.GetPoint(row, col+1) {
		return 0
	}
	return int(point) + 1
}

func (sb *SmokeBasin) GetPoint(row, col int) int {
	if row < 0 || row >= len(sb.Data) || col < 0 || col >= len(sb.Data[0]) {
		return 9
	}
	// We know we're only dealing with digits so can easy-convert from
	// ascii byte value to "real" value by -48 from the ascii code.
	return int(sb.Data[row][col]) - 48
}
