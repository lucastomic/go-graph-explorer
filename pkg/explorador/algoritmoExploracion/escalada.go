package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/camino"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
)

// Algoritmo de exploracion "Escalada"
type Escalada struct {
	heuristico   heuristico.HeuristicoEstado
	alOrdenacion algoritmoOrdenacion.AlgoritmoOrdenacion
	grafo        [][]float64
}

// Constructor del algoritmo de exploración "Escalada".
// heuristico es el heuristico de estado mediante el cual el algoritmo evalua los estados para tomar decisiones de expansión.
// alOrdenacion es el algoritmo de ordenacion mediante el cual "escalada" ordena los caminos una vez son evaluados con el heuristico
func NewEscalada(heuristico heuristico.HeuristicoEstado, alOrdenacion algoritmoOrdenacion.AlgoritmoOrdenacion) Escalada {
	return Escalada{
		heuristico:   heuristico,
		alOrdenacion: alOrdenacion,
	}
}

// Primero ordena los nuevos algoritmos obtenidos con la expanción del estado actual, y
// luego los concatena con los demás caminos pendientes por explorar.
func (e Escalada) Mezclar(nuevos *[]camino.Camino, viejos *[]camino.Camino) {
	heuristicoCamino := heuristico.NewPathHeurFromStateHeur(e.heuristico, e.grafo)
	e.alOrdenacion.Ordenar(nuevos, heuristicoCamino)
	*viejos = append(*viejos, *nuevos...)
}
