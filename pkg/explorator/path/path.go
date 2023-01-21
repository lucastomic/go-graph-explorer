package path

import (
	"math"

	"github.com/lucastomic/ExploracionDeEspacios/internals/sliceUtils"
)

// A path is an acyclic succession of states. These states are represented
// for an Int that matches the index that represents them in the adjacency matrix of the graph.
type Path struct {
	states *[]int
}

// Returns a path with the states specified by argument
func NewPath(estados *[]int) Path {
	return Path{states: estados}
}

// Returns a empty path
func NewEmptyPath() Path {
	states := &[]int{}
	return Path{states: states}
}

// Returns the total cost of the path
//For example, given the next graph:

// (0)───0.5──(1)───5────┐
//  │          │        (4)
// 0.5        2.5
//  │          │
// (2)        (3)

// And the path [2,0,1,4]
// The total cost wold be 0.5+0.5+5 = 6

// To do this iterates the costs of all transitions.
// Transitions are changes from one state to another.
func (c Path) GetTotalCost(graph [][]float64) float64 {
	var totalCost float64
	//We iterate transitions of states. Until len(path)-1 because the last state does not transition to any other state.
	for i := 0; i < len(*c.states)-1; i++ {
		currentState := (*c.states)[i]
		nexState := (*c.states)[i+1]
		totalCost += graph[currentState][nexState]
	}

	return totalCost
}

// getPathWithNewState Returns a path with the same states but with the state passed as an argument at the end.
// It doesn't modfy the actually states.
// For exmale, having a path with the next states:
// [0,4,3,9,6]
// getPathWithNewState(5) would return a different path with the next states:
// [0,4,3,9,6,5]
func (c Path) getPathWithNewState(newState int) Path {
	states := append(*c.states, newState)
	return NewPath(&states)
}

// Convert partial path into a set of new paths, the result of adding each one to it
// of the possible successors that do not give rise to cyclic paths
//
// For example, given the next graph:

// (5)───0.5──(0)──0.5──(1)───5────┐
//             │         │         │
//             │         │        (4)
//            0.5       2.5        │
//             │         │         2
//             │         │         │
//            (2)       (3)────────┘

// and the nex path: [5,0,1]
// the result of expanding this path wold be:
// [ [5,0,1,4], [5,0,1,3] ]
func (p Path) Expand(graph [][]float64) *[]Path {
	var res []Path
	state := *p.states
	currentState := state[len(state)-1]
	adjacentOfCurrentState := graph[currentState]

	// For each state adjacent to the current one (the last one in the list) that is not in the path,
	// add a path to the response with path+state adjacent
	for i := range adjacentOfCurrentState {
		if p.notInfinit(adjacentOfCurrentState[i]) && !sliceUtils.Contains(*p.states, i) {
			res = append(res, p.getPathWithNewState(i))
		}
	}
	return &res
}

// If [val] is different from infinity (This in the graph means that they are connected)
func (c Path) notInfinit(val float64) bool {
	return val != math.MaxFloat64
}

// Return the last state of the path (the current state)
func (c Path) GetCurrentState() int {
	states := *c.states
	return states[len(states)-1]
}