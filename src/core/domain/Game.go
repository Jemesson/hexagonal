package domain

const (
	GameStateWon  = "won"
	GameStateLost = "lost"
	GameStateNew  = "new"
)

type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"board_settings"`
	Board         Board         `json:"board"`
}

func NewGame(id string, name string, size uint, bombs uint) Game {
	return Game{
		ID:    id,
		Name:  name,
		State: GameStateNew,
		BoardSettings: BoardSettings{
			Size:  size,
			Bombs: bombs,
		},
		Board: NewBoard(size, bombs),
	}
}

func (game *Game) IsOver() bool {
	return game.State == GameStateLost || game.State == GameStateWon
}
