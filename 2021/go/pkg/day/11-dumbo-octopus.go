package day

import (
	"fmt"
	"strings"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
)

type DumboOctopus struct{}

func (*DumboOctopus) Name() string {
	return "Dumbo Octopus"
}

type Octopus struct {
	Flashed bool
	Energy  int
}

type OctopusGrid [][]*Octopus // y, x

func (og OctopusGrid) Print() {
	for _, row := range og {
		for _, o := range row {
			fmt.Printf("%d ", o.Energy)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (og OctopusGrid) String() string {
	stringified := ""
	for _, row := range og {
		for _, o := range row {
			stringified += fmt.Sprintf("%d", o.Energy)
		}
	}
	return stringified
}

func (og OctopusGrid) Get(x, y int) *Octopus {
	if y > len(og)-1 || y < 0 || x < 0 || x > len(og[0])-1 {
		return nil
	}
	return og[y][x]
}

func (og OctopusGrid) Increment(x, y int) {
	if octo := og.Get(x, y); octo != nil && !octo.Flashed {
		octo.Energy++
	}
}

func (d *DumboOctopus) Answer() (int, int, error) {
	data, err := advent.LoadData("dumbo-octopus.txt")
	if err != nil {
		return -1, -1, err
	}

	part1, _ := d.Flash(d.MakeGrid(data), 100)
	_, part2 := d.Flash(d.MakeGrid(data), 600)

	return part1, part2, nil
}

func (d *DumboOctopus) Flash(og OctopusGrid, maxSteps int) (int, int) {
	totalFlashes := 0
	firstAllFlash := 0

	for step := 1; step <= maxSteps; step++ {
		stepFlashes := 0
		// Increment all by 1
		for y, row := range og {
			for x := range row {
				octo := og.Get(x, y)
				octo.Energy++
				octo.Flashed = false
			}
		}

		// Loop through flashes until no more flashes occur
		for {
			flashes := 0
			for y, row := range og {
				for x := range row {
					if octo := og.Get(x, y); !octo.Flashed && octo.Energy > 9 {
						octo.Energy = 0
						octo.Flashed = true
						flashes++
						// Set surrounding octos +1
						og.Increment(x-1, y-1)
						og.Increment(x-1, y)
						og.Increment(x-1, y+1)

						og.Increment(x, y-1)
						og.Increment(x, y+1)

						og.Increment(x+1, y-1)
						og.Increment(x+1, y)
						og.Increment(x+1, y+1)
					}
				}
			}
			if step <= maxSteps {
				totalFlashes += flashes
			}
			stepFlashes += flashes
			if flashes == 0 {
				break
			}
		}
		if stepFlashes == 100 && firstAllFlash == 0 {
			firstAllFlash = step
		}
	}
	return totalFlashes, firstAllFlash
}

func (d *DumboOctopus) MakeGrid(data []string) OctopusGrid {
	og := make(OctopusGrid, 10)
	for y, row := range data {
		og[y] = make([]*Octopus, 0, 10)
		for _, octo := range strings.Split(row, "") {
			og[y] = append(og[y], &Octopus{Energy: int([]byte(octo)[0]) - 48})
		}
	}
	return og
}
