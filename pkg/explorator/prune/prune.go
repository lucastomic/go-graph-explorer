package prune

import "github.com/lucastomic/go-graph-explorer/pkg/explorator/path"

// Whenever two partial paths lead to the same state, the most expensive of them is removed. Since
// whatever state can be reached with the most expensive can also be reached with the other
//
// For example:
// [
//
//	[0,1,2,3],
//	[0,1,3],
//	[0,1,4],
//	[0,1,2,5],
//
// ]
//
// Will turn into:
//
// [
//
//	[0,1,3],
//	[0,1,4],
//	[0,1,2,5],
//
// ]
//
// becuase cost([0,1,2,3]) < cost([0,1,3])
func Prune(paths *[]path.Path, graph [][]float64) {
	// Hash table that indicates the path of each final state.
	// This means the following structure [endState:pathIndex]
	//
	// An example would be:
	// For the following paths: [[0,4,3], [0,4,2],[0,3,4]]
	// The following hash table:
	// {
	//   3:0, (final state 3 already exists in a path and is the one with index 0)
	// 	 2:1,
	// 	 4:2,
	// }

	hash := make(map[int]int)
	var removeIndexes []int

	for i := range *paths {
		endState := (*paths)[i].GetCurrentState()
		//pathAlreadyExists indicates if a path with that final state has already been evaluated
		//oldPathIndex indicates the index of the path that has that state as its final state
		oldPathIndex, pathAlreadyExists := hash[endState]
		if pathAlreadyExists {
			oldPathCost := (*paths)[oldPathIndex].GetTotalCost(graph)
			newPathCost := (*paths)[i].GetTotalCost(graph)
			if oldPathCost >= newPathCost {
				hash[endState] = i
				removeIndexes = append(removeIndexes, oldPathIndex)
			}
		} else {
			hash[endState] = i
		}
	}
	removeSeveralIndex(paths, removeIndexes)
}

// removeSeveralIndex removes the elements from the first slice whose index are passed on the second one.
// For example, given this slice:
//
// [apple, banana, pencil, tree]
//
// and this slice:
//
// [1,2]
//
// it will remove the "banana" and "pencil" elements
//
// [apple, tree]
func removeSeveralIndex(paths *[]path.Path, indexes []int) {
	for _, index := range indexes {
		*paths = append((*paths)[:index], (*paths)[index+1:]...)
	}
}
