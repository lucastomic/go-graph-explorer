package algoritmoOrdenacion

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

type MergeSort struct {
}

// Devuelve un objeto del algoritmo de ordenacion "Merge Sort"
// heuristico es el heuristico con el que comprará los caminos para ordenarlos entre si
func NewMergeSort() MergeSort {
	return MergeSort{}
}

// Ordena los caminos dado el heuristico pasado en el constructor
func (m MergeSort) Ordenar(caminos *[]path.Path, heuristico heuristico.HeuristicoCamino) {
	m.mergeSortAux(caminos, 0, len(*caminos)-1, heuristico)
}

// Dado un vector de caminos que incluye dos subvectores ordenados por el heuristico indicado,
// desde i0 a k el primer subvector, y de k a iN el segundo, ordena el vector de i0 a iN dependiendo el heuristico indicado.

// Por ejemplo, dado el vector de caminos con los siguientes heuristico:
// [...,-2,0,1,4,19,-5,-1,4,7,...]
// donde i0 sería la posición que ocupa el elemento -2
// k sería la posición que ocupa el elemento 19
// iN sería la posición que ocupa el elemento 7

// La función merge(&vector, indexDe(-2),indexDe(19), indexDe(7))
// ordena el vector de esta forma:
// [...,-5,-2,-1,0,1,4,4,7,19,...]
func (m MergeSort) merge(vectorPointer *[]path.Path, i0 int, k int, iN int, heuristico heuristico.HeuristicoCamino) {
	i, d, f := i0, k+1, 0
	aux := make([]path.Path, iN-i0+1)
	vector := *vectorPointer
	for i <= k && d <= iN {
		if heuristico.Heuristico(vector[i]) <= heuristico.Heuristico(vector[d]) {
			aux[f] = vector[i]
			i++
		} else {
			aux[f] = vector[d]
			d++
		}
		f++
	}

	for a := i; a <= k; a++ {
		aux[f] = vector[a]
		f++
	}
	for a := d; a <= iN; a++ {
		aux[f] = vector[a]
		f++
	}
	for el := range aux {
		vector[i0+el] = aux[el]
	}
}

// Siempre y cuando i0 e iN no sean la misma posición, busca una posición intermedia y utiliza recursividad para ordenar ambos lados de la partición.
// Una vez ordenadas ambas partes, utliiza el método merge para ordenar el vector desde i0 a iN partiendo de los dos subVectores ordenados.
// Si i0 e iN son iguales, nos encontramos en el caso base, por lo que no altera el vector y lo deja como estaba.
func (m MergeSort) mergeSortAux(vector *[]path.Path, i0 int, iN int, heuristico heuristico.HeuristicoCamino) {
	if i0 < iN {
		k := (i0 + iN) / 2
		m.mergeSortAux(vector, i0, k, heuristico)
		m.mergeSortAux(vector, k+1, iN, heuristico)
		m.merge(vector, i0, k, iN, heuristico)
	}
}
