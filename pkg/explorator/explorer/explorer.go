package explorer

import (
	"errors"

	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/explorationAlgorithm"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/heuristic"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/solution"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/sortAlgorithm"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/utils/sliceUtils"
)

// explorer is a structure in charge of exploring a graph.
// You must pass by parameter:
// The graph to be evaluated in the form of an adjacency matrix
// The scan algorithm is defined by an object that implements the scanAlgorithm interface, which
// defines how the pending paths to be explored will be ordered once the current state is extended.
// The solution state is measured with the isSolution() method of the Solution structure
type explorer struct {
	graph      [][]float64
	algorithm  explorationAlgorithm.ExplorationAlgorithm
	solution   solution.Solution
	startState int
}

// Indicates whether the explorer should keep looking for an optimal path.
// This will be so as long as there are still pending paths to search for and
// no solution found yet
func (e explorer) keepSearching(pendingPaths []path.Path, currentPath path.Path) bool {
	return !sliceUtils.IsEmpty(pendingPaths) && !e.solution.IsSolution(currentPath.GetCurrentState(), e.graph)
}

// Explore the graph obtaining the searched path.
// To do this, create a list of pending paths, starting with the first node according to the adjacency matrix.
// Then, in each iteration, replace the last path with what is returned by the [e.expand()] method and
// reorders the paths depending on the [shuffle()] method of the scan algorithm passed by argument.
// Iterate until the method [e.seguirBuscando()] returns false
// If all the iterations are finished and the current state is not the solution, return an error indicating that the problem has no solution.
// Otherwise, return the current state.
func (e explorer) Explore() (path.Path, error) {
	var pendingPaths []path.Path
	pendingPaths[0] = path.NewPath(&[]int{e.startState})
	currentPath := pendingPaths[len(pendingPaths)-1]

	for e.keepSearching(pendingPaths, currentPath) {
		newPaths := currentPath.Expand(e.graph)
		sliceUtils.RemoveLast(newPaths)

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
// algoritmoexploracion.AlBranchAndBonud (Branch&Bound)
// algoritmoexploracion.AlDepthFirst(Depth)
// algoritmoexploracion.AlAmplitude (Amplitude)
func ExploreWithUninformed(
	graph [][]float64,
	solution solution.Solution,
	startState int,
	algorithmType explorationAlgorithm.ExpAlgorithmType,
) (path.Path, error) {

	var algorithm explorationAlgorithm.ExplorationAlgorithm
	switch algorithmType {
	case explorationAlgorithm.AlBranchAndBonud:
		algorithm = explorationAlgorithm.NewBranchAndBound(sortAlgorithm.NewMergeSort(), graph)
	case explorationAlgorithm.AlDepthFirst:
		algorithm = explorationAlgorithm.NewProfundidad()
	case explorationAlgorithm.AlAmplitude:
		algorithm = explorationAlgorithm.NewAmplitude()
	}
	explorer := explorer{
		graph:      graph,
		solution:   solution,
		algorithm:  algorithm,
		startState: startState,
	}
	return explorer.Explore()
}

// Explore a graph with an informed algorithm until it finds a solution state.
// graph is the graph to be explored
// solution must be a struct that implements the method IsSolution(int, [][]float64) and will indicate when a state is a solution to the problem
// heuristic is a struct that must implement the method HeuristicoEstado(int)float64, which will return a value
// quantitative of how good that state is in terms of the solution
// algorithmType is the scan algorithm to use. Your options are:
//
//	algoritmoexploracion.AlAStar (A*)
//	algoritmoexploracion.AlClimbing (Climbing)
//	algoritmoexploracion.AlBestFirst (Best First)
func ExploreWithInformed(
	graph [][]float64,
	solution solution.Solution,
	heuristic heuristic.StateHeuristic,
	startState int,
	algorithmType explorationAlgorithm.ExpAlgorithmType,
) (path.Path, error) {

	var algorithm explorationAlgorithm.ExplorationAlgorithm
	switch algorithmType {
	case explorationAlgorithm.AlAStar:
		algorithm = explorationAlgorithm.NewAStar(sortAlgorithm.NewMergeSort(), graph, heuristic)
	case explorationAlgorithm.AlClimbing:
		algorithm = explorationAlgorithm.NewEscalada(sortAlgorithm.NewMergeSort(), graph, heuristic)
	case explorationAlgorithm.AlBestFirst:
		algorithm = explorationAlgorithm.NewPrimeroElMejor(sortAlgorithm.NewMergeSort(), graph, heuristic)
	}
	explorer := explorer{
		graph:      graph,
		solution:   solution,
		algorithm:  algorithm,
		startState: startState,
	}
	return explorer.Explore()
}
