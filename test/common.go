package test

import (
	"math"
	"testing"

	"github.com/lucastomic/go-graph-explorer/pkg/explorator/path"
)

// (5)───0.5──(0)────0.5──(1)───5────┐
//             │           │         │
//             │           │        (4)
//            0.5         2.5        │
//             │           │         │
//             │           │         2
//             │           │         │
//            (2)         (3)────────┘

var maxF float64 = math.MaxFloat64
var testGraph [][]float64 = [][]float64{
	{maxF, 0.5, 0.5, maxF, maxF, 0.5},
	{0.5, maxF, maxF, 2.5, 5, maxF},
	{0.5, maxF, maxF, maxF, maxF, maxF},
	{maxF, 2.5, maxF, maxF, 2, maxF},
	{maxF, 5, maxF, 2, maxF, maxF},
	{0.5, maxF, maxF, maxF, maxF, maxF},
}

// Displays two path slices to can see the difference between them.
// It's usefull to compare two path slices when they were supposed to be equals in unit testing.
func DisplaySlicesExpectedGot(expected *[]path.Path, got *[]path.Path, t *testing.T) {
	var stringExpected string
	for _, path := range *expected {
		stringExpected += path.ToString() + "\n"
	}
	var stringGot string
	for _, path := range *got {
		stringGot += path.ToString() + "\n"
	}
	t.Errorf("Expected: \n%v \ngot: \n%v", stringExpected, stringGot)
}
