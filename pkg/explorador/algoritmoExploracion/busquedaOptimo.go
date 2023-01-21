package algoritmoexploracion

import (
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/prune"
)

// Son algoritmos que garantizan obtener el camino óptimo
type BusquedaOptimo struct {
	algoritmoOrdenacion algoritmoOrdenacion.AlgoritmoOrdenacion
	heuristico          heuristico.HeuristicoCamino
	grafo               [][]float64
}

// It concatenates the old paths with the new ones.
// Then prune them and finally sort all his elemnets.
// It sorts them by cost.
func (b BusquedaOptimo) Mezclar(viejos *[]path.Path, nuevos *[]path.Path) {
	*viejos = append(*viejos, *nuevos...)
	prune.Prune(viejos, b.grafo)
	b.algoritmoOrdenacion.Ordenar(viejos, b.heuristico)
}
