package sliceUtils

// Eliminar último elemento del slice
func EliminarUltimo(sPointer *[][]int) {
	slice := *sPointer
	slice = slice[:len(slice)-1]
}

// Whether [list] contains [val]
func Contains(list []int, val int) bool {
	for _, el := range list {
		if el == val {
			return true
		}
	}
	return false
}

// Whether [list] is not empty
func IsEmpty(list [][]int) bool {
	return len(list) == 0
}
