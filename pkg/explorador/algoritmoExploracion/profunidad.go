package algoritmoexploracion

// Explora un grafo en profundidad
type AlgoritmoEnProfundidad struct {
}

// Reordena caminos pendientes por explorar poniendo primero los recién expandidos (los que incluyen los
// sucesores del estado acutal) y luego los que ya estaban
func (e AlgoritmoEnProfundidad) Mezclar(viejos *[][]int, nuevos *[][]int) {
	*viejos = append(*nuevos, *viejos...)
}

// Constructor del struct
func newProfundidad() AlgoritmoEnProfundidad {
	return AlgoritmoEnProfundidad{}
}
