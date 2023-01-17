package solucion

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/camino"

// Método que evalua si un estado es la solución del problema
type Solucion interface {
	EsSolucion(camino.Camino) bool
}
