package gear

import (
	"fmt"
	"strconv"

	"github.com/necrohonic/advent-of-code/advent"
	"github.com/necrohonic/advent-of-code/advent/data"
)

type Gear struct {
	data.Store
}

func (g *Gear) Init() error {
	data, err := data.Load("03-gear")
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}
	g.D = data

	return nil
}

func isPart(engine []string, x, y int) bool {
	return isSymbol(engine[y][x])
}

func isSymbol(r byte) bool {
	return (r < '0' || r > '9') && r != '.'
}

type Match struct {
	C Coord
	S byte
}

type PartNumber struct {
	Val        string
	Start, End Coord
	Valid      bool
	Matches    []Match
}

func (pn *PartNumber) String() string {
	return fmt.Sprintf("%-5s [%v %v] (Matches: %v)", pn.Val, pn.Start, pn.End, pn.Matches)
}

type Coord struct{ X, Y int }

func (c Coord) String() string { return fmt.Sprintf("%d-%d", c.X, c.Y) }

func (g Gear) Part1() (any, error) {
	numbers, _, err := parseEngine(g.D)
	if err != nil {
		return advent.NoAnswer, err
	}

	total := 0
	for _, n := range numbers {
		if !n.Valid {
			continue
		}
		add, err := strconv.Atoi(n.Val)
		if err != nil {
			return nil, err
		}
		total += add
	}
	return total, nil
}

func (g Gear) Part2() (any, error) {
	_, gears, err := parseEngine(g.D)
	if err != nil {
		return advent.NoAnswer, err
	}

	total := 0
	for _, gearLinks := range gears {
		if len(gearLinks) != 2 {
			continue
		}
		link1, err := strconv.Atoi(gearLinks[0])
		if err != nil {
			return advent.NoAnswer, err
		}
		link2, err := strconv.Atoi(gearLinks[1])
		if err != nil {
			return advent.NoAnswer, err
		}
		total += link1 * link2
	}
	return total, nil
}

func parseEngine(dta data.Data) (numbers []*PartNumber, gears map[string][]string, err error) {
	gears = map[string][]string{}

	// Now we have the basic data, parse it into the engine
	engine := []string{}
	for _, d := range dta {
		engine = append(engine, d)
	}

	var number *PartNumber
	for ri, row := range engine {
		lastci := 0
		for ci, r := range row {
			lastci = ci
			rv := advent.RuneToDigit(r)
			switch {
			case rv >= 0 && rv <= 9:
				if number == nil {
					number = &PartNumber{
						Start:   Coord{X: ci, Y: ri},
						Matches: []Match{},
					}
				}
				number.Val += string(r)

			case isSymbol(byte(r)), r == '.':
				if number != nil {
					coord := Coord{X: ci - 1}

					// X is -1 as we'll have read off the end.
					number.End = coord
					numbers = append(numbers, number)
				}
				number = nil
			}
		}
		if number != nil {
			// In this case we've reached the end of the row, so we need to
			// output if we've got a number in progress. Don't need to minus
			// one as the index will not have moved forward.
			number.End = Coord{X: lastci, Y: ri}
			numbers = append(numbers, number)
			number = nil
		}
	}

	// Go through each number and search for parts around then.
	for _, n := range numbers {
		xs, xe := leftRight(n, engine, n.Start.Y)

		// Above at y-1, below at y+1
		for _, ydelta := range []int{-1, 0, 1} {
			y := n.Start.Y + ydelta
			if y < 0 || y > len(engine)-1 {
				continue
			}
			for x := xs; x <= xe; x++ {
				if isPart(engine, x, y) {
					if engine[y][x] == '*' {
						gearLinks := gears[Coord{X: x, Y: y}.String()]
						gearLinks = append(gearLinks, n.Val)
						gears[Coord{X: x, Y: y}.String()] = gearLinks
					}
					n.Matches = append(n.Matches, Match{S: engine[y][x], C: Coord{X: x, Y: y}})
					n.Valid = true
				}
			}
		}
	}
	return numbers, gears, nil
}

func leftRight(n *PartNumber, engine []string, y int) (int, int) {
	xs := n.Start.X - 1
	if xs < 0 {
		xs = 0 // Left bound
	}
	xe := n.End.X + 1
	if xe > len(engine[y])-1 {
		xe = len(engine[y]) - 1 // Right bound
	}
	return xs, xe
}
