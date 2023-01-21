package test

import (
	"fmt"
	"testing"

	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
)

var totalCostTests = []struct {
	path     path.Path
	expected float64
}{
	{
		path.NewPath(&[]int{0, 1, 4}),
		5.5,
	},
	{
		path.NewPath(&[]int{0, 4}),
		maxF,
	},
	{
		path.NewPath(&[]int{3, 1, 0}),
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
	path     path.Path
	expected []path.Path
}{
	{
		path.NewPath(&[]int{0, 1, 4}),
		[]path.Path{
			path.NewPath(&[]int{0, 1, 4, 3}),
		},
	},
	{
		path.NewPath(&[]int{0, 1}),
		[]path.Path{
			path.NewPath(&[]int{0, 1, 3}),
			path.NewPath(&[]int{0, 1, 4}),
		},
	},
	{
		path.NewPath(&[]int{0}),
		[]path.Path{
			path.NewPath(&[]int{0, 1}),
			path.NewPath(&[]int{0, 2}),
			path.NewPath(&[]int{0, 5}),
		},
	},
	{
		path.NewPath(&[]int{1, 0}),
		[]path.Path{
			path.NewPath(&[]int{1, 0, 2}),
			path.NewPath(&[]int{1, 0, 5}),
		},
	},
}

func TestExpand(t *testing.T) {
	for i, tt := range expandTests {
		testName := fmt.Sprintf("Expand test number %v", i)
		t.Run(testName, func(t *testing.T) {
			res := tt.path.Expand(testGraph)
			if !path.ComparePathsSlices(*res, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, res)
			}
		})
	}
}
