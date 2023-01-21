package explorationAlgorithm

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/heuristic"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/sortAlgorithm"
)

// Exploration algorithm "Climbing"
type Climbing struct {
	heuristic heuristic.StateHeuristic
	sortAlg   sortAlgorithm.SortAlgorithm
	graph     [][]float64
}

// Constructor of the "Escalation" scanning algorithm.
// heuristic is the state heuristic by which the algorithm evaluates the states to make expansion decisions.
// alSort is the sort algorithm by which "escalation" sorts the paths once they are evaluated with the heuristic
func NewEscalada(
	sortAlg sortAlgorithm.SortAlgorithm,
	graph [][]float64,
	heuristic heuristic.StateHeuristic,
) Climbing {
	return Climbing{
		heuristic: heuristic,
		sortAlg:   sortAlg,
		graph:     graph,
	}
}

// First sorts the new algorithms obtained by expanding the current state, and
// then concatenates them with the remaining paths to be explored.
func (e Climbing) Merge(news *[]path.Path, olds *[]path.Path) {
	pathHeuristic := heuristic.NewPathHeurFromStateHeur(e.heuristic, e.graph)
	e.sortAlg.Sort(news, pathHeuristic)
	*olds = append(*olds, *news...)
}
