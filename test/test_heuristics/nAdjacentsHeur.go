package testheuristics

import floatutils "github.com/lucastomic/ExploracionDeEspacios/pkg/utils/floatUtils"

// Heuristic which returns the number of adjacents
type NAdjacentsHeur struct{ testGraph [][]float64 }

func (h NAdjacentsHeur) Heuristic(state int) float64 {
	var adyacentsAmount float64
	for _, cost := range h.testGraph[state] {
		if floatutils.NotInfinit(cost) {
			adyacentsAmount++
		}
	}
	return adyacentsAmount
}

func NewNAdjacentHeur(graph [][]float64) NAdjacentsHeur {
	return NAdjacentsHeur{graph}
}
