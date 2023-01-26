package explorationAlgorithm

import "github.com/lucastomic/go-graph-explorer/pkg/explorator/path"

// Interface that covers the various system-dependent aspects of an exploration problem.
// This is defined by:
// The way in which it orders the pending paths to be explored
// How to define when a state is solution

type ExplorationAlgorithm interface {
	Merge(*[]path.Path, *[]path.Path)
}
