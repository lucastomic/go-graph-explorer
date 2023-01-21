package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

// Algoritmo de exploración "Primero el mejor"
type PrimeroElMejor struct {
	heuristico   heuristico.HeuristicoEstado
	alOrdenacion algoritmoOrdenacion.AlgoritmoOrdenacion
	grafo        [][]float64
}

// Constructor de la clase
func NewPrimeroElMejor(
	alOrdenacion algoritmoOrdenacion.AlgoritmoOrdenacion,
	grafo [][]float64,
	heuristico heuristico.HeuristicoEstado,
) PrimeroElMejor {
	return PrimeroElMejor{
		heuristico:   heuristico,
		alOrdenacion: alOrdenacion,
		grafo:        grafo,
	}
}

// Primero se concatenan los caminos pendientes de explorar viejos con los recien expandidos
// y luego se ordena según el heuristico y el método de ordenación
func (p PrimeroElMejor) Mezclar(viejos *[]path.Path, nuevos *[]path.Path) {
	*viejos = append(*viejos, *nuevos...)
	heuristicoCamino := heuristico.NewPathHeurFromStateHeur(p.heuristico, p.grafo)
	p.alOrdenacion.Ordenar(viejos, heuristicoCamino)
}
