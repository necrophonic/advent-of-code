package luggage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// DataFile defines where to read input data from
var DataFile = "data/luggage.txt"

// BagMap is a mapping of bag names to definitions
type BagMap map[string]*Bag

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}
	rules := bytes.Split(data, []byte("\n"))
	bags := ParseRules(rules)

	possible := make(BagMap)

	possibleCount := FindPossibleBags("shiny gold", possible, bags)

	totalBagsContained := NumberBagsContained("shiny gold", bags)

	return possibleCount, totalBagsContained - 1, nil
}

// Bag is a luggage bag
type Bag struct {
	Name     string
	Contains map[string]int
}

// CanContain returns true if the given bag
// can be contained in this bag.
func (b Bag) CanContain(name string) bool {
	_, ok := b.Contains[name]
	return ok
}

// NumberBagsContained returns the total number of bags contained
// in the given bag type (including the bag itself)
func NumberBagsContained(name string, bags BagMap) int {
	count := 1
	for bagName := range bags[name].Contains {
		count += NumberBagsContained(bagName, bags) * bags[name].Contains[bagName]
	}
	return count
}

// FindPossibleBags returns all the bags that could possibly
// enventually contain the given bag.
func FindPossibleBags(findName string, possible BagMap, bags BagMap) int {
	for bagName, bag := range bags {
		if bag.CanContain(findName) {
			fmt.Printf("%s can contain %s\n", bagName, findName)
			if _, exists := possible[bagName]; !exists {
				FindPossibleBags(bagName, possible, bags)
			}
			possible[bagName] = bag
		}
	}
	return len(possible)
}

var reRule = regexp.MustCompile(`(\d+) (\w+ \w+)(?:, |)`)

// ParseRules creates a set of bags from the given rules
func ParseRules(rules [][]byte) BagMap {
	bags := make(BagMap)
	for _, rule := range rules {
		if len(rule) == 0 {
			continue
		}

		s := strings.Split(string(rule), " bags contain")

		bag := &Bag{
			Name:     s[0],
			Contains: make(map[string]int),
		}

		matches := reRule.FindAllStringSubmatch(s[1], -1)

		for _, match := range matches {
			amount, err := strconv.Atoi(match[1])
			if err != nil {
				// Don't like to panic, but in this case if this
				// fails then things are corrupt and we can't continue.
				panic(err)
			}
			bag.Contains[match[2]] = amount
		}
		bags[bag.Name] = bag
	}
	return bags
}
