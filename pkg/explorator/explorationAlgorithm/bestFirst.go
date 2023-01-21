package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

// Explration algorithm "Best First"
type BestFirst struct {
	heuristic heuristico.StateHeuristic
	sortAlg   algoritmoOrdenacion.SortAlgorithm
	graph     [][]float64
}

// Struct constructor
func NewPrimeroElMejor(
	sortAlg algoritmoOrdenacion.SortAlgorithm,
	graph [][]float64,
	heuristic heuristico.StateHeuristic,
) BestFirst {
	return BestFirst{
		heuristic: heuristic,
		sortAlg:   sortAlg,
		graph:     graph,
	}
}

// First, the old paths pending exploration are concatenated with the newly expanded ones
// and then sort according to the heuristic and sort method
func (p BestFirst) Merge(olds *[]path.Path, news *[]path.Path) {
	*olds = append(*olds, *news...)
	pathHeuristic := heuristico.NewPathHeurFromStateHeur(p.heuristic, p.graph)
	p.sortAlg.Sort(olds, pathHeuristic)
}
