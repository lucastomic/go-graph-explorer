package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/camino"
)

// Algoritmo de exploracion "Branch And Bound"
// No informado
// Encuentra la solcuión optimo
// Utiliza poda
// Se basa en elegir siempre para expandir el camino parcial
// de menor coste. Esto garantiza que, cuando el camino
// seleccionado sea una solución al problema, será
// precisamente la de menor coste

// Heuristico que se utiliza en el algoritmo de Branch&Bound
type branchAndBoundHeuristico struct {
	grafo [][]float64
}

// Devuelve el coste total de un camino
func (b branchAndBoundHeuristico) Heuristico(camino camino.Camino) float64 {
	return camino.GetTotalCost(b.grafo)
}

// Constructor de la clase. Devuelve un algoritmo de busqueda ópitmo, pasando como heuristico el coste del camino.
func newBranchAndBound(algoritmoOrdenacion algoritmoOrdenacion.AlgoritmoOrdenacion, grafo [][]float64) BusquedaOptimo {
	return BusquedaOptimo{
		algoritmoOrdenacion: algoritmoOrdenacion,
		grafo:               grafo,
		heuristico: branchAndBoundHeuristico{
			grafo: grafo,
		},
	}
}
