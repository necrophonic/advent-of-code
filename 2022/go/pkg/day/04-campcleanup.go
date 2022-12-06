package day

import (
	"strconv"
	"strings"

	"github.com/necrophonic/advent-of-code/pkg/advent"
	"github.com/pkg/errors"
)

type CampCleanup struct{}

func (c CampCleanup) Name() string {
	return "Camp Cleanup"
}

func (c CampCleanup) Answer() (int, int, error) {
	data, err := advent.LoadData("campcleanup.txt")
	if err != nil {
		return -1, -1, err
	}

	part1, err := c.Part1(data)
	if err != nil {
		return -1, -1, errors.WithMessage(err, "part2")
	}

	part2, err := c.Part2(data)
	if err != nil {
		return -1, -1, errors.WithMessage(err, "part2")
	}

	return part1, part2, nil
}

func (cc CampCleanup) Part1(data []string) (int, error) {
	count := 0
	for ai, d := range data {
		assignments := strings.Split(d, ",")
		cr1, err := cc.NewCampArea(assignments[0])
		if err != nil {
			return -1, errors.WithMessagef(err, "failed assignment '%d'", ai)
		}
		cr2, err := cc.NewCampArea(assignments[1])
		if err != nil {
			return -1, errors.WithMessagef(err, "failed assignment '%d'", ai)
		}

		if cr1.Contains(cr2) || cr2.Contains(cr1) {
			count++
		}
	}
	return count, nil
}

func (cc CampCleanup) Part2(data []string) (int, error) {
	count := 0
	for ai, d := range data {
		assignments := strings.Split(d, ",")
		cr1, err := cc.NewCampArea(assignments[0])
		if err != nil {
			return -1, errors.WithMessagef(err, "failed assignment '%d'", ai)
		}
		cr2, err := cc.NewCampArea(assignments[1])
		if err != nil {
			return -1, errors.WithMessagef(err, "failed assignment '%d'", ai)
		}

		if cr1.Overlaps(cr2) || cr2.Overlaps(cr1) {
			count++
		}
	}
	return count, nil
}

type CampArea struct {
	Lower int
	Upper int
}

func (CampCleanup) NewCampArea(rnge string) (*CampArea, error) {
	ends := strings.Split(rnge, "-")

	lowerInt, err := strconv.Atoi(ends[0])
	if err != nil {
		return nil, err
	}
	upperInt, err := strconv.Atoi(ends[1])
	if err != nil {
		return nil, err
	}

	return &CampArea{
		Upper: upperInt,
		Lower: lowerInt,
	}, nil
}

func (ca *CampArea) Contains(check *CampArea) bool {
	return ca.Lower <= check.Lower && ca.Upper >= check.Upper
}

func (ca *CampArea) Overlaps(check *CampArea) bool {
	return (ca.Lower <= check.Lower && ca.Upper >= check.Lower) || (ca.Lower <= check.Upper && ca.Upper >= check.Upper)
}
