package algoritmoOrdenacion

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

// Different vector sorting algorithms.
type SortAlgorithm interface {
	Sort(*[]path.Path, heuristico.PathHeuristic)
}
