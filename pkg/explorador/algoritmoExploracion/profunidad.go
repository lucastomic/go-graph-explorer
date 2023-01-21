package algoritmoexploracion

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"

// Explora un grafo en profundidad
type Profundidad struct {
}

// Reordena caminos pendientes por explorar poniendo primero los recién expandidos (los que incluyen los
// sucesores del estado acutal) y luego los que ya estaban
func (e Profundidad) Mezclar(viejos *[]path.Path, nuevos *[]path.Path) {
	*viejos = append(*nuevos, *viejos...)
}

// Constructor del struct
func NewProfundidad() Profundidad {
	return Profundidad{}
}
