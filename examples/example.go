package examples

import (
	"fmt"
	"math"

	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorationAlgorithm/enums/informedAlgorithm"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorationAlgorithm/enums/uninformedAlgorithm"
	"github.com/lucastomic/go-graph-explorer/pkg/explorator/explorer"
)

// (5)───0.5──(0)────0.5──(1)───5────┐
//             │           │         │
//             │           │        (4)
//            0.5         2.5        │
//             │           │         │
//             │           │         2
//             │           │         │
//            (2)         (3)────────┘

var maxF float64 = math.MaxFloat64
var testGraph [][]float64 = [][]float64{
	{maxF, 0.5, 0.5, maxF, maxF, 0.5},
	{0.5, maxF, maxF, 2.5, 5, maxF},
	{0.5, maxF, maxF, maxF, maxF, maxF},
	{maxF, 2.5, maxF, maxF, 2, maxF},
	{maxF, 5, maxF, 2, maxF, maxF},
	{0.5, maxF, maxF, maxF, maxF, maxF},
}

// Heuristic which dislikes 3
type Dislike1Heur struct{}

// Rerturns 10 if the state == 1, returns 1 otherwise
func (h Dislike1Heur) Heuristic(state int) float64 {
	if state == 1 {
		return 10
	} else {
		return 1
	}
}

// This solution is true when the sate is three
type ThreeSolution struct {
}

// Returns true if the state is 3
func (s ThreeSolution) IsSolution(state int, graph [][]float64) bool {
	return state == 3
}

func main() {
	res, _ := explorer.ExploreWithInformed(testGraph, ThreeSolution{}, Dislike1Heur{}, 5, informedAlgorithm.AStar)
	// explorer.ExploreWithInformed(testGraph, ThreeSolution{}, Dislike1Heur{}, 5, informedAlgorithm.Climbing)
	// explorer.ExploreWithInformed(testGraph, ThreeSolution{}, Dislike1Heur{}, 5, informedAlgorithm.BestFirst)
	fmt.Println(res.ToString())

	res2, _ := explorer.ExploreWithUninformed(testGraph, ThreeSolution{}, 5, uninformedAlgorithm.Amplitude)
	// explorer.ExploreWithUninformed(testGraph,ThreeSolution{},5,uninformedAlgorithm.DepthFirst)
	// explorer.ExploreWithUninformed(testGraph,ThreeSolution{},5,uninformedAlgorithm.BranchAndBonud)
	fmt.Println(res2.ToString())
}
