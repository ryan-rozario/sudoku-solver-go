package sudoku

type SudokuBoard struct {
	Grid   [][]int
	Height int
	Width  int
}

func NewBoard(board [][]int) *SudokuBoard {
	b := SudokuBoard{Grid: board, Height: len(board), Width: len(board[0])}
	return &b
}
