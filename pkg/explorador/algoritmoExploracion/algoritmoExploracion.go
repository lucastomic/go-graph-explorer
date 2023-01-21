package algoritmoexploracion

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"

// Interfaz que cubre los distintos aspectos dependientes del sistema de un problema de exploracion.
// Esto es defenido por:
//	La forma en que ordena los caminos pendientes por explorar
//	Como define cuando un estado e solucion

type AlgoritmoExploracion interface {
	Mezclar(viejos *[]path.Path, nuevos *[]path.Path)
}

// Estos son los diferentes tipos de algoritmo que hay:
type TipoAlgoritmoExp int

const (
	AlBranchAndBonud = iota
	AlAEstrella
	AlEscalada
	AlPrimeroElMejor
	AlProfundidad
	AlAmplitud
)
