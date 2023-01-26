package path

import (
	"reflect"
	"strconv"

	floatutils "github.com/lucastomic/go-graph-explorer/pkg/utils/floatUtils"
	"github.com/lucastomic/go-graph-explorer/pkg/utils/sliceUtils"
)

// A path is an acyclic succession of states. These states are represented
// for an Int that matches the index that represents them in the adjacency matrix of the graph.
type Path struct {
	states *[]int
}

// Returns a path with the states specified by argument
func NewPath(states *[]int) Path {
	return Path{states: states}
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
// It copies the path's state in another slice to don't modify the path's state with the append() method.
func (c Path) getPathWithNewState(newState int) Path {
	var states []int = make([]int, len(*c.states))
	copy(states, *c.states)
	states = append(states, newState)
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
	var res []Path = make([]Path, 0)
	adjacentOfCurrentState := graph[p.GetCurrentState()]

	// For each state adjacent to the current one (the last one in the list) that is not in the path,
	// add a path to the response with path+state adjacent
	for i := range adjacentOfCurrentState {
		if p.isValid(adjacentOfCurrentState, i) {
			res = append(res, p.getPathWithNewState(i))
		}
	}
	return &res
}

// isValid cheks if the state passed as an argument is connected and is not in path yet, according
// to the vector of adjacent states
func (p Path) isValid(adjacentOfCurrentState []float64, state int) bool {
	isConnected := floatutils.NotInfinit(adjacentOfCurrentState[state])
	isInPath := sliceUtils.Contains(*p.states, state)
	return isConnected && !isInPath
}

// Return the last state of the path (the current state)
func (c Path) GetCurrentState() int {
	states := *c.states
	return states[len(states)-1]
}

// Converts a path to string. It shows all his states connected by an arrow.
// For example, a path with the next states : [0,4,2,3]
// Is returned like this "0 -> 4 -> 2 -> 3"
func (c Path) ToString() string {
	var s string
	for i := 0; i < len(*c.states)-1; i++ {
		s += strconv.Itoa((*c.states)[i])
		s += " -> "
	}
	s += strconv.Itoa(c.GetCurrentState())
	return s
}

// Compare if this paht is equal to another slice. It doesn't compare
// if it's the same object, it compares if it has the same paht.
func (p Path) Equal(otherPath Path) bool {
	otherStates := otherPath.states
	if !reflect.DeepEqual(otherStates, p.states) {
		return false
	}
	return true
}

// ComparePathsSlices check whether two paths slices have the sames paths.
// This means that every path have the same states
func ComparePathsSlices(pathA, pathB []Path) bool {
	if len(pathA) != len(pathB) {
		return false
	}
	for i := range pathA {
		aStates := pathA[i].states
		bStates := pathB[i].states
		if !reflect.DeepEqual(*aStates, *bStates) {
			return false
		}
	}
	return true
}
