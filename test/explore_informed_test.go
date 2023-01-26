package test

import (
	"fmt"
	"testing"

	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorationAlgorithm/enums/informedAlgorithm"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorer"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/heuristic"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/path"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/solution"
	solutiontests "github.com/lucastomic/go-graph-explorer/test/solution_tests"
	testheuristics "github.com/lucastomic/go-graph-explorer/test/test_heuristics"
)

var exploreWithInformedTests = []struct {
	solution    solution.Solution
	heuristic   heuristic.StateHeuristic
	initalState int
	expected    path.Path
}{
	{
		solutiontests.NewThreeSolution(),
		testheuristics.NewDislike3Heur(),
		5,
		path.NewPath(&[]int{5, 0, 1, 3}),
	},
	{
		solutiontests.NewThreeSolution(),
		testheuristics.NewDislike3Heur(),
		2,
		path.NewPath(&[]int{2, 0, 1, 3}),
	},
	{
		solutiontests.NewThreeSolution(),
		testheuristics.NewDislike1Heur(),
		2,
		path.NewPath(&[]int{2, 0, 1, 3}),
	},
	{
		solutiontests.NewThreeSolution(),
		testheuristics.NewDislike1Heur(),
		1,
		path.NewPath(&[]int{1, 3}),
	},
	{
		solutiontests.NewThreeSolution(),
		testheuristics.NewNAdjacentHeur(testGraph),
		5,
		path.NewPath(&[]int{5, 0, 1, 3}),
	},
	{
		solutiontests.NewThreeSolution(),
		testheuristics.NewNAdjacentHeur(testGraph),
		2,
		path.NewPath(&[]int{2, 0, 1, 3}),
	},
}

func TestExploreAStar(t *testing.T) {
	for i, tt := range exploreWithInformedTests {
		testName := fmt.Sprintf("Explore AStar test number %v", i)
		t.Run(testName, func(t *testing.T) {
			pathSolution, err := explorer.ExploreWithInformed(testGraph, tt.solution, tt.heuristic, tt.initalState, informedAlgorithm.AStar)
			if err != nil {
				t.Error(err)
			}
			if !pathSolution.Equal(tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected.ToString(), pathSolution.ToString())
			}
		})
	}
}

func TestExploreBestFirst(t *testing.T) {
	for i, tt := range exploreWithInformedTests {
		testName := fmt.Sprintf("Explore best first test number %v", i)
		t.Run(testName, func(t *testing.T) {
			pathSolution, err := explorer.ExploreWithInformed(testGraph, tt.solution, tt.heuristic, tt.initalState, informedAlgorithm.BestFirst)
			if err != nil {
				t.Error(err)
			}
			if !tt.solution.IsSolution(pathSolution.GetCurrentState(), testGraph) {
				t.Errorf("Path solution is not a valid solution, %v", pathSolution.GetCurrentState())
			}
		})
	}
}

func TestExploreClimbing(t *testing.T) {
	for i, tt := range exploreWithInformedTests {
		testName := fmt.Sprintf("Explore climbing test number %v", i)
		t.Run(testName, func(t *testing.T) {
			pathSolution, err := explorer.ExploreWithInformed(testGraph, tt.solution, tt.heuristic, tt.initalState, informedAlgorithm.Climbing)
			if err != nil {
				t.Error(err)
			}
			if !tt.solution.IsSolution(pathSolution.GetCurrentState(), testGraph) {
				t.Errorf("Path solution is not a valid solution, %v", pathSolution.GetCurrentState())
			}
		})
	}
}
