package solucion

// Método que evalua si un estado es la solución del problema
type Solucion interface {
	EsSolucion([]int) bool
}
