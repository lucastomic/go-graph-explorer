package heuristico

// Heuristic state evaluates the states and returns a verifiable value of how promising they are
// in order to reach the solution
type StateHeuristic interface {
	Heuristic(int) float64
}
