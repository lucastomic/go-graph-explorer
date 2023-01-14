package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
)

// Algoritmo de exploración "Primero el mejor"
type PrimeroElMejor struct {
	heuristico heuristico.Heuristico
	algoritmoOrdenacion.AlgoritmoOrdenacion
}

// Primero se concatenan los caminos pendientes de explorar viejos con los recien expandidos
// y luego se ordena según el heuristico y el método de ordenación
func (p PrimeroElMejor) Mezclar(viejos *[][]int, nuevos *[][]int) {
	*viejos = append(*viejos, *nuevos...)
	p.Ordenar(viejos, p.heuristico)
}
