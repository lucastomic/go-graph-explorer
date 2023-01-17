package camino

import (
	"math"

	"github.com/lucastomic/ExploracionDeEspacios/internals/sliceUtils"
)

// Un camino es una sucesión aciclica de estados. Estos estados son representados
// por un Int que coincide con el indice que los representa en la matriz de adyacencia del grafo.
type Camino struct {
	estados *[]int
}

// Returns a path with the states specified by argument
func NewCamino(estados *[]int) Camino {
	return Camino{estados: estados}
}

// Returns a empty path
func NewEmptyPath() Camino {
	states := &[]int{}
	return Camino{estados: states}
}

// Devuelve el coste total que tiene el camino
// Para ello hace una iteración de los costes de todas las transiciones.
// Las transiciones son los cambios de un estado a otro.
func (c Camino) GetTotalCost(grafo [][]float64) float64 {
	var costeTotal float64
	//Iteramos transiciones de estados. Hasta len(camino)-1 porque el último estado no hace transicion a ningún otro estado.
	for i := 0; i < len(*c.estados)-1; i++ {
		//i:estado actual
		//i+1: estado siguiente
		costeTotal += grafo[i][i+1]
	}

	return costeTotal
}

// getWithNewState Returns a path with the same states but with the state passed as an argument at the end.
// It doesn't modfy the actually states.
// For exmale, having a path with the next states:
// [0,4,3,9,6]
// getWithNewState(5) would return a different path with the next states:
// [0,4,3,9,6,5]
func (c Camino) getWithNewState(newState int) Camino {
	states := append(*c.estados, newState)
	return NewCamino(&states)
}

// Convierte camino parcial en un conjunto de nuevos caminos, resultado de añadir a éste cada uno
// de los posibles sucesores que no den lugar a caminos cíclicos
func (c Camino) Expandir(grafo [][]float64) []Camino {
	var res []Camino
	estados := *c.estados
	estadoActual := estados[len(estados)-1]
	adyacentesDeEstadoActual := grafo[estadoActual]

	// Por cada estado adyacente al acual (el último de la lista) que no esté en el camino,
	// se agrega un camino a la respuesta con el camino+estado adyacente
	for i := range adyacentesDeEstadoActual {
		if c.diferenteDeInfinito(adyacentesDeEstadoActual[i]) && !sliceUtils.Contains(*c.estados, i) {
			res = append(res, c.getWithNewState(i))
		}
	}
	return res
}

// Si [val] es diferente a infinito (Esto en el grafo quiere decir que están conectados)
func (c Camino) diferenteDeInfinito(val float64) bool {
	return val != math.MaxFloat64
}

// Devuelve el último estado del camino (el estado actual)
func (c Camino) GetCurrentState() int {
	states := *c.estados
	return states[len(*c.estados)-1]
}
