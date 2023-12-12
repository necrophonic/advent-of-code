package scratch

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/necrohonic/advent-of-code/advent"
	"github.com/necrohonic/advent-of-code/advent/data"
)

type Scratch struct {
	data.Store
}

func (s *Scratch) Init() error {
	data, err := data.Load("04-scratch")
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}
	s.D = data
	return nil
}

func (s Scratch) Part1() (any, error) {
	total := 0.0
	for cnmo, card := range s.D {
		cn := cnmo - 1
		winning, myNumbers, err := parseTicket(card)
		if err != nil {
			return advent.NoAnswer, fmt.Errorf("game %d: %w", cn, err)
		}
		wins := countWins(winning, myNumbers)
		if wins > 0 {
			total += math.Pow(2, float64(wins)-1)
		}
	}
	return int(total), nil
}

func (s Scratch) Part2() (any, error) {
	cardMap := map[int]int{}

	for cnmo, card := range s.D {
		cn := cnmo + 1
		cardMap[cn]++

		winning, myNumbers, err := parseTicket(card)
		if err != nil {
			return advent.NoAnswer, fmt.Errorf("game %d: %w", cn, err)
		}

		wins := countWins(winning, myNumbers)
		for w := 1; w <= wins; w++ {
			// Increment that next card by the number of this
			// card we have (as it'll "win" that many times)
			cardMap[cn+w] += cardMap[cn]
		}
	}

	total := 0
	for _, count := range cardMap {
		total += count
	}
	return total, nil
}

func countWins(winning, myNumbers []string) int {
	count := 0
	for _, number := range myNumbers {
		if number == "" {
			continue // Deal with double spaces in formatted input creating null numbers
		}
		if slices.Contains(winning, number) {
			count++
		}
	}
	return count
}

func parseTicket(ticket string) (winning, myNumbers []string, err error) {
	game := strings.Split(ticket, ":")
	if len(game) != 2 {
		return nil, nil, fmt.Errorf("invalid game: %s", ticket)
	}

	numbers := strings.Split(game[1], "|")
	if len(game) != 2 {
		return nil, nil, fmt.Errorf("invalid game: %s", ticket)
	}

	winning = strings.Split(strings.Trim(numbers[0], " "), " ")
	myNumbers = strings.Split(strings.Trim(numbers[1], " "), " ")
	return winning, myNumbers, nil
}
