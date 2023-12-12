package trebuchet

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/necrohonic/advent-of-code/advent/data"
)

type Trebuchet struct {
	data.Store
}

func (t *Trebuchet) Init() error {
	data, err := data.Load("01-trebuchet")
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}
	t.D = data
	return nil
}

func (t Trebuchet) Part1() (any, error) {
	reFirst := regexp.MustCompile(`^[^0-9]*(\d)`)
	reLast := regexp.MustCompile(`(\d)[^0-9]*$`)

	total := 0
	for _, calibration := range t.D {
		f := reFirst.FindStringSubmatch(calibration)[1]
		l := reLast.FindStringSubmatch(calibration)[1]

		v, err := strconv.Atoi(f + l)
		if err != nil {
			return nil, err
		}
		total += v
	}
	return total, nil
}

var digitMap = map[string]string{
	"1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9",
	"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
}

func (t Trebuchet) Part2() (any, error) {
	total := 0
	for _, calibration := range t.D {
		var f, l string
		fi := -1
		li := -1

		for digit := range digitMap {
			if i := strings.Index(calibration, digit); i > -1 && (i < fi || fi == -1) {
				fi = i
				f = digitMap[digit]
			}
			if i := strings.LastIndex(calibration, digit); i > -1 && i > li {
				li = i
				l = digitMap[digit]
			}
		}
		v, err := strconv.Atoi(f + l)
		if err != nil {
			return nil, err
		}
		total += v
	}

	return total, nil
}
