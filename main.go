package main

import (
	"fmt"

	"github.com/ryan-rozario/sudoku-solver/solver"
	"github.com/ryan-rozario/sudoku-solver/sudoku"
)

func main() {
	test_mat := [][]int{
		{0, 2, 0, 5, 0, 1, 0, 9, 0},
		{8, 0, 0, 2, 0, 3, 0, 0, 6},
		{0, 3, 0, 0, 6, 0, 0, 7, 0},
		{0, 0, 1, 0, 0, 0, 6, 0, 0},
		{5, 4, 0, 0, 0, 0, 0, 1, 9},
		{0, 0, 2, 0, 0, 0, 7, 0, 0},
		{0, 9, 0, 0, 3, 0, 0, 8, 0},
		{2, 0, 0, 8, 0, 4, 0, 0, 7},
		{0, 1, 0, 9, 0, 7, 0, 6, 0},
	}
	nb := sudoku.NewBoard(test_mat)
	solver.FillGrid(nb)
	fmt.Println(nb.Grid)

}
