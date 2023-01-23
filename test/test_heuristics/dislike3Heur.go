package testheuristics

// Heuristic which dislikes 3
type Dislike3Heur struct{}

// Rerturns 10 if the state == 3, returns 1 otherwise
func (h Dislike3Heur) Heuristic(state int) float64 {
	if state == 3 {
		return 10
	} else {
		return 1
	}
}
func NewDislike3Heur() Dislike3Heur {
	return Dislike3Heur{}
}
