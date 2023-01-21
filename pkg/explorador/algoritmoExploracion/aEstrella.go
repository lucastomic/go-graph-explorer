package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

//TODO: IMPLEMENTAR FACTOR ALPHA

// El algoritmo A* es un algoritmo que se basa en la estimación del coste total que tiene el camino.
// Esta estimación es conseguida mediante la suma del coste del camino ya conocido, y la estimacion de lo que queda
// hasta le meta (mediante el heuristico)
func NewAEstrella(
	algoritmoOrdenacion algoritmoOrdenacion.AlgoritmoOrdenacion,
	grafo [][]float64,
	heuristico heuristico.HeuristicoEstado,
) BusquedaOptimo {
	return BusquedaOptimo{
		algoritmoOrdenacion: algoritmoOrdenacion,
		grafo:               grafo,
		heuristico: heuristicoAEstrella{
			grafo:                grafo,
			heuristicoEstimacion: heuristico,
		},
	}
}

type heuristicoAEstrella struct {
	// heuristicoEstimacion es el heuristico que estima cuanto costará llegar desde un estado a la solución
	heuristicoEstimacion heuristico.HeuristicoEstado
	grafo                [][]float64
}

// Devuelve la suma entre el coste del camino ya recorrido (ya lo sabemos) y una estimaion del coste
// de camino que faltará por recorrer (estimacion aproximada en base al heuristico pasado)
func (h heuristicoAEstrella) Heuristico(camino path.Path) float64 {
	costeCaminoYaRecorrido := camino.GetTotalCost(h.grafo)
	estimacionCosteFaltante := h.heuristicoEstimacion.Heuristico(camino.GetCurrentState())
	return estimacionCosteFaltante + costeCaminoYaRecorrido
}
