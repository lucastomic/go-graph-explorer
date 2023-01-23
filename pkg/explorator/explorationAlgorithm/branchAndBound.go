package explorationAlgorithm

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/sortAlgorithm"
)

// Heuristic used in the Branch&Bound algorithm
type branchAndBoundHeuristic struct {
	graph [][]float64
}

// Returns the total cost of a path
func (b branchAndBoundHeuristic) Heuristic(path path.Path) float64 {
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
	sortAlgorithm sortAlgorithm.SortAlgorithm,
	graph [][]float64,
) OptimalSerch {
	return OptimalSerch{
		sortAlgorithm: sortAlgorithm,
		graph:         graph,
		heuristic: branchAndBoundHeuristic{
			graph: graph,
		},
	}
}
