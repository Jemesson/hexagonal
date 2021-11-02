package memory_kvs

import (
	"encoding/json"
	"github.com/matiasvarela/errors"
	"hexagonal/src/config/apperrors"
	"hexagonal/src/config/messages"
	"hexagonal/src/core/domain"
)

type MemoryKVS struct {
	kvs map[string][]byte
}

func NewMemKVS() *MemoryKVS {
	return &MemoryKVS{
		kvs: map[string][]byte{},
	}
}

func (repo *MemoryKVS) Get(id string) (domain.Game, error) {

	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return domain.Game{}, errors.New(apperrors.Internal, err, messages.GameNotFoundFromKVS)
		}

		return game, nil
	}

	return domain.Game{}, errors.New(apperrors.NotFound, nil, "game not found in kvs")
}

func (repo *MemoryKVS) Save(game domain.Game) error {
	bytes, err := json.Marshal(game)
	if err != nil {
		return errors.New(apperrors.InvalidInput, err, messages.GameMarshalingFailed)
	}

	repo.kvs[game.ID] = bytes
	return nil
}
