package sliceUtils

// Eliminar último elemento del slice
func RemoveLast[T any](sPointer *[]T) {
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
	return len(list) == 1
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
