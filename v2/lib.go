package workspace_lib

import (
	"golang.org/x/exp/constraints"
)

// Number is a type constraint for integers and floats.
type Number interface {
	constraints.Integer | constraints.Float
}

// AddNums adds two numbers of a type that satisfies the Number constraint.
func AddNums[T Number](a, b T) T {
	return a + b
}

func SubNums[T Number](a, b T) T {
	return a - b
}
