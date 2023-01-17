package poda

import "github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/camino"

// Siempre que dos caminos parciales conducen a un mismo
// estado, se elimina el más costoso de ellos, puesto que
// cualquier estado al que pueda llegarse con el más costoso
// también puede alcanzarse con el otro
func Podar(caminos []camino.Camino, grafo [][]float64) []camino.Camino {
	//Tabla hash que indica camino de cada estado final.
	//Esto significa la siguiente estructura [estadoFinal:indiceCamino]
	//Un ejemplo sería:
	//Para el siguiente grafo: [[0,4,3], [0,4,2],[0,3,4]]
	//La siguiente tabla hash:
	//{
	//  3:0, (el estado final 3 ya existe en un camino y es el que tiene indice 0)
	//  2:1,
	//	4:2
	//}
	hash := make(map[int]int)
	for i := range caminos {
		//Estado final del camino actual
		estadoFinal := caminos[i].GetCurrentState()
		//yaExisteCamino indica si ya se evaluó un camino con ese estado final
		//indiceCaminoAntiguo indica el indice del camino que tiene ese estado como estado final
		indiceCaminoAntiguo, yaExisteCamino := hash[estadoFinal]
		if yaExisteCamino {
			costeCaminoAntiguo := caminos[indiceCaminoAntiguo].GetTotalCost(grafo)
			costeCaminoNuevo := caminos[i].GetTotalCost(grafo)
			if costeCaminoAntiguo > costeCaminoNuevo {
				hash[estadoFinal] = i
			}
		} else {
			hash[estadoFinal] = i
		}
	}
	return getCaminosFromHash(caminos, hash)
}

// Devuelve una matriz con los caminos cuyos indices están en los valores de la tabla hash pasada.
// Esto quiere decir:
// Si la tabla hash tiene el siguiente cuerpo:
//
//	{
//	  clave1:0,
//	  clave2:3,
//	  clave3:5
//	}
//
// La matriz devuelta será:
// [caminos[0], caminos[3], caminos[5]]
func getCaminosFromHash(caminos []camino.Camino, hash map[int]int) []camino.Camino {
	var caminosPodados []camino.Camino
	for _, indiceCamino := range hash {
		caminosPodados = append(caminosPodados, caminos[indiceCamino])
	}
	return caminosPodados
}
