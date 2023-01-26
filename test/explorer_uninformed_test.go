package test

import (
	"fmt"
	"testing"

	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/explorationAlgorithm/enums/uninformedAlgorithm"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/explorer"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/solution"
	solutiontests "github.com/lucastomic/ExploracionDeEspacios/test/solution_tests"
)

var exploreWithUninformedTests = []struct {
	solution    solution.Solution
	initalState int
	expected    path.Path
}{
	{
		solutiontests.NewThreeSolution(),
		5,
		path.NewPath(&[]int{5, 0, 1, 3}),
	},
	{
		solutiontests.NewThreeSolution(),
		0,
		path.NewPath(&[]int{0, 1, 3}),
	},
	{
		solutiontests.NewThreeSolution(),
		2,
		path.NewPath(&[]int{2, 0, 1, 3}),
	},
	{
		solutiontests.NewThreeSolution(),
		1,
		path.NewPath(&[]int{1, 3}),
	},
}

func TestExploreBranchAndBound(t *testing.T) {
	for i, tt := range exploreWithUninformedTests {
		testName := fmt.Sprintf("Explore branch and bound test number %v", i)
		t.Run(testName, func(t *testing.T) {
			pathSolution, err := explorer.ExploreWithUninformed(testGraph, tt.solution, tt.initalState, uninformedAlgorithm.BranchAndBonud)
			if err != nil {
				t.Error(err)
			}
			if !pathSolution.Equal(tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected.ToString(), pathSolution.ToString())
			}
		})
	}
}
func TestExploreDepthFirst(t *testing.T) {
	for i, tt := range exploreWithUninformedTests {
		testName := fmt.Sprintf("Explore depth first test number %v", i)
		t.Run(testName, func(t *testing.T) {
			pathSolution, err := explorer.ExploreWithUninformed(testGraph, tt.solution, tt.initalState, uninformedAlgorithm.DepthFirst)
			if err != nil {
				t.Error(err)
			}
			if !tt.solution.IsSolution(pathSolution.GetCurrentState(), testGraph) {
				t.Errorf("Path solution is not a valid solution, %v", pathSolution.GetCurrentState())
			}
		})
	}
}

func TestExploreAmplitude(t *testing.T) {
	for i, tt := range exploreWithUninformedTests {
		testName := fmt.Sprintf("Explore depth first test number %v", i)
		t.Run(testName, func(t *testing.T) {
			pathSolution, err := explorer.ExploreWithUninformed(testGraph, tt.solution, tt.initalState, uninformedAlgorithm.Amplitude)
			if err != nil {
				t.Error(err)
			}
			if !tt.solution.IsSolution(pathSolution.GetCurrentState(), testGraph) {
				t.Errorf("Path solution is not a valid solution, %v", pathSolution.GetCurrentState())
			}
		})
	}
}
