package test

import (
	"fmt"
	"testing"

	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/explorationAlgorithm"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
)

var amplitudeTests = []struct {
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

func TestAmplitude(t *testing.T) {
	for i, tt := range amplitudeTests {
		testName := fmt.Sprintf("Expand test number %v", i)
		t.Run(testName, func(t *testing.T) {
			explorationAlgorithm.NewAmplitude().Merge(tt.olds, tt.news)
			if !path.ComparePathsSlices(*tt.olds, *tt.expected) {
				t.Errorf("Expected: %v, got: %v", *tt.expected, *tt.olds)
			}
		})
	}
}
