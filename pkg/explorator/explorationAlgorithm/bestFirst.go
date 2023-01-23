package explorationAlgorithm

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/heuristic"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/sortAlgorithm"
)

// Explration algorithm "Best First"
type BestFirst struct {
	heuristic heuristic.StateHeuristic
	sortAlg   sortAlgorithm.SortAlgorithm
	graph     [][]float64
}

// Struct constructor
func NewBestFirst(
	sortAlg sortAlgorithm.SortAlgorithm,
	graph [][]float64,
	heuristic heuristic.StateHeuristic,
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
	pathHeuristic := heuristic.NewPathHeurFromStateHeur(p.heuristic, p.graph)
	p.sortAlg.Sort(olds, pathHeuristic)
}
