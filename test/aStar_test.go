package test

import (
	"fmt"
	"testing"

	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/explorationAlgorithm"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/heuristic"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/sortAlgorithm"
	floatutils "github.com/lucastomic/ExploracionDeEspacios/pkg/utils/floatUtils"
)

// Heuristic which returns the number of adjacents
type nAdjacentsHeur struct{}

func (h nAdjacentsHeur) Heuristic(state int) float64 {
	var adyacentsAmount float64
	for _, cost := range testGraph[state] {
		if floatutils.NotInfinit(cost) {
			adyacentsAmount++
		}
	}
	return adyacentsAmount
}

// Heuristic which dislikes 3
type dislike3Heur struct{}

// Rerturns 10 if the state == 3, returns 1 otherwise
func (h dislike3Heur) Heuristic(state int) float64 {
	if state == 3 {
		return 10
	} else {
		return 1
	}
}

var aStarTests = []struct {
	heuristic heuristic.StateHeuristic
	olds      *[]path.Path
	news      *[]path.Path
	expected  *[]path.Path
}{
	{
		nAdjacentsHeur{},
		&[]path.Path{
			path.NewPath(&[]int{0, 2}),
		},
		&[]path.Path{
			path.NewPath(&[]int{0, 1, 4}),
			path.NewPath(&[]int{0, 1, 3}),
		},

		&[]path.Path{
			path.NewPath(&[]int{0, 2}),
			path.NewPath(&[]int{0, 1, 3}),
			path.NewPath(&[]int{0, 1, 4}),
		},
	},
	{
		dislike3Heur{},
		&[]path.Path{
			path.NewPath(&[]int{0, 2}),
		},
		&[]path.Path{
			path.NewPath(&[]int{0, 1, 4}),
			path.NewPath(&[]int{0, 1, 3}),
		},

		&[]path.Path{
			path.NewPath(&[]int{0, 2}),
			path.NewPath(&[]int{0, 1, 4}),
			path.NewPath(&[]int{0, 1, 3}),
		},
	},
}

func TestAStar(t *testing.T) {
	for i, tt := range aStarTests {
		testName := fmt.Sprintf("A* test number %v", i)
		t.Run(testName, func(t *testing.T) {
			explorationAlgorithm.NewAStar(sortAlgorithm.NewMergeSort(), testGraph, tt.heuristic).Merge(tt.olds, tt.news)
			if !path.ComparePathsSlices(*tt.olds, *tt.expected) {
				t.Errorf("Expected: %v, got: %v", *tt.expected, *tt.olds)
			}
		})
	}
}
