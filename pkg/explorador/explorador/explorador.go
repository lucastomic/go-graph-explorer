package explorador

import (
	"errors"

	"github.com/lucastomic/ExploracionDeEspacios/internals/sliceUtils"
	algoritmoexploracion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmoExploracion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/camino"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/solucion"
)

// Explorador es una estructura encargada de la exploracion de un grafo.
// Deben pasaras por parametro:
// El grafo que será evaluado en forma de matriz de adyacencia
// El algoritmo de exploracion es definido por un objeto que implemente la interfaz algoritmoExploracion, el cual
// define como se ordenaran los caminos pendientes por explorar una vez se extiend el estado actual.
// El estado solución se mide con el méotdo esSolucion() de la estructura Solucion
type Explorador struct {
	grafo [][]float64
	algoritmoexploracion.AlgoritmoExploracion
	solucion.Solucion
}

// Indica si el explorador debe seguir buscando un camino optimo.
// Esto será asi siempre y cuando todavía queden caminos pendientes por los que buscar y
// no se haya encontrado aun la solucion
func (e Explorador) seguirBuscando(caminosPendientes []camino.Camino, caminoActual camino.Camino) bool {
	return !sliceUtils.IsEmpty(caminosPendientes) && !e.EsSolucion(caminoActual)
}

// Explora el grafo obteniendo el camnio optimo.
// Para ello crea una lista de caminos pendientes, empezando por el primer nodo según la matriz de adyacencia.
// Luego, en cada iteracion reemplaza el último camino con lo devuelto por el método [e.expandir()] y
// reordena los caminos dependiendo del método [mezclar()] del algoritmo de exploracion pasado por argumento.
// Itera hasta que el método [e.seguirBuscando()] devuelve false
// Si se terminaron todas las iteraciones y el estado actual no es la solucion, devuelve un error indicando que le problema no tiene solucion.
// En caso contrario, devuelve el estado actual.
func (e Explorador) Explorar() (camino.Camino, error) {
	var caminosPendientes []camino.Camino
	//Inicializamos camino vacío
	caminosPendientes[0] = camino.NewEmptyPath()
	caminoActual := caminosPendientes[len(caminosPendientes)-1]

	for e.seguirBuscando(caminosPendientes, caminoActual) {
		caminosNuevos := caminoActual.Expandir(e.grafo)
		sliceUtils.EliminarUltimo(&caminosNuevos)

		e.Mezclar(&caminosPendientes, &caminosNuevos)
		caminoActual = caminosPendientes[len(caminosPendientes)-1]
	}

	if !e.EsSolucion(caminoActual) {
		return camino.NewEmptyPath(), errors.New("no existe solucion")
	} else {
		return caminoActual, nil
	}
}
