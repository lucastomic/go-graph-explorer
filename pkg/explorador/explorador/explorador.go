package explorador

import (
	"errors"
	"math"

	"github.com/lucastomic/ExploracionDeEspacios/internals/sliceUtils"
	algoritmoexploracion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmoExploracion"
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
	algoritmoexploracion.AlogritmoExploracion
	solucion.Solucion
}

// Si [val] es diferente a infinito (Esto en el grafo quiere decir que están conectados)
func (e Explorador) diferenteDeInfinito(val float64) bool {
	return val != math.MaxFloat64
}

// Convierte camino parcial en un conjunto de nuevos caminos, resultado de añadir a éste cada uno
// de los posibles sucesores que no den lugar a caminos cíclicos
func (e Explorador) expandir(grafo [][]float64, caminoParcial []int) [][]int {
	var res [][]int

	estadoActual := caminoParcial[len(caminoParcial)-1]
	adyacentesDeEstadoActual := grafo[estadoActual]

	for i := range adyacentesDeEstadoActual {
		if e.diferenteDeInfinito(adyacentesDeEstadoActual[i]) && !sliceUtils.Contains(caminoParcial, i) {
			res = append(res, append(caminoParcial, i))
		}
	}
	return res
}

// Indica si el explorador debe seguir buscando un camino optimo.
// Esto será asi siempre y cuando todavía queden caminos pendientes por los que buscar y
// no se haya encontrado aun la solucion
func (e Explorador) seguirBuscando(caminosPendientes [][]int, estadoActual []int) bool {
	return !sliceUtils.IsEmpty(caminosPendientes) && !e.EsSolucion(estadoActual)
}

// Explora el grafo obteniendo el camnio optimo.
// Para ello crea una lista de caminos pendientes, empezando por el primer nodo según la matriz de adyacencia.
// Luego, en cada iteracion reemplaza el último camino con lo devuelto por el método [e.expandir()] y
// reordena los caminos dependiendo del método [mezclar()] del algoritmo de exploracion pasado por argumento.
// Itera hasta que el método [e.seguirBuscando()] devuelve false
// Si se terminaron todas las iteraciones y el estado actual no es la solucion, devuelve un error indicando que le problema no tiene solucion.
// En caso contrario, devuelve el estado actual.
func (e Explorador) Explorar() ([]int, error) {
	var caminosPendientes [][]int
	caminosPendientes[0] = []int{0}
	estadoActual := caminosPendientes[len(caminosPendientes)-1]

	for e.seguirBuscando(caminosPendientes, estadoActual) {
		caminosNuevos := e.expandir(e.grafo, estadoActual)
		sliceUtils.EliminarUltimo(&caminosNuevos)
		e.Mezclar(&caminosPendientes, &caminosNuevos)
		estadoActual = caminosPendientes[len(caminosPendientes)-1]
	}

	if !e.EsSolucion(estadoActual) {
		return make([]int, 0), errors.New("no existe solucion")
	} else {
		return estadoActual, nil
	}
}
