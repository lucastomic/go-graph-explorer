package explorationAlgorithm

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/heuristic"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/prune"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/sortAlgorithm"
)

// They are algorithms that guarantee to obtain the optimal path
type optimalSerch struct {
	sortAlgorithm sortAlgorithm.SortAlgorithm
	heuristic     heuristic.PathHeuristic
	graph         [][]float64
}

// It concatenates the old paths with the new ones.
// Then prune them and finally sort all his elemnets.
// It sorts them by cost.
func (b optimalSerch) Merge(olds *[]path.Path, news *[]path.Path) {
	*olds = append(*olds, *news...)
	prune.Prune(olds, b.graph)
	b.sortAlgorithm.Sort(olds, b.heuristic)
}
