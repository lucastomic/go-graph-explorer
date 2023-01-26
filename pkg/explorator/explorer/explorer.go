package explorer

import (
	"errors"

	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorationAlgorithm"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorationAlgorithm/enums/informedAlgorithm"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorationAlgorithm/enums/uninformedAlgorithm"

	"github.com/lucastomic/go-graph-explorer/pkg/explorator/heuristic"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/path"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/solution"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/sortAlgorithm"
	"github.com/lucastomic/go-graph-explorer/pkg/utils/sliceUtils"
)

// explorer is a structure in charge of exploring a graph.
// You must pass by parameter:
// The graph to be evaluated in the form of an adjacency matrix
// The scan algorithm is defined by an object that implements the scanAlgorithm interface, which
// defines how the pending paths to be explored will be ordered once the current state is extended.
// The solution state is measured with the isSolution() method of the Solution structure
type explorer struct {
	graph       [][]float64
	algorithm   explorationAlgorithm.ExplorationAlgorithm
	solution    solution.Solution
	initalState int
}

// Indicates whether the explorer should keep looking for an optimal path.
// This will be so as long as there are still pending paths to search for and
// no solution found yet
func (e explorer) keepSearching(pendingPaths []path.Path, currentPath path.Path) bool {
	return !sliceUtils.IsEmpty(pendingPaths) && !e.solution.IsSolution(currentPath.GetCurrentState(), e.graph)
}

// explore the graph obtaining the searched path.
// To do this, create a list of pending paths, starting with the first node according to the adjacency matrix.
// Then, in each iteration, replace the last path with what is returned by the [e.expand()] method and
// reorders the paths depending on the [shuffle()] method of the scan algorithm passed by argument.
// Iterate until the method [e.seguirBuscando()] returns false
// If all the iterations are finished and the current state is not the solution, return an error indicating that the problem has no solution.
// Otherwise, return the current state.
func (e explorer) explore() (path.Path, error) {
	var pendingPaths []path.Path = make([]path.Path, 0)
	pendingPaths = append(pendingPaths, path.NewPath(&[]int{e.initalState}))
	currentPath := pendingPaths[len(pendingPaths)-1]

	for e.keepSearching(pendingPaths, currentPath) {
		newPaths := currentPath.Expand(e.graph)
		sliceUtils.RemoveLast(&pendingPaths)

		e.algorithm.Merge(&pendingPaths, newPaths)
		currentPath = pendingPaths[len(pendingPaths)-1]
	}

	if !e.solution.IsSolution(currentPath.GetCurrentState(), e.graph) {
		return path.NewEmptyPath(), errors.New("there is no solution")
	} else {
		return currentPath, nil
	}
}

// Explore a graph with an uninformed algorithm until it finds a solution state.
// graph is the graph to be explored
// solution must be a struct that implements the method IsSolution(int, [][]float64) and will indicate when a state is a solution to the problem
// algorithmType is the scan algorithm to use. Your options are:
//
// uninformedAlgorithm.BranchAndBonud (Branch&Bound)
// uninformedAlgorithm.DepthFirst(Depth)
// uninformedAlgorithm.Amplitude (Amplitude)
func ExploreWithUninformed(
	graph [][]float64,
	solution solution.Solution,
	initialState int,
	algorithmType uninformedAlgorithm.UniformedExpAlgorithmType,
) (path.Path, error) {
	var algorithm explorationAlgorithm.ExplorationAlgorithm
	switch algorithmType {
	case uninformedAlgorithm.BranchAndBonud:
		algorithm = explorationAlgorithm.NewBranchAndBound(sortAlgorithm.NewMergeSort(), graph)
	case uninformedAlgorithm.DepthFirst:
		algorithm = explorationAlgorithm.NewDepthFirst()
	case uninformedAlgorithm.Amplitude:
		algorithm = explorationAlgorithm.NewAmplitude()
	default:
		return path.NewEmptyPath(), errors.New("this algotirhm is not uninformed. Use ExploreWithInformed()")

	}
	explorer := explorer{
		graph:       graph,
		solution:    solution,
		algorithm:   algorithm,
		initalState: initialState,
	}
	return explorer.explore()
}

// Explore a graph with an informed algorithm until it finds a solution state.
// graph is the graph to be explored
// solution must be a struct that implements the method IsSolution(int, [][]float64) and will indicate when a state is a solution to the problem
// heuristic is a struct that must implement the method HeuristicoEstado(int)float64, which will return a value
// quantitative of how good that state is in terms of the solution
// algorithmType is the scan algorithm to use. Your options are:
//
//	informedAlgorithm.AlAStar (A*)
//	informedAlgorithm.AlClimbing (Climbing)
//	informedAlgorithm.AlBestFirst (Best First)
func ExploreWithInformed(
	graph [][]float64,
	solution solution.Solution,
	heuristic heuristic.StateHeuristic,
	initalState int,
	algorithmType informedAlgorithm.InformedExpAlgorithmType,
) (path.Path, error) {
	var algorithm explorationAlgorithm.ExplorationAlgorithm
	switch algorithmType {
	case informedAlgorithm.AStar:
		algorithm = explorationAlgorithm.NewAStar(sortAlgorithm.NewMergeSort(), graph, heuristic)
	case informedAlgorithm.Climbing:
		algorithm = explorationAlgorithm.NewClimbing(sortAlgorithm.NewMergeSort(), graph, heuristic)
	case informedAlgorithm.BestFirst:
		algorithm = explorationAlgorithm.NewBestFirst(sortAlgorithm.NewMergeSort(), graph, heuristic)
	default:
		return path.NewEmptyPath(), errors.New("this algotirhm is not informed. Use ExploreWithUninformed()")
	}
	explorer := explorer{
		graph:       graph,
		solution:    solution,
		algorithm:   algorithm,
		initalState: initalState,
	}
	return explorer.explore()
}
