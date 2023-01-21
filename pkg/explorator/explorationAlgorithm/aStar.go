package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

//TODO: IMPLEMENTAR FACTOR ALPHA

// Algorithm A* is an algorithm that is based on the estimation of the total cost of the path.
// This estimate is obtained by adding the cost of the path already known, and the estimate of what remains
// up to the goal (using the heuristic)
func NewAEstrella(
	sortAlgorithm algoritmoOrdenacion.SortAlgorithm,
	graph [][]float64,
	heuristic heuristico.StateHeuristic,
) OptimalSerch {
	return OptimalSerch{
		sortAlgorithm: sortAlgorithm,
		graph:         graph,
		heuristic: AStarHeuristic{
			graph:               graph,
			estimationHeuristic: heuristic,
		},
	}
}

type AStarHeuristic struct {
	// estimation Heuristic is the heuristic that estimates how long it will take to get from a state to the solution
	estimationHeuristic heuristico.StateHeuristic
	graph               [][]float64
}

// Returns the sum between the cost of the path already traveled (we already know it) and an estimate of the cost
// of path that will remain to be traveled (rough estimate based on the past heuristic)
func (h AStarHeuristic) Heuristic(path path.Path) float64 {
	costPathAlreadyExplored := path.GetTotalCost(h.graph)
	missingCostEstimation := h.estimationHeuristic.Heuristic(path.GetCurrentState())
	return missingCostEstimation + costPathAlreadyExplored
}