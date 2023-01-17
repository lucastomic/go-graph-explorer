package valoradorcamino

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/camino"

type ValoradorCamino interface {
	//Método que devuelve el valor que se le asigna a cada camino para ser comparado con los otros.
	//Por ejemplo, si en el algoritmo se comparan por el coste, el método deverá devolver el coste de ese camino.
	ObtenerValorComparable(camino camino.Camino, grafo [][]float64) float64
}
