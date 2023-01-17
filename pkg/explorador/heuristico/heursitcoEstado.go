package heuristico

// Heuristico estado evalua los estados y devuelve un valor comprabale de que tan prometedores son
// de cara a llegar a la solución
type HeuristicoEstado interface {
	Heuristico(int) float64
}
