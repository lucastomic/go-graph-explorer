package algoritmoexploracion

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"

// Explore a graph in breadth.
// Alogirthm not reported.
// It does not return the optimal solution, but the first one it finds.
type Amplitude struct{}

// Reorders pending paths to be explored, putting the ones that were already there first
// and then the newly expanded ones (the ones that include the successors of the current state)
func (e Amplitude) Merge(olds *[]path.Path, news *[]path.Path) {
	*olds = append(*olds, *news...)
}

// struct constructor
func NewAmplitude() Amplitude {
	return Amplitude{}
}
