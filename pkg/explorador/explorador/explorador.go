package explorador

import (
	"errors"

	"github.com/lucastomic/ExploracionDeEspacios/internals/sliceUtils"
	algoritmoexploracion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmoExploracion"
	algoritmoOrdenacion "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/algoritmosOrdenacion"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/solucion"
)

// explorador es una estructura encargada de la exploracion de un grafo.
// Deben pasaras por parametro:
// El grafo que será evaluado en forma de matriz de adyacencia
// El algoritmo de exploracion es definido por un objeto que implemente la interfaz algoritmoExploracion, el cual
// define como se ordenaran los caminos pendientes por explorar una vez se extiend el estado actual.
// El estado solución se mide con el méotdo esSolucion() de la estructura Solucion
type explorador struct {
	grafo     [][]float64
	algoritmo algoritmoexploracion.AlgoritmoExploracion
	solucion  solucion.Solucion
}

// Indica si el explorador debe seguir buscando un camino optimo.
// Esto será asi siempre y cuando todavía queden caminos pendientes por los que buscar y
// no se haya encontrado aun la solucion
func (e explorador) seguirBuscando(caminosPendientes []path.Path, caminoActual path.Path) bool {
	return !sliceUtils.IsEmpty(caminosPendientes) && !e.solucion.EsSolucion(caminoActual.GetCurrentState(), e.grafo)
}

// Explora el grafo obteniendo el camnio buscado.
// Para ello crea una lista de caminos pendientes, empezando por el primer nodo según la matriz de adyacencia.
// Luego, en cada iteracion reemplaza el último camino con lo devuelto por el método [e.expandir()] y
// reordena los caminos dependiendo del método [mezclar()] del algoritmo de exploracion pasado por argumento.
// Itera hasta que el método [e.seguirBuscando()] devuelve false
// Si se terminaron todas las iteraciones y el estado actual no es la solucion, devuelve un error indicando que le problema no tiene solucion.
// En caso contrario, devuelve el estado actual.
func (e explorador) Explorar() (path.Path, error) {
	var caminosPendientes []path.Path
	caminosPendientes[0] = path.NewEmptyPath()
	caminoActual := caminosPendientes[len(caminosPendientes)-1]

	for e.seguirBuscando(caminosPendientes, caminoActual) {
		caminosNuevos := caminoActual.Expand(e.grafo)
		sliceUtils.EliminarUltimo(caminosNuevos)

		e.algoritmo.Mezclar(&caminosPendientes, caminosNuevos)
		caminoActual = caminosPendientes[len(caminosPendientes)-1]
	}

	if !e.solucion.EsSolucion(caminoActual.GetCurrentState(), e.grafo) {
		return path.NewEmptyPath(), errors.New("no existe solucion")
	} else {
		return caminoActual, nil
	}
}

// Explora un grafo con un algoritmo no informado hasta que encuentra un estado solución.
// grafo es el grafo que se explorará
// solucion debe ser un struct que implemente el método EsSolucion(int, [][]float64) e indicará cuando un estado es solución del problema
// tipoAlgoritmo es el algoritmo de exploración que se utilizará. Sus opciones son:
//
//	algoritmoexploracion.AlBranchAndBonud (Branch&Bound)
//	algoritmoexploracion.AlProfundidad (Profunidad)
//	algoritmoexploracion.AlAmplitud (Amplitud)
func ExplorarNoInformado(
	grafo [][]float64,
	solucion solucion.Solucion,
	tipoAlgoritmo algoritmoexploracion.TipoAlgoritmoExp,
) (path.Path, error) {

	var algoritmo algoritmoexploracion.AlgoritmoExploracion
	switch tipoAlgoritmo {
	case algoritmoexploracion.AlBranchAndBonud:
		algoritmo = algoritmoexploracion.NewBranchAndBound(algoritmoOrdenacion.NewMergeSort(), grafo)
	case algoritmoexploracion.AlProfundidad:
		algoritmo = algoritmoexploracion.NewProfundidad()
	case algoritmoexploracion.AlAmplitud:
		algoritmo = algoritmoexploracion.NewAmplitud()
	}
	explorer := explorador{
		grafo:     grafo,
		solucion:  solucion,
		algoritmo: algoritmo,
	}
	return explorer.Explorar()
}

// Explora un grafo con un algoritmo informado hasta que encuentra un estado solución.
// grafo es el grafo que se explorará
// solucion debe ser un struct que implemente el método EsSolucion(int, [][]float64) e indicará cuando un estado es solución del problema
// heuristico es un struct que debe implementar el método HeuristicoEstado(int)float64, el cual devolverá un valor
// cuantitativo de que tan bueno es ese estado de cara a la solución
// tipoAlgoritmo es el algoritmo de exploración que se utilizará. Sus opciones son:
//
//	algoritmoexploracion.AlAEstrella (A*)
//	algoritmoexploracion.AlEscalada (Escalada)
//	algoritmoexploracion.AlPrimeroElMejor (Primero el mejor)
func ExplorarInformado(
	grafo [][]float64,
	solucion solucion.Solucion,
	heuristico heuristico.HeuristicoEstado,
	tipoAlgoritmo algoritmoexploracion.TipoAlgoritmoExp,
) (path.Path, error) {

	var algoritmo algoritmoexploracion.AlgoritmoExploracion
	switch tipoAlgoritmo {
	case algoritmoexploracion.AlAEstrella:
		algoritmo = algoritmoexploracion.NewAEstrella(algoritmoOrdenacion.NewMergeSort(), grafo, heuristico)
	case algoritmoexploracion.AlEscalada:
		algoritmo = algoritmoexploracion.NewEscalada(algoritmoOrdenacion.NewMergeSort(), grafo, heuristico)
	case algoritmoexploracion.AlPrimeroElMejor:
		algoritmo = algoritmoexploracion.NewPrimeroElMejor(algoritmoOrdenacion.NewMergeSort(), grafo, heuristico)
	}
	explorer := explorador{
		grafo:     grafo,
		solucion:  solucion,
		algoritmo: algoritmo,
	}
	return explorer.Explorar()
}
