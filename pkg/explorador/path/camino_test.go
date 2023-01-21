package path

import (
	"fmt"
	"math"
	"reflect"
	"testing"
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

var totalCostTests = []struct {
	path     Path
	expected float64
}{
	{
		NewPath(&[]int{0, 1, 4}),
		5.5,
	},
	{
		NewPath(&[]int{0, 4}),
		maxF,
	},
	{
		NewPath(&[]int{3, 1, 0}),
		3,
	},
}

func TestTotalCost(t *testing.T) {
	for i, tt := range totalCostTests {
		testName := fmt.Sprintf("Total cost test number %v", i)
		t.Run(testName, func(t *testing.T) {
			res := tt.path.GetTotalCost(testGraph)
			if res != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, res)
			}
		})
	}
}

var expandTests = []struct {
	path     Path
	expected []Path
}{
	{
		NewPath(&[]int{0, 1, 4}),
		[]Path{
			NewPath(&[]int{0, 1, 4, 3}),
		},
	},
	{
		NewPath(&[]int{0, 1}),
		[]Path{
			NewPath(&[]int{0, 1, 3}),
			NewPath(&[]int{0, 1, 4}),
		},
	},
	{
		NewPath(&[]int{0}),
		[]Path{
			NewPath(&[]int{0, 1}),
			NewPath(&[]int{0, 2}),
			NewPath(&[]int{0, 5}),
		},
	},
	{
		NewPath(&[]int{1, 0}),
		[]Path{
			NewPath(&[]int{1, 0, 2}),
			NewPath(&[]int{1, 0, 5}),
		},
	},
}

// ComparePaths check whether two paths have the same state
func ComparePaths(pathA, pathB []Path) bool {
	if len(pathA) != len(pathB) {
		return false
	}
	for i := range pathA {
		aStates := pathA[i].states
		bStates := pathB[i].states
		if !reflect.DeepEqual(*aStates, *bStates) {
			return false
		}
	}
	return true
}

func TestExpand(t *testing.T) {
	for i, tt := range expandTests {
		testName := fmt.Sprintf("Expand test number %v", i)
		t.Run(testName, func(t *testing.T) {
			res := tt.path.Expand(testGraph)
			if !ComparePaths(*res, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, res)
			}
		})
	}
}
