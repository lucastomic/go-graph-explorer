package floatutils

import "math"

// If [val] is different from infinity (This in the graph means that they are connected)
func NotInfinit(val float64) bool {
	return val != math.MaxFloat64
}
