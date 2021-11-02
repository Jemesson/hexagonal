package ports

import "hexagonal/src/core/domain"

type GameRepositoryPort interface {
	Get(id string) (domain.Game, error)
	Save(game domain.Game) error
}
