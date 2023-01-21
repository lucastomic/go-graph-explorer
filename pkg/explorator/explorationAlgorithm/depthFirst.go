package algoritmoexploracion

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"

// "DepthFirst" exploration algorithm
type DepthFirst struct {
}

// Reorders pending paths to be explored putting the newly expanded ones first (the ones that include the
// successors of the current state) and then those that were already
func (e DepthFirst) Merge(olds *[]path.Path, news *[]path.Path) {
	*olds = append(*news, *olds...)
}

// Constructor
func NewProfundidad() DepthFirst {
	return DepthFirst{}
}
