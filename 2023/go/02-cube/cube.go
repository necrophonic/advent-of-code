package cube

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/necrohonic/advent-of-code/advent"
	"github.com/necrohonic/advent-of-code/advent/data"
)

type Cube struct {
	data.Store
}

func (c *Cube) Init() error {
	data, err := data.Load("02-cube")
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}
	c.D = data
	return nil
}

var (
	reGame = regexp.MustCompile(`^Game (\d+)`)
	rePull = regexp.MustCompile(`^ (\d+) (red|green|blue)`)
)

// Part1 discovers which games could be played with the given bag of cubes.
func (c Cube) Part1() (any, error) {
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	total := 0

	for _, round := range c.D {

		game, err := parseGame(round)
		if err != nil {
			return advent.NoAnswer, fmt.Errorf("error parsing game: %v", err)
		}

		add := game.Number
		for _, pull := range game.Pulls {
			for colour, amount := range pull {
				if amount > bag[colour] {
					add = 0
				}
			}
		}
		total += add
	}
	return total, nil
}

// Part2 discovers the minimum number of each cube required to play each game.
func (c Cube) Part2() (any, error) {
	total := 0
	for _, round := range c.D {
		game, err := parseGame(round)
		if err != nil {
			return advent.NoAnswer, fmt.Errorf("error parsing game: %v", err)
		}
		mins := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, pull := range game.Pulls {
			for _, colour := range []string{"red", "green", "blue"} {
				if a := pull[colour]; a > mins[colour] {
					mins[colour] = a
				}
			}
		}
		total += mins["red"] * mins["green"] * mins["blue"]
	}
	return total, nil
}

type Game struct {
	Number int
	Pulls  []map[string]int

	parts []string
}

func parseGame(raw string) (*Game, error) {
	g := &Game{Number: -1, Pulls: []map[string]int{}}

	// Game <number>: <pulls>
	g.parts = strings.Split(raw, ":")
	if len(g.parts) != 2 {
		return nil, fmt.Errorf("error parsing game (invalid structure): %s", raw)
	}

	// Game number
	gameNumberStr := reGame.FindStringSubmatch(g.parts[0])[1]
	var err error
	g.Number, err = strconv.Atoi(gameNumberStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing game number: %v", err)
	}

	for _, pull := range strings.Split(g.parts[1], ";") {
		thisPull := map[string]int{}
		for _, colourPull := range strings.Split(pull, ",") {
			amountAndColour := rePull.FindStringSubmatch(colourPull)
			if len(amountAndColour) != 3 {
				return nil, fmt.Errorf("error parsing pull for colour %s: %s", colourPull, raw)
			}
			amount, err := strconv.Atoi(amountAndColour[1])
			if err != nil {
				return nil, fmt.Errorf("error parsing pull for colour %s: %v", colourPull, err)
			}
			thisPull[amountAndColour[2]] = amount
		}
		g.Pulls = append(g.Pulls, thisPull)
	}
	return g, nil
}
