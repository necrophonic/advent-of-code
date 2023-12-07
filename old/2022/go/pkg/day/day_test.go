package day_test

import (
	"os"
	"testing"

	"github.com/necrophonic/advent-of-code/pkg/advent"
	"github.com/necrophonic/advent-of-code/pkg/debug"
)

func TestMain(m *testing.M) {
	advent.DataFolder = "../../data"
	debug.On = true
	os.Exit(m.Run())
}
