package day

import (
	"strconv"
	"strings"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
	"github.com/pkg/errors"
)

type hyperGrid [][]int

func (hg hyperGrid) CountOverlaps() int {
	count := 0
	for _, row := range hg {
		for _, col := range row {
			if col >= 2 {
				count++
			}
		}
	}
	return count
}

type Hyper struct{}

func (*Hyper) Name() string {
	return "Hyperthermal Vents"
}

func (h *Hyper) ParseVents(data []string) ([]*Vent, error) {
	vents := []*Vent{}
	for _, ventString := range data {
		vent, err := NewVent(ventString)
		if err != nil {
			return nil, err
		}
		vents = append(vents, vent)
	}
	return vents, nil
}

type Point struct {
	X int
	Y int
}

func NewPointFromCoordsString(coords string) (*Point, error) {
	xy := strings.Split(coords, ",")
	p := &Point{}
	var err error
	if p.X, err = strconv.Atoi(xy[0]); err != nil {
		return nil, err
	}
	if p.Y, err = strconv.Atoi(xy[1]); err != nil {
		return nil, err
	}
	return p, nil
}

type Vent struct {
	Points [2]*Point
}

func NewVent(ventString string) (*Vent, error) {
	vent := &Vent{}
	coords := strings.Split(ventString, " -> ")
	var err error
	for i := 0; i < 2; i++ {
		if vent.Points[i], err = NewPointFromCoordsString(coords[i]); err != nil {
			return nil, errors.WithMessagef(err, "failed to parse vent coord: '%s'", coords[i])
		}
	}
	return vent, nil
}

func (v *Vent) DrawNoDiagonal(grid [][]int) {
	if v.Points[0].X == v.Points[1].X || v.Points[0].Y == v.Points[1].Y {
		v.Draw(grid)
	}
}

func (v *Vent) Draw(grid [][]int) {

	x1 := v.Points[0].X
	y1 := v.Points[0].Y

	x2 := v.Points[1].X
	y2 := v.Points[1].Y

	grid[y1][x1]++
	moveX, moveY := 1, 1

	if x2 < x1 {
		moveX = -1
	}
	if x2 == x1 {
		moveX = 0
	}
	if y2 < y1 {
		moveY = -1
	}
	if y2 == y1 {
		moveY = 0
	}

	for y2 != y1 || x2 != x1 {
		x1 += moveX
		y1 += moveY
		grid[y1][x1]++
	}
}

func (h *Hyper) Answer() (int, int, error) {
	data, err := advent.LoadData("hyperthermal-vents.txt")
	if err != nil {
		return -1, -1, nil
	}

	vents, err := h.ParseVents(data)
	if err != nil {
		return -1, -1, err
	}

	ans1, err := h.Part1(vents)
	if err != nil {
		return -1, -1, errors.WithMessage(err, "part 1 failed")
	}

	ans2, err := h.Part2(vents)
	if err != nil {
		return -1, -1, errors.WithMessage(err, "part 2 failed")
	}

	return ans1, ans2, nil
}

func (h *Hyper) Part1(vents []*Vent) (int, error) {
	// We know the grid can't be bigger than 1000x1000 from the input data.
	// Don't be this confident with production data ;)
	size := 1000
	grid := make(hyperGrid, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	for _, vent := range vents {
		vent.DrawNoDiagonal(grid)
	}
	return grid.CountOverlaps(), nil
}

func (h *Hyper) Part2(vents []*Vent) (int, error) {
	// We know the grid can't be bigger than 1000x1000 from the input data.
	// Don't be this confident with production data ;)
	size := 1000
	grid := make(hyperGrid, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	for _, vent := range vents {
		vent.Draw(grid)
	}
	return grid.CountOverlaps(), nil
}
