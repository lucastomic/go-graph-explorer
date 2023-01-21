package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/prune"
)

// They are algorithms that guarantee to obtain the optimal path
type OptimalSerch struct {
	sortAlgorithm algoritmoOrdenacion.SortAlgorithm
	heuristic     heuristico.PathHeuristic
	graph         [][]float64
}

// It concatenates the old paths with the new ones.
// Then prune them and finally sort all his elemnets.
// It sorts them by cost.
func (b OptimalSerch) Merge(olds *[]path.Path, news *[]path.Path) {
	*olds = append(*olds, *news...)
	prune.Prune(olds, b.graph)
	b.sortAlgorithm.Sort(olds, b.heuristic)
}
