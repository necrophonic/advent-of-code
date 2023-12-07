package data

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var DataFolder = "../data"

func open(fileName string) (*os.File, error) {
	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}
	f, err := os.Open(filepath.Join(DataFolder, fileName))
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %w", fileName, err)
	}
	return f, nil
}

type Data []string

func (d Data) Copy() Data {
	d2 := make(Data, len(d))
	copy(d2, d)
	return d2
}

func (d Data) Len() int {
	return len(d)
}

func (d Data) Ints() ([]int, error) {
	ints := make([]int, 0, len(d))
	for _, s := range d {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to int", s)
		}
		ints = append(ints, i)
	}
	return ints, nil
}

// Load attempts to complete load a data file and returns the data in a slice
// containing one element per newline delimited line.
func Load(file string) (Data, error) {
	f, err := open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to load data: %w", err)
	}
	defer f.Close()

	data := Data{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			data = append(data, text)
		}
	}
	return data, nil
}
