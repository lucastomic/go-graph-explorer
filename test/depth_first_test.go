package test

import (
	"fmt"
	"testing"

	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/explorationAlgorithm"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
)

var depthFirstTests = []struct {
	olds     *[]path.Path
	news     *[]path.Path
	expected *[]path.Path
}{
	{
		&[]path.Path{
			path.NewPath(&[]int{5, 0, 2}),
		},
		&[]path.Path{
			path.NewPath(&[]int{5, 0, 1}),
		},
		&[]path.Path{
			path.NewPath(&[]int{5, 0, 2}),
			path.NewPath(&[]int{5, 0, 1}),
		},
	},
	{
		&[]path.Path{
			path.NewPath(&[]int{5}),
		},
		&[]path.Path{
			path.NewPath(&[]int{5, 1}),
			path.NewPath(&[]int{5, 2}),
		},
		&[]path.Path{
			path.NewPath(&[]int{5}),
			path.NewPath(&[]int{5, 1}),
			path.NewPath(&[]int{5, 2}),
		},
	},
	{
		&[]path.Path{
			path.NewPath(&[]int{1}),
			path.NewPath(&[]int{2}),
			path.NewPath(&[]int{3}),
		},
		&[]path.Path{
			path.NewPath(&[]int{4}),
			path.NewPath(&[]int{5}),
		},
		&[]path.Path{
			path.NewPath(&[]int{1}),
			path.NewPath(&[]int{2}),
			path.NewPath(&[]int{3}),
			path.NewPath(&[]int{4}),
			path.NewPath(&[]int{5}),
		},
	},
}

func TestDepthFirst(t *testing.T) {
	for i, tt := range depthFirstTests {
		testName := fmt.Sprintf("Merge test number %v", i)
		t.Run(testName, func(t *testing.T) {
			explorationAlgorithm.NewDepthFirst().Merge(tt.olds, tt.news)
			if !path.ComparePathsSlices(*tt.olds, *tt.expected) {
				DisplaySlicesExpectedGot(tt.expected, tt.olds, t)
			}
		})
	}
}
