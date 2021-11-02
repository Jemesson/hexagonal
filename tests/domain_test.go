package tests

import (
	"github.com/stretchr/testify/assert"
	"hexagonal/src/core/domain"
	"testing"
)

func TestNewBoard(t *testing.T) {
	size := uint(10)
	bombs := uint(10)

	board := domain.NewBoard(size, bombs)

	countBombs := uint(0)
	for i := range board {
		for j := range board[0] {
			if board[i][j] == domain.CellBomb {
				countBombs++
			}
		}
	}

	assert.Equal(t, uint(len(board)), size)
	assert.Equal(t, uint(len(board[0])), size)
	assert.Equal(t, countBombs, bombs)
}

func TestBoardHiddenBombs(t *testing.T) {
	board := domain.NewBoard(10, 50).HideBombs()
	isBombHidden := true

	for i := range board {
		if !isBombHidden {
			break
		}

		for j := range board[0] {
			if board[i][j] != domain.CellEmpty {
				isBombHidden = false
				break
			}
		}
	}

	assert.True(t, isBombHidden)
}

func TestBoardIsPositionValid(t *testing.T) {
	board := domain.NewBoard(10, 50)

	assert.True(t, board.IsValidPosition(5, 5))
	assert.False(t, board.IsValidPosition(11, 5))
	assert.False(t, board.IsValidPosition(5, 11))
}

func TestBoardContainsBombs(t *testing.T) {
	board := domain.NewBoard(10, 50)
	board[1][2] = domain.CellBomb

	assert.True(t, board.Contains(1, 2, domain.CellBomb))
}

func TestBoardIsCellEmpty(t *testing.T) {
	boardWithEmptyCells := domain.NewBoard(10, 50)
	boardWithoutEmptyCells := domain.NewBoard(10, 100)

	assert.True(t, boardWithEmptyCells.IsCellEmpty())
	assert.False(t, boardWithoutEmptyCells.IsCellEmpty())
}

// ··· GAME TESTS ··· //

func TestNewGame(t *testing.T) {
	game := domain.NewGame("1001", "new game", 10, 50)

	assert.Equal(t, "1001", game.ID)
	assert.Equal(t, "new game", game.Name)
	assert.Equal(t, domain.BoardSettings{Size: 10, Bombs: 50}, game.BoardSettings)
	assert.Equal(t, domain.GameStateNew, game.State)
	assert.Equal(t, 10, len(game.Board))
	assert.Equal(t, 10, len(game.Board[0]))
}

func TestGame_IsOver(t *testing.T) {
	gameNew := domain.NewGame("1001", "new game", 10, 50)

	gameWon := domain.NewGame("1001", "won game", 10, 50)
	gameWon.State = domain.GameStateWon

	gameLost := domain.NewGame("1001", "lost game", 10, 50)
	gameLost.State = domain.GameStateLost

	assert.False(t, gameNew.IsOver())
	assert.True(t, gameWon.IsOver())
	assert.True(t, gameLost.IsOver())
}
