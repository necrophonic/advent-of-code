package wires

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// DataFile defines where to read input data from
var DataFile = "data/wires.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	f, err := os.Open(DataFile)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	wire1 := scanner.Text()

	scanner.Scan()
	wire2 := scanner.Text()

	intersections, err := FindIntersections(wire1, wire2)
	if err != nil {
		return 0, 0, errors.WithMessage(err, "failed to trace wires")
	}

	return FindNearestIntersection(intersections).Distance(), FindShortestIntersection(intersections), nil
}

// FindNearestIntersection finds the closest wire intersection
func FindNearestIntersection(intersections map[string][]*Point) *Point {
	closestPoint := &Point{99999, 99999, 0}
	for _, points := range intersections {
		if points[0].Distance() < closestPoint.Distance() {
			closestPoint = points[0]
		}
	}
	return closestPoint
}

// FindShortestIntersection finds the intersection with the fewest steps
func FindShortestIntersection(intersections map[string][]*Point) int {
	shortest := 99999
	for _, points := range intersections {
		steps := points[0].Step + points[1].Step
		if steps < shortest {
			shortest = steps
		}
	}
	return shortest
}

// FindIntersections finds all wire intersections
func FindIntersections(w1, w2 string) (map[string][]*Point, error) {

	// Trace the first wire
	wiring := make(map[string][]*Point)
	cx := 0
	cy := 0

	i := 0
	for _, step := range strings.Split(w1, ",") {
		direction, distance, err := parseDirection(step)
		if err != nil {
			return nil, err
		}

		mx, my := moveMods(direction)
		for m := 0; m < distance; m++ {
			cx += mx
			cy += my
			i++
			wiring[fmt.Sprintf("%d-%d", cx, cy)] = []*Point{{cx, cy, i}, nil}
		}
	}

	// Trace the second wire
	cx = 0
	cy = 0
	i = 0

	intersections := make(map[string][]*Point)

	for _, step := range strings.Split(w2, ",") {
		direction, distance, err := parseDirection(step)
		if err != nil {
			return nil, err
		}

		mx, my := moveMods(direction)
		for m := 0; m < distance; m++ {
			cx += mx
			cy += my
			i++
			key := fmt.Sprintf("%d-%d", cx, cy)
			if p, exists := wiring[key]; exists {
				intersections[key] = []*Point{p[0], {cx, cy, i}}
			}
		}
	}

	return intersections, nil
}

func moveMods(direction string) (x, y int) {
	mx := 0
	my := 0
	switch direction {
	case "U":
		my = -1
	case "D":
		my = 1
	case "L":
		mx = -1
	case "R":
		mx = 1
	}
	return mx, my
}

// Point is a point on the wiring grid. It contains
// which step in the wiring it was laid down on.
type Point struct {
	X    int
	Y    int
	Step int
}

// Distance returns the Manhattan distance
// between the point and origin.
func (p Point) Distance() int {
	x := math.Abs(float64(p.X))
	y := math.Abs(float64(p.Y))
	return int(x + y)
}

var reDirection = regexp.MustCompile(`([RLUD])(\d+)`)

func parseDirection(raw string) (string, int, error) {
	matches := reDirection.FindStringSubmatch(raw)
	direction := matches[1]
	distance, err := strconv.Atoi(matches[2])
	if err != nil {
		return "", 0, err
	}
	return direction, distance, nil
}
