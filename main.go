package main

import (
	"fmt"

	"github.com/antonio-muniz/sudoku-solver/sudoku"
)

func main() {
	set := sudoku.NewValueSet(1, 2, 3, 4, 5)
	powerSet := set.PowerSet()
	fmt.Println(powerSet)
}
