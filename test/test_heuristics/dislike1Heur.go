package testheuristics

// Heuristic which dislikes 1
type Dislike1Heur struct{}

// Rerturns 10 if the state == 1, returns 1 otherwise
func (h Dislike1Heur) Heuristic(state int) float64 {
	if state == 1 {
		return 10
	} else {
		return 1
	}
}
func NewDislike1Heur() Dislike1Heur {
	return Dislike1Heur{}
}
