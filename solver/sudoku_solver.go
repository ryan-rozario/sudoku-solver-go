package solver

import (
	"github.com/ryan-rozario/sudoku-solver/sudoku"
)

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

func checkValid(board *sudoku.SudokuBoard, row int, col int, val int) bool {
	for h := 0; h < board.Height; h++ {
		if board.Grid[h][col] == val {
			if val == 6 {
			}
			return false
		}
	}

	for w := 0; w < board.Width; w++ {
		if board.Grid[row][w] == val {
			if val == 6 {
			}
			return false
		}
	}

	var box_h int
	var box_w int
	box_h = (row / 3) * 3
	box_w = (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			if board.Grid[box_h+i][box_w+j] == val {

				return false
			}
		}

	}

	return true

}

func FillGrid(board *sudoku.SudokuBoard) bool {
	row, col := findEmpty(board)
	if row == -1 && col == -1 {
		return true
	}

	for cell_val := 1; cell_val <= 9; cell_val++ {
		if checkValid(board, row, col, cell_val) {
			board.Grid[row][col] = cell_val
			if FillGrid(board) {
				return true
			}
		}
		board.Grid[row][col] = 0
	}

	return false

}
