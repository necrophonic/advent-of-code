package bus

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
)

// DataFile defines where to read input data from
var DataFile = "data/bus.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}
	data = bytes.TrimRight(data, "\n")

	dataParts := strings.Split(string(data), "\n")

	earliestDepart, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return 0, 0, err
	}

	schedule, err := ParseSchedule(dataParts[1])
	if err != nil {
		return 0, 0, err
	}

	answerPart1 := FindEarliest(earliestDepart, schedule)
	answerPart2 := FindSpecialEarliest(schedule)

	return answerPart1, answerPart2, nil
}

// FindSpecialEarliest finds the earliest for part 2
func FindSpecialEarliest(schedule []int) int {
	// for t := schedule[0]; ; t += t {
	// 	if busSequence(t, schedule) {
	// 		return t
	// 	}
	// }
	return 0
}

func busSequence(start int, schedule []int) bool {
	for b := 1; b < len(schedule); b++ {
		if s := schedule[b]; s != 0 && s != schedule[b-1] {
			return false
		}
	}
	return true
}

// FindEarliest finds the earliest possible bus
func FindEarliest(depart int, schedule []int) int {
	for t := depart; ; t++ {
		for _, bus := range schedule {
			if bus > 0 && t%bus == 0 {
				return bus * (t - depart)
			}
		}
	}
}

// ParseSchedule parses a string of scheules into a
// numeric slice of buses
func ParseSchedule(raw string) ([]int, error) {
	buses := []int{}
	for _, bus := range strings.Split(raw, ",") {
		if bus == "x" {
			buses = append(buses, 0)
			continue
		}
		busNumber, err := strconv.Atoi(bus)
		if err != nil {
			return nil, err
		}
		buses = append(buses, busNumber)
	}
	return buses, nil
}
