package algoritmoexploracion

// Explora un grafo en amplitud.
// Alogirtmo no informado.
// No devuelve solución optima, sino la primera que encuentra.
type AlgoritmoEnAmplitud struct {
}

// Reordena caminos pendientes por explorar poniendo primero los que ya estaban
// y luego los recién expandidos (los que incluyen los sucesores del estado acutal)
func (e AlgoritmoEnAmplitud) Mezclar(viejos *[][]int, nuevos *[][]int) {
	*viejos = append(*viejos, *nuevos...)
}

// constructor del struct
func newAmplitud() AlgoritmoEnAmplitud {
	return AlgoritmoEnAmplitud{}
}
