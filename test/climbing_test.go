package test

import (
	"fmt"
	"testing"

	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorationAlgorithm"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/heuristic"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/path"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/sortAlgorithm"
	testheuristics "github.com/lucastomic/go-graph-explorer/test/test_heuristics"
)

var climbingtests = []struct {
	heuristic heuristic.StateHeuristic
	olds      *[]path.Path
	news      *[]path.Path
	expected  *[]path.Path
}{
	{
		testheuristics.NewNAdjacentHeur(testGraph),
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
	{
		testheuristics.NewDislike3Heur(),
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
}

func TestClimbing(t *testing.T) {
	for i, tt := range climbingtests {
		testName := fmt.Sprintf("Climbing test number %v", i)
		t.Run(testName, func(t *testing.T) {
			explorationAlgorithm.NewClimbing(sortAlgorithm.NewMergeSort(), testGraph, tt.heuristic).Merge(tt.olds, tt.news)
			if !path.ComparePathsSlices(*tt.olds, *tt.expected) {
				DisplaySlicesExpectedGot(tt.expected, tt.olds, t)
			}
		})
	}
}
