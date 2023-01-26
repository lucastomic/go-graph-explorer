package solutiontests

// This solution is true when the sate is three
type ThreeSolution struct {
}

// Returns true if the state is 3
func (s ThreeSolution) IsSolution(state int, graph [][]float64) bool {
	return state == 3
}

func NewThreeSolution() ThreeSolution {
	return ThreeSolution{}
}
