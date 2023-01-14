package algoritmoOrdenacion

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"

// Diferentes algoritmos de ordenación de vectores.
type AlgoritmoOrdenacion interface {
	Ordenar(caminos *[][]int, heuristico heuristico.Heuristico)
}
