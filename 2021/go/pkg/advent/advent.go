package advent

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Answer(day int, p1, p2 interface{}) {
	fmt.Printf("Answers to Day %d:\n\tPart 1: %v\n\tPart 2: %v\n", day, p1, p2)
}

var DataFolder = "data/"

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
		if item != "" {
			i, err := strconv.Atoi(item)
			if err != nil {
				return nil, errors.WithMessage(err, "failed to parse data item to int")
			}
			data = append(data, i)
		}
	}
	return data, nil
}
