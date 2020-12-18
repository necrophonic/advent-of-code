package seating

import (
	"bytes"
	"io/ioutil"
)

// DataFile defines where to read input data from
var DataFile = "data/game.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}
	data = bytes.TrimRight(data, "\n")

	changes := 0

	g1 := NewGrid(data)
	for {
		if changes = g1.Round(); changes == 0 {
			break
		}
	}

	changes = 0
	g2 := NewGrid(data)
	for {
		if changes = g2.PickyRound(); changes == 0 {
			break
		}
	}

	return g1.NumberOccupied, g2.NumberOccupied, nil
}

// Grid represents the seating grid
type Grid struct {
	g              []byte // Real grid
	s              []byte // Shadow grid for writing
	RowLength      int
	Height         int
	NumberOccupied int
}

const (
	occupied byte = '#'
	empty         = 'L'
	floor         = '.'
	space         = ' '
)

// NewGrid creates a new grid from the given input
func NewGrid(data []byte) *Grid {
	rows := bytes.Split(data, []byte("\n"))
	g := append(bytes.Join(rows, []byte(" ")), []byte(" ")...)

	grid := &Grid{
		g:         g, // Grid
		s:         make([]byte, len(g)),
		RowLength: len(rows[0]),
		Height:    len(rows),
	}
	copy(grid.s, grid.g)
	return grid
}

// GetString returns the current grid string
func (g *Grid) GetString() string {
	return string(g.g)
}

// GetXY fetches a value from the virtual grid
func (g *Grid) GetXY(x, y int) byte {
	if g.invalidLocation(x, y) {
		return space
	}
	return g.g[g.xyToLinear(x, y)]
}

// GetShadowXY fetches a value from the shadow grid
func (g *Grid) GetShadowXY(x, y int) byte {
	if g.invalidLocation(x, y) {
		return space
	}
	return g.s[g.xyToLinear(x, y)]
}

// SetXY sets an element at the given location on the shadow grid
func (g *Grid) SetXY(x, y int, element byte) {
	if g.invalidLocation(x, y) {
		return
	}
	g.s[g.xyToLinear(x, y)] = element
}

func (g *Grid) xyToLinear(x, y int) int {
	return x + (y * (g.RowLength + 1))
}

func (g *Grid) invalidLocation(x, y int) bool {
	if x < 0 || y < 0 || x >= g.RowLength || y >= g.Height {
		return true
	}
	return false
}

// Round performs a round of seating-of-life
func (g *Grid) Round() (changes int) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.RowLength; x++ {
			element := g.GetXY(x, y)
			switch element {
			case empty:
				if g.Neighbours(x, y) == 0 {
					g.SetXY(x, y, occupied)
					changes++
					g.NumberOccupied++
				}
			case occupied:
				if g.Neighbours(x, y) >= 4 {
					g.SetXY(x, y, empty)
					changes++
					g.NumberOccupied--
				}
			}
		}
	}
	copy(g.g, g.s)
	return changes
}

// Neighbours returns the number of occupied neighbours to the given location
func (g *Grid) Neighbours(x, y int) (numOccupied int) {
	for sy := y - 1; sy <= y+1; sy++ {
		for sx := x - 1; sx <= x+1; sx++ {
			if x == sx && y == sy {
				continue
			}
			if g.GetXY(sx, sy) == occupied {
				numOccupied++
			}
		}
	}
	return numOccupied
}

type matrix struct{ x, y int }

var checks = []matrix{
	{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1},
}

// PickyRound performs a round of seating-of-life
func (g *Grid) PickyRound() (changes int) {
	i := 0
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.RowLength; x++ {
			element := g.GetXY(x, y)
			switch element {
			case empty:
				if g.PickyNeighbours(x, y) == 0 {
					g.SetXY(x, y, occupied)
					changes++
					g.NumberOccupied++
				}
			case occupied:
				if g.PickyNeighbours(x, y) >= 5 {
					g.SetXY(x, y, empty)
					changes++
					g.NumberOccupied--
				}
			}
			i++
		}
	}
	copy(g.g, g.s)
	return changes
}

// PickyNeighbours returns the number of "neighbours" in eight directions
func (g *Grid) PickyNeighbours(x, y int) (numOccupied int) {
	for _, check := range checks {
		if numOccupied += g.checkLocation(x, y, check); numOccupied == 5 {
			return numOccupied
		}
	}
	return numOccupied
}

func (g *Grid) checkLocation(x, y int, modifier matrix) int {
	nx := x + modifier.x
	ny := y + modifier.y
	for !g.invalidLocation(nx, ny) {
		switch g.GetXY(nx, ny) {
		case empty:
			return 0
		case occupied:
			return 1
		}
		nx += modifier.x
		ny += modifier.y
	}
	return 0
}
