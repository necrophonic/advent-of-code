package day

import (
	"strconv"
	"strings"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
	"github.com/pkg/errors"
)

var bingoBoardLineLength = 5

type BingoBoard []int
type BingoLine []int

func (bb BingoBoard) Row(index int) BingoLine {
	offset := index * 5
	return BingoLine(bb[offset : offset+5])
}

func (bb BingoBoard) Col(index int) BingoLine {
	col := make(BingoLine, 5)
	for i := 0; i < 5; i++ {
		col[i] = bb[index+(i*5)]
	}
	return col
}

// Mark will set a number in the board (if found) to zero
func (bb BingoBoard) Mark(number int) {
	for i := 0; i < len(bb); i++ {
		if bb[i] == number {
			bb[i] = 0
		}
	}
}

// IsWinner checks whether a given line is a winning line
func (line BingoLine) IsWinner() bool {
	for _, number := range line {
		if number != 0 {
			return false
		}
	}
	return true
}

// IsWinner checks whether the board has any winning lines
func (bb BingoBoard) IsWinner() bool {
	for i := 0; i < 5; i++ {
		if bb.Row(i).IsWinner() || bb.Col(i).IsWinner() {
			return true
		}
	}
	return false
}

func (bb BingoBoard) Total() int {
	total := 0
	for _, number := range bb {
		total += number
	}
	return total
}

func NewBingoBoard(rows []string) (BingoBoard, error) {
	b := make(BingoBoard, bingoBoardLineLength*bingoBoardLineLength)
	for r, row := range rows {
		for n, number := range strings.Split(strings.ReplaceAll(strings.TrimSpace(row), "  ", " "), " ") {
			value, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			b[(r*5)+n] = value
		}
	}
	return b, nil
}

type GiantSquid struct{}

func (GiantSquid) Name() string {
	return "Giant Squid"
}

func (gs *GiantSquid) Parse(data []string) ([]int, []BingoBoard, error) {
	// First line is the bingo numbers
	numbers := []int{}
	for _, number := range strings.Split(data[0], ",") {
		value, err := strconv.Atoi(number)
		if err != nil {
			return nil, nil, errors.WithMessagef(err, "failed to parse number: '%s'", number)
		}
		numbers = append(numbers, value)
	}

	// Start reading through boards (skip the numbers line
	// and the first blank line)
	boards := []BingoBoard{}
	for boardData := data[1:]; len(boardData) >= 5; boardData = boardData[5:] {
		board, err := NewBingoBoard(boardData[:5])
		if err != nil {
			return nil, nil, errors.WithMessagef(err, "failed to parse board data")
		}
		boards = append(boards, board)
	}

	return numbers, boards, nil
}

func (g *GiantSquid) Answer() (int, int, error) {
	data, err := advent.LoadData("giantsquid.txt")
	if err != nil {
		return -1, -1, nil
	}

	numbers, boards, err := g.Parse(data)
	if err != nil {
		return -1, -1, err
	}

	part1Boards := make([]BingoBoard, len(boards))
	copy(part1Boards, boards)
	ans1, err := g.Part1(numbers, part1Boards)
	if err != nil {
		return -1, -1, err
	}

	part2Boards := make([]BingoBoard, len(boards))
	copy(part2Boards, boards)
	ans2, err := g.Part2(numbers, part1Boards)
	if err != nil {
		return -1, -1, err
	}

	return ans1, ans2, nil
}

func (g *GiantSquid) Part1(numbers []int, boards []BingoBoard) (int, error) {
	for _, number := range numbers {
		for b := 0; b < len(boards); b++ {
			boards[b].Mark(number)
			if boards[b].IsWinner() {
				return boards[b].Total() * number, nil
			}
		}
	}
	return 0, nil
}

func (g *GiantSquid) Part2(numbers []int, boards []BingoBoard) (int, error) {
	var lastWinner BingoBoard
	var lastNumber int
	for _, number := range numbers {

		for b := 0; b < len(boards); b++ {
			if boards[b] == nil {
				continue
			}
			boards[b].Mark(number)
			if boards[b].IsWinner() {
				lastWinner = boards[b]
				lastNumber = number
				// Remove the board from the search
				boards[b] = nil
			}
		}
	}
	return lastWinner.Total() * lastNumber, nil
}
