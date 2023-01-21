package test

import "math"

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
