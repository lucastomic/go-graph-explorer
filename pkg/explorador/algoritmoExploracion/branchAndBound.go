package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
)

// Algoritmo de exploracion "Branch And Bound"
// No informado
// Encuentra la solcuión optim
// Se basa en elegir siempre para expandir el camino parcial
// de menor coste. Esto garantiza que, cuando el camino
// seleccionado sea una solución al problema, será
// precisamente de menor costeá
type BranchAndBound struct {
	heuristico heuristico.Heuristico
	algoritmoOrdenacion.AlgoritmoOrdenacion
}

func (b BranchAndBound) Mezclar(viejos *[][]int, nuevos *[][]int) {

}
