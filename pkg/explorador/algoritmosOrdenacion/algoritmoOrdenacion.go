package algoritmoOrdenacion

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

// Diferentes algoritmos de ordenación de vectores.
type AlgoritmoOrdenacion interface {
	Ordenar(caminos *[]path.Path, heuristico heuristico.HeuristicoCamino)
}
