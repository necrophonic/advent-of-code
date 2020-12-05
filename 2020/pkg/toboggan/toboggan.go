package toboggan

import (
	"bytes"
	"io/ioutil"
)

// DataFile defines where to read input data from
var DataFile = "data/toboggan.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}

	rows := bytes.Split(data, []byte("\n"))
	return runPart1(rows), runPart2(rows), nil
}

func runPart1(rows [][]byte) (numTrees int) {
	tree := 0
	x := 0
	for _, row := range rows {
		tree, x = checkForTree(row, x, 3, len(row))
		numTrees += tree
	}
	return numTrees
}

func runPart2(rows [][]byte) (numTrees int) {

	type strategy struct {
		right int
		down  int
	}

	routes := []strategy{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	// Need to start with a base 1, or the
	// multiplication will end up zero!
	numTrees = 1

	for _, route := range routes {
		routeTrees := 0
		tree := 0
		x := 0
		for r := 0; r < len(rows); r += route.down {
			row := rows[r]
			tree, x = checkForTree(row, x, route.right, len(row))
			routeTrees += tree
		}
		if routeTrees > 0 {
			numTrees *= routeTrees
		}
	}
	return numTrees
}

func checkForTree(row []byte, curX, incX, width int) (int, int) {
	if width == 0 {
		return 0, curX + incX
	}
	if row[curX%width] == '#' {
		return 1, curX + incX
	}
	return 0, curX + incX
}
