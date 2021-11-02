package messages

const (
	GameNotFound                      = "game not found"
	GameFailedFromRepository          = "get game from repository has failed"
	GameCannotBeCreatedFromRepository = "create game into repository has failed"
	GameCannotBeUpdateFromRepository  = "update game into repository has failed"
	GameBombsTooHigh                  = "the number of bombs is too high"
	GameOver                          = "game is over"
	GameInvalidPosition               = "invalid position"
	GameNotFoundFromKVS               = "fail to get value from kvs"
	GameMarshalingFailed              = "game fails at marshal into json string"
)
