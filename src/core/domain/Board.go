package domain

import (
	"math/rand"
	"time"
)

const (
	CellBomb       = "X"
	CellBombHidden = "-"
	CellEmpty      = "-"
	CellRevealed   = "0"
)

type Board [][]string

func NewBoard(size uint, bombs uint) Board {
	board := NewEmptyBoard(size)
	board.fillWithBombs(bombs)

	return board
}

func NewEmptyBoard(size uint) Board {
	board := make([][]string, size)

	for row := range board {
		board[row] = make([]string, size)
	}

	for row := range board {
		for col := range board[0] {
			board[row][col] = CellEmpty
		}
	}

	return board
}

func (board Board) fillWithBombs(bombs uint) {

	rows := len(board)
	cols := len(board[0])
	positions := _getRandomPositions(rows*cols, bombs)

	var row, col int
	for _, pos := range positions {
		row = pos / cols
		col = pos - row*rows
		board[row][col] = CellBomb
	}
}

func (board Board) HideBombs() Board {
	newBoard := NewEmptyBoard(uint(len(board)))

	for row := range board {
		for col := range board[0] {
			if board[row][col] == CellBomb {
				newBoard[row][col] = CellBombHidden
			} else {
				newBoard[row][col] = board[row][col]
			}
		}
	}

	return newBoard
}

func (board Board) IsValidPosition(row uint, col uint) bool {
	return row < uint(len(board)) && col < uint(len(board[0]))
}

func (board Board) Contains(row uint, col uint, element string) bool {
	return board[row][col] == element
}

func (board Board) Set(row uint, col uint, element string) {
	board[row][col] = element
}

func (board Board) IsCellEmpty() bool {
	for row := range board {
		for col := range board[0] {
			if board[row][col] == CellEmpty {
				return true
			}
		}
	}

	return false
}

// ··· Private Functions ··· //
func _getRandomPositions(size int, n uint) []int {
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(size)

	var positions []int

	for _, r := range p[:n] {
		positions = append(positions, r)
	}

	return positions
}
