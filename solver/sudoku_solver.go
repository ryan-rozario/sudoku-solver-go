package solver

import "github.com/ryan-rozario/sudoku-solver/sudoku"

func findEmpty(board *sudoku.SudokuBoard) (int, int) {

	for h := 0; h < board.Height; h++ {
		for w := 0; w < board.Width; w++ {
			if board.Grid[h][w] == 0 {
				return h, w
			}
		}
	}

	return -1, -1

}
