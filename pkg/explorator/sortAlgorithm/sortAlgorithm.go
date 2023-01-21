package sortAlgorithm

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/heuristic"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
)

// Different vector sorting algorithms.
type SortAlgorithm interface {
	Sort(*[]path.Path, heuristic.PathHeuristic)
}
