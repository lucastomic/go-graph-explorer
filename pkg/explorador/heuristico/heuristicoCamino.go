package heuristico

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"

// Heuristico camino evalua un camino y devuelve un valor comprarble de que tan bueno es
type HeuristicoCamino interface {
	Heuristico(path.Path) float64
}

// Devuelve un heuristico de camino en base a un heuristico de estado.
// El heuristico de camino devuelto devuelve un float64 resultante de aplicar el heuristico
// de estado pasado por parametro al último estado del camino
func NewPathHeurFromStateHeur(stateHeur HeuristicoEstado, grafo [][]float64) HeuristicoCamino {
	return PathHeurFromStateHeur{
		stateHeur: stateHeur,
		grafo:     grafo,
	}
}

// Heuristico de camino sacado a partir de un heuristicoEstado
type PathHeurFromStateHeur struct {
	stateHeur HeuristicoEstado
	grafo     [][]float64
}

// Devuelve el valor del heuristico (heuristicoEstado) del estado final del camino pasado por parametro.
func (p PathHeurFromStateHeur) Heuristico(camino path.Path) float64 {
	currentState := camino.GetCurrentState()
	return p.stateHeur.Heuristico(currentState)
}
