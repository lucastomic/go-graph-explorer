package prune

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/lucastomic/ExploracionDeEspacios/pkg/explorador/path"
)

// Differents paths
var path1 path.Path = path.NewPath(&[]int{0, 1, 4, 3})
var path2 path.Path = path.NewPath(&[]int{0, 1, 3})
var path3 path.Path = path.NewPath(&[]int{0, 2})
var path4 path.Path = path.NewPath(&[]int{0, 5})
var path5 path.Path = path.NewPath(&[]int{0, 5})

// Graph used
var maxF float64 = math.MaxFloat64
var graph [][]float64 = [][]float64{
	{maxF, 0.5, 0.5, maxF, maxF, 0.5},
	{0.5, maxF, maxF, 2.5, 5, maxF},
	{0.5, maxF, maxF, maxF, maxF, maxF},
	{maxF, 2.5, maxF, maxF, 2, maxF},
	{maxF, 5, maxF, 2, maxF, maxF},
	{0.5, maxF, maxF, maxF, maxF, maxF},
}

// Differents sets of tests.
var pruneTests = []struct {
	paths    []path.Path
	expected []path.Path
}{
	{
		[]path.Path{path1, path2, path3, path4},
		[]path.Path{path2, path3, path4},
	},
	{
		[]path.Path{path2, path3, path4},
		[]path.Path{path2, path3, path4},
	},
	{
		[]path.Path{path2, path3, path4, path5},
		[]path.Path{path2, path3, path5},
	},
}

func TestPodar(t *testing.T) {
	for i, tt := range pruneTests {
		testName := fmt.Sprintf("Prune test number %v", i)
		t.Run(testName, func(t *testing.T) {
			Prune(&tt.paths, graph)
			if !reflect.DeepEqual(tt.paths, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, tt.paths)
			}
		})
	}
}
