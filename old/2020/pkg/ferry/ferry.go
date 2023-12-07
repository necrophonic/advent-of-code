package ferry

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// DataFile defines where to read input data from
var DataFile = "data/ferry.txt"

// Answer provides the day's answers
func Answer() (p1, p2 int, err error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}

	instructions := string(data)

	ferry := &Ferry{
		Facing: E,
	}
	if err := ferry.PlotCourse(instructions, Move); err != nil {
		return 0, 0, err
	}

	// waypointFerry := &Ferry{
	// 	Facing:   E,
	// 	Waypoint: Waypoint{10, -1},
	// }
	// if err := waypointFerry.PlotCourse(instructions, MoveWaypoint); err != nil {
	// 	return 0, 0, err
	// }

	return ferry.Distance(), -1, nil
}

// Direction constants
const (
	N = iota
	E
	S
	W
	F
)

// Turn constants - used for rotational calculations
const (
	L = 180
	R = 360 // 360 instead of zero so as not to clash with N!
)

var (
	directions   = []string{"N", "E", "S", "W"}
	directionMap = map[string]int{"N": N, "E": E, "S": S, "W": W, "F": F, "L": L, "R": R}
)

// Waypoint is a waypoint
type Waypoint []int

func (w Waypoint) X() int {
	return w[0]
}

func (w Waypoint) Y() int {
	return w[1]
}

func (w Waypoint) rotate(direction, angle int) {
	angle = (angle + direction) % 360
	for steps := angle / 90; steps > 0; steps-- {
		z := w[0]
		w[0] = w[1] * -1
		w[1] = z
	}
}

// Ferry represents the current state of the ferry
type Ferry struct {
	Facing   int
	X, Y     int
	Waypoint Waypoint
}

func (f *Ferry) String() string {
	r := map[int]string{N: "N", E: "E", S: "S", W: "W", F: "F", L: "L", R: "R"}
	return fmt.Sprintf("X: %4d, Y: %4d [%s]", f.X, f.Y, r[f.Facing])
}

// Distance returns the Manhattan distance between the current
// ferry location and it's starting origin.
func (f *Ferry) Distance() int {
	return int(math.Abs(float64(f.X)) + math.Abs(float64(f.Y)))
}

// PlotCourse plots a course
func (f *Ferry) PlotCourse(directions string, move func(*Ferry, int, int)) error {
	for _, directionInstruction := range strings.Split(directions, "\n") {
		fmt.Println(directionInstruction)
		if directionInstruction == "" {
			continue
		}
		ds := strings.SplitN(directionInstruction, "", 2)
		direction := ds[0]
		distance, err := strconv.Atoi(ds[1])
		if err != nil {
			return err
		}
		move(f, directionMap[direction], distance)
	}
	return nil
}

// MoveWaypoint will move the ferry and it's waypoints
func MoveWaypoint(f *Ferry, direction int, distance int) {
	switch direction {
	case N:
		f.Waypoint = Waypoint{f.Waypoint.X(), f.Waypoint.Y() - distance}
	case S:
		f.Waypoint = Waypoint{f.Waypoint.X(), f.Waypoint.Y() + distance}
	case E:
		f.Waypoint = Waypoint{f.Waypoint.X() + distance, f.Waypoint.Y()}
	case W:
		f.Waypoint = Waypoint{f.Waypoint.X() - distance, f.Waypoint.Y()}
	case F:
		for d := 0; d < distance; d++ {
			f.X += f.Waypoint.X()
			f.Y += f.Waypoint.Y()
		}
	default:
		f.Waypoint.rotate(direction, distance)
	}
	fmt.Println(f)
	return
}

// Move will move the ferry
func Move(f *Ferry, direction int, distance int) {
	switch direction {
	case N:
		f.Y -= distance
	case S:
		f.Y += distance
	case E:
		f.X += distance
	case W:
		f.X -= distance
	case F:
		Move(f, f.Facing, distance)
	default:
		f.ChangeHeading(direction, distance)
	}
	return
}

// ChangeHeading will rotate the heading of the ferry
func (f *Ferry) ChangeHeading(direction, distance int) {
	steps := distance / 90
	if direction == L {
		f.Facing = (f.Facing + 4 - steps) % 4
		return
	}
	f.Facing = (f.Facing + steps) % 4
}
