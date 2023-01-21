package heuristico

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"

// Heuristic path evaluates a path and returns a comparable value of how good it is
type PathHeuristic interface {
	Heuristic(path.Path) float64
}

// Returns a path heuristic based on a state heuristic.
// The returned path heuristic returns a float64 resulting from applying the heuristic
// from state passed by parameter to the last state of the path
func NewPathHeurFromStateHeur(stateHeur StateHeuristic, grafo [][]float64) PathHeuristic {
	return PathHeurFromStateHeur{
		stateHeur: stateHeur,
		graph:     grafo,
	}
}

// Path heuristic derived from a stateheuristic
type PathHeurFromStateHeur struct {
	stateHeur StateHeuristic
	graph     [][]float64
}

// Returns the value of the heuristic (State heuristic) of the final state of the path passed by parameter.
func (p PathHeurFromStateHeur) Heuristic(path path.Path) float64 {
	currentState := path.GetCurrentState()
	return p.stateHeur.Heuristic(currentState)
}
