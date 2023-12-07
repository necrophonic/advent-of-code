package day

import (
	"strconv"
	"strings"

	"github.com/necrophonic/advent-of-code/pkg/advent"
	"github.com/pkg/errors"
)

type NoSpace struct {
	Structure *Directory
}

func (d NoSpace) Name() string {
	return "NoSpace Trouble"
}

func (d NoSpace) Answer() (any, any, error) {
	data, err := advent.LoadData("nospace.txt")
	if err != nil {
		return -1, -1, err
	}

	err = d.BuildStructure(data)
	if err != nil {
		return -1, -1, errors.WithMessage(err, "failed to build structure")
	}

	return d.Part1(data), d.Part2(data), nil
}

func (d NoSpace) Part1(data []string) int {
	return -1
}

func (d NoSpace) Part2(data []string) int {
	return -1
}

type Directory struct {
	Name    string
	Content []any
}

type File struct {
	Name string
	Size int
}

var ErrNoCurrentDirectory = errors.New("no current directory selected")
var ErrBadFileSize = errors.New("bad file size")

func (d *NoSpace) BuildStructure(commands []string) error {

	var currentDir *Directory

	for _, line := range commands {
		switch {
		case strings.HasPrefix(line, "$ cd "):
			// Change dir
			// name := strings.TrimPrefix(line, "$ cd ")

		case strings.HasPrefix(line, "$ ls"):
			// List
			// no-op

		case strings.HasPrefix(line, "dir"):
			// Declare a directory
			if currentDir == nil {
				return ErrNoCurrentDirectory
			}

			currentDir.Content = append(
				currentDir.Content,
				&Directory{
					Name:    strings.TrimPrefix(line, "dir "),
					Content: make([]any, 0),
				},
			)

		default:
			// A file to add
			if currentDir == nil {
				return ErrNoCurrentDirectory
			}
			parts := strings.Split(line, " ")
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				return errors.WithMessage(err, ErrBadFileSize.Error())
			}
			name := parts[1]

			currentDir.Content = append(
				currentDir.Content,
				&File{
					Name: name,
					Size: size,
				},
			)
		}
	}
	return nil
}
