package algoritmoOrdenacion

import (
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/heuristico"
	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

type MergeSort struct {
}

// Returns an object of the sorting algorithm "Merge Sort"
// heuristic is the heuristic with which you will compare the paths to order them among themselves
func NewMergeSort() MergeSort {
	return MergeSort{}
}

// Sort the paths given the heuristic passed into the constructor
func (m MergeSort) Sort(caminos *[]path.Path, heuristico heuristico.PathHeuristic) {
	m.mergeSortAux(caminos, 0, len(*caminos)-1, heuristico)
}

// Given a vector of paths that includes two subvectors ordered by the indicated heuristic,
// from i0 to k the first subvector, and from k to iN the second, orders the vector from i0 to iN depending on the indicated heuristic.

// For example, given the vector of paths with the following heuristics:
// [...,-2,0,1,4,19,-5,-1,4,7,...]
// where i0 would be the position occupied by the -2 element
// k would be the position occupied by element 19
// iN would be the position occupied by element 7

// The function merge(&vector, indexDe(-2),indexDe(19), indexDe(7))
// sort the array like this:
// [...,-5,-2,-1,0,1,4,4,7,19,...]
func (m MergeSort) merge(vectorPointer *[]path.Path, i0 int, k int, iN int, heuristic heuristico.PathHeuristic) {
	i, d, f := i0, k+1, 0
	aux := make([]path.Path, iN-i0+1)
	vector := *vectorPointer
	for i <= k && d <= iN {
		if heuristic.Heuristic(vector[i]) <= heuristic.Heuristic(vector[d]) {
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

// As long as i0 and iN are not the same position, look for an intermediate position and use recursion to sort both sides of the partition.
// Once both parts are sorted, use the merge method to sort the vector from i0 to iN starting from the two sorted subVectors.
// If i0 and iN are equal, we are in the base case, so it does not alter the vector and leaves it as it was.
func (m MergeSort) mergeSortAux(vector *[]path.Path, i0 int, iN int, heuristic heuristico.PathHeuristic) {
	if i0 < iN {
		k := (i0 + iN) / 2
		m.mergeSortAux(vector, i0, k, heuristic)
		m.mergeSortAux(vector, k+1, iN, heuristic)
		m.merge(vector, i0, k, iN, heuristic)
	}
}
