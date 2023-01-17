package sliceUtils

// Eliminar último elemento del slice
func EliminarUltimo[T any](sPointer *[]T) {
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
func IsEmpty[T any](list []T) bool {
	return len(list) == 0
}
