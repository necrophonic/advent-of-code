package encoding

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// DataFile defines where to read input data from
var DataFile = "data/encoding.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}

	numbers := make([]int, 0)
	for _, n := range strings.Split(string(data), "\n") {
		if n == "" {
			continue
		}
		d, err := strconv.Atoi(n)
		if err != nil {
			return 0, 0, err
		}
		numbers = append(numbers, d)
	}
	weakness := FindWeakness(numbers, 25)
	return weakness, FindContiguous(numbers, weakness), nil
}

// FindWeakness finds the weakness in the encoding.
func FindWeakness(numbers []int, preamble int) int {
	for i := preamble; i < len(numbers); i++ {
		pw := numbers[i-preamble : i]
		// Make a lookup map so we
		// only have to scan once.
		m := make(map[int]bool)
		for _, pn := range pw {
			m[pn] = true
		}

		ok := false
		next := numbers[i]
		for _, pn := range pw {
			if _, exists := m[next-pn]; exists && pn != next-pn {
				ok = true
				break
			}
		}
		if !ok {
			// Return this number if we
			// didn't manage to sum it.
			return next
		}
	}
	return 0
}

// FindContiguous finds the contiguous set of numbers that
// makes the target and returns the sum of the lowest and highest.
// Uses expanding/contracting slices over the input to find ranges.
func FindContiguous(numbers []int, target int) int {
	start := 0
	end := 0
	for start < len(numbers) {
		w := numbers[start:end]
		t := addSet(w)
		switch {
		case t == target:
			// Found it
			return sumSmallestAndLargest(w)
		case t > target:
			// Overshot, inc start and start expanding again
			start++
			end = start
		default:
			// Otherwise, expand the search window
			end++
		}
	}
	return 0
}

func addSet(numbers []int) (total int) {
	for _, n := range numbers {
		total += n
	}
	return
}

func sumSmallestAndLargest(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0] + numbers[len(numbers)-1]
}
