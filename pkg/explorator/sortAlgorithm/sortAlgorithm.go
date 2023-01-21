package sortAlgorithm

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/heuristic"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorator/path"
)

// Different vector sorting algorithms.
type SortAlgorithm interface {
	//Sorts the paths slice in an ascending way depending on the heuristic passed as argument
	Sort(*[]path.Path, heuristic.PathHeuristic)
}
