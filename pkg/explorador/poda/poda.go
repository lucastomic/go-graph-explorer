package poda

type Poda struct {
}

// Siempre que dos caminos parciales conducen a un mismo
// estado, se elimina el más costoso de ellos, puesto que
// cualquier estado al que pueda llegarse con el más costoso
// también puede alcanzarse con el otro
func (p Poda) podar(caminos *[][]int, grafo [][]int) {

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
	
	for i:= range *caminos{
	  estadoFinal := caminos[i][len(caminos[i])]
	  if camino,hayColision := hash[estadoFinal], yaExiste{
		 
	  }else{
		hash[estadoFinal] = i
	  }
	}
}

//Devuelve el coste total que tiene un camino
func (p Poda) costeCamino(camino []int , grafo[][]int) float64{
  var costeTotal int
  for  estado := range camino {
		
  }
  
}
