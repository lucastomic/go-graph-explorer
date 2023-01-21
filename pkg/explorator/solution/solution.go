package solution

// Method that evaluates if a state is the solution of the problem
type Solution interface {
	IsSolution(int, [][]float64) bool
}
