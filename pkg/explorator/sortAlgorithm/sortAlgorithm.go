package sortAlgorithm

import (
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/heuristic"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/path"
)

// Different vector sorting algorithms.
type SortAlgorithm interface {
	//Sorts the paths slice in an ascending way depending on the heuristic passed as argument
	Sort(*[]path.Path, heuristic.PathHeuristic)
}
