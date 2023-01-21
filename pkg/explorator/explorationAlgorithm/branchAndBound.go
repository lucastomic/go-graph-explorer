package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

// Heuristic used in the Branch&Bound algorithm
type branchAndBoundHHeuristic struct {
	graph [][]float64
}

// Returns the total cost of a path
func (b branchAndBoundHHeuristic) Heuristic(path path.Path) float64 {
	return path.GetTotalCost(b.graph)
}

// "Branch And Bound" scanning algorithm
// Not informed
// Find the optimal solution
// Use pruning
// Relies on always choosing to expand partial path
// lowest cost. This ensures that when the road
// selected is a solution to the problem, it will be
// precisely the one with the lowest cost
// Constructor of the class. Returns an optimal search algorithm, passing the cost of the path as the heuristic.
func NewBranchAndBound(
	sortAlgorithm algoritmoOrdenacion.SortAlgorithm,
	graph [][]float64,
) OptimalSerch {
	return OptimalSerch{
		sortAlgorithm: sortAlgorithm,
		graph:         graph,
		heuristic: branchAndBoundHHeuristic{
			graph: graph,
		},
	}
}
