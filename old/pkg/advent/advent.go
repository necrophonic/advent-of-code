// Package advent contains
package advent

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

var DataFolder = "data"

type EmptyDay struct{}

func (EmptyDay) Name() string {
	return "TODO"
}

func (EmptyDay) Answer() (any, any, error) {
	return nil, nil, nil
}

// Day defines the interface for a day of an advent challenge
type Day interface {
	Answer() (any, any, error)
	Name() string
}

func Run(number int8, day Day) {
	part1, part2, err := day.Answer()
	if err != nil {
		fmt.Printf(color.RedString("- [%2d] %35s errored: %v\n"), number, day.Name(), err)
		return
	}

	ans1 := color.YellowString("<no answer>")
	if part1 != nil {
		ans1 = color.GreenString(fmt.Sprintf("%v", part1))
	}
	ans2 := color.YellowString("<no answer>")
	if part2 != nil {
		ans2 = color.GreenString(fmt.Sprintf("%v", part2))
	}

	name := color.BlueString(day.Name())
	if strings.Contains(name, "TODO") {
		name = color.YellowString(day.Name())
	}

	fmt.Printf("- [%2d] %35s part 1: %-20s part 2: %-20s\n", number, name, ans1, ans2)
}

func OpenFile(fileName string) (*os.File, error) {
	f, err := os.Open(filepath.Join(DataFolder, fileName))
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to open file: %s", fileName)
	}
	return f, nil
}

// LoadData attempts to complete load a data file and returns the data in a slice
// containing one element per newline delimited line.
func LoadData(file string) ([]string, error) {
	f, err := OpenFile(file)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to load data")
	}
	defer f.Close()

	data := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			data = append(data, text)
		}
	}
	return data, nil
}

// LoadIntSliceData loads a set of ints, comma delimited, from a single line
func LoadIntSliceData(fileName string) ([]int, error) {
	data, err := ioutil.ReadFile(filepath.Join(DataFolder, fileName))
	if err != nil {
		return nil, err
	}
	sdata := strings.TrimSuffix(string(data), "\n")

	alphas := strings.Split(sdata, ",")
	ints := make([]int, len(alphas))
	for i, alpha := range alphas {
		number, err := strconv.Atoi(alpha)
		if err != nil {
			return nil, err
		}
		ints[i] = number
	}
	return ints, nil
}

// LoadIntDataSliceOfSlices loads blank line delimited chunks of newline delimited
// ints into grouped slices.
func LoadIntDataSliceOfSlices(file string) ([][]int, error) {
	f, err := os.Open(filepath.Join(DataFolder, file))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := [][]int{}
	slice := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		item := scanner.Text()
		if item == "" {
			data = append(data, slice)
			slice = make([]int, 0)
			continue
		}
		i, err := strconv.Atoi(item)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to parse data item to int")
		}
		slice = append(slice, i)

	}
	return data, nil
}

func LoadIntData(file string) ([]int, error) {
	f, err := os.Open(filepath.Join(DataFolder, file))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := []int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		item := scanner.Text()
		if item == "" {
			continue
		}
		i, err := strconv.Atoi(item)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to parse data item to int")
		}
		data = append(data, i)

	}
	return data, nil
}

// ByteToInt converts a single digit as a byte into a real int
// e.g '1' -> 1
func ByteToInt(b byte) int {
	return int(b) - 48
}

func StringToInt(b string) int {
	return int(b[0]) - 48
}
