package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
)

// Algoritmo de exploracion "Escalada"
type Escalada struct {
	heuristico heuristico.Heuristico
	algoritmoOrdenacion.AlgoritmoOrdenacion
}

// Primero ordena los nuevos algoritmos obtenidos con la expanción del estado actual, y
// luego los concatena con los demás caminos pendientes por explorar.
func (e Escalada) Mezclar(nuevos *[][]int, viejos *[][]int) {
	e.Ordenar(nuevos, e.heuristico)
	*viejos = append(*viejos, *nuevos...)
}
