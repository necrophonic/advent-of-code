package fuel

import (
	"bufio"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// DataFile defines where to read input data from
var DataFile = "data/fuel.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	f, err := os.Open(DataFile)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	fuelMass := 0
	totalFuelMass := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		componentMass := scanner.Text()
		mass, err := strconv.Atoi(componentMass)
		if err != nil {
			return 0, 0, errors.WithMessage(err, "failed to parse mass value")
		}
		componentFuelMass := ForMass(mass)
		fuelMass += componentFuelMass
		totalFuelMass += ForFuelMass(componentFuelMass)
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return fuelMass, totalFuelMass, nil
}

// ForMass calculates the amount of
// fuel needed for a given mass.
func ForMass(mass int) int {
	return mass/3 - 2
}

// ForFuelMass calculates the amount of
// fuel needed for a given mass of fuel.
func ForFuelMass(mass int) int {
	add := ForMass(mass)
	if add < 1 {
		return mass
	}
	total := add + mass
	for {
		add = ForMass(add)
		if add < 1 {
			break
		}
		total += add
	}
	return total
}
