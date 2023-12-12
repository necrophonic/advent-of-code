package main

import (
	trebuchet "github.com/necrohonic/advent-of-code/2023/go/01-trebuchet"
	cube "github.com/necrohonic/advent-of-code/2023/go/02-cube"
	gear "github.com/necrohonic/advent-of-code/2023/go/03-gear"
	scratch "github.com/necrohonic/advent-of-code/2023/go/04-scratch"
	"github.com/necrohonic/advent-of-code/advent"
)

func main() {
	advent.SolveFor(advent.Solutions{
		1: &trebuchet.Trebuchet{},
		2: &cube.Cube{},
		3: &gear.Gear{},
		4: &scratch.Scratch{},
	})
}
