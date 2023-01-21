package algoritmoexploracion

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"

// Interface that covers the various system-dependent aspects of an exploration problem.
// This is defined by:
// The way in which it orders the pending paths to be explored
// How to define when a state is solution

type ExplorationAlgorithm interface {
	Merge(viejos *[]path.Path, nuevos *[]path.Path)
}

// These are the different types of algorithm that there are:
type ExpAlgorithmType int

const (
	AlBranchAndBonud = iota
	AlAStar
	AlClimbing
	AlBestFirst
	AlDepthFirst
	AlAmplitude
)
