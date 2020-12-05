package expenses

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
)

// DataFile defines where to read input data from
var DataFile = "data/expenses.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}

	expenses := dataToInt(bytes.Split(data, []byte("\n")))

	a1, b1, err := findPair(expenses, 2020, len(expenses))
	if err != nil {
		return 0, 0, err
	}
	part1 := a1 * b1

	a2, b2, c2, err := findTriplet(expenses, 2020, len(expenses))
	if err != nil {
		log.Fatal(err)
	}
	part2 := a2 * b2 * c2

	return part1, part2, nil
}

func dataToInt(in [][]byte) []int {
	out := make([]int, len(in))
	for i, s := range in {
		// Assuming the data is always good - don't do
		// this in production code!
		out[i], _ = strconv.Atoi(string(s))
	}
	return out
}

func findPair(exp []int, target, length int) (int, int, error) {
	for i, a := range exp[:len(exp)-1] {
		for j := i; j < len(exp); j++ {
			if a+exp[j] == target {
				return a, exp[j], nil
			}
		}
	}
	return 0, 0, errors.New("no matching pair")
}

func findTriplet(exp []int, target, length int) (int, int, int, error) {
	for i, a := range exp[:len(exp)-1] {
		remaining := target - a
		b, c, err := findPair(exp[i:], remaining, len(exp[i:]))
		if err != nil {
			continue
		}
		return a, b, c, nil
	}
	return 0, 0, 0, nil
}
