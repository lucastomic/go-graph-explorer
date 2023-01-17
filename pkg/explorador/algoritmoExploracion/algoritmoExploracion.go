package algoritmoexploracion

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/camino"

// Interfaz que cubre los distintos aspectos dependientes del sistema de un problema de exploracion.
// Esto es defenido por:
//	La forma en que ordena los caminos pendientes por explorar
//	Como define cuando un estado e solucion

type AlgoritmoExploracion interface {
	Mezclar(viejos *[]camino.Camino, nuevos *[]camino.Camino)
}
