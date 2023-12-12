package table_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/necrohonic/advent-of-code/advent/table"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	tble := table.New()
	tble.Add(1, 10000, 200, 0, nil)
	tble.Add(2, 10000, 200, 0, errors.New("something went wrong"))
	tble.Add(21, "cat", "elephant", 0, nil)
	tble.Add(5, -1, true, 0, nil)

	fmt.Println(tble.String())

	assert.Equal(t, "table", tble.String())

}
