package test

import (
	"fmt"
	"testing"

	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorationAlgorithm"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/path"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/sortAlgorithm"
)

var branchAndBoundTests = []struct {
	olds     *[]path.Path
	news     *[]path.Path
	expected *[]path.Path
}{
	{
		&[]path.Path{
			path.NewPath(&[]int{0, 2}),
			path.NewPath(&[]int{0, 1, 4, 3}),
		},
		&[]path.Path{
			path.NewPath(&[]int{0, 1, 4}),
			path.NewPath(&[]int{0, 1, 3}),
		},

		&[]path.Path{
			path.NewPath(&[]int{0, 1, 4}),
			path.NewPath(&[]int{0, 1, 3}),
			path.NewPath(&[]int{0, 2}),
		},
	},
	{
		&[]path.Path{
			path.NewPath(&[]int{0, 2}),
			path.NewPath(&[]int{0, 5}),
		},
		&[]path.Path{
			path.NewPath(&[]int{0, 1, 4}),
			path.NewPath(&[]int{0, 1, 3, 4}),
			path.NewPath(&[]int{0, 1, 3}),
		},

		&[]path.Path{
			path.NewPath(&[]int{0, 1, 3, 4}),
			path.NewPath(&[]int{0, 1, 3}),
			path.NewPath(&[]int{0, 2}),
			path.NewPath(&[]int{0, 5}),
		},
	},
}

func TestBranchAndBound(t *testing.T) {
	for i, tt := range branchAndBoundTests {
		testName := fmt.Sprintf("Branch&Bound test number %v", i)
		t.Run(testName, func(t *testing.T) {
			bAndB := explorationAlgorithm.NewBranchAndBound(sortAlgorithm.NewMergeSort(), testGraph)
			bAndB.Merge(tt.olds, tt.news)
			if !path.ComparePathsSlices(*tt.olds, *tt.expected) {
				DisplaySlicesExpectedGot(tt.expected, tt.olds, t)
			}
		})
	}
}
