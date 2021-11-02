package uuid

import "github.com/google/uuid"

type Generator interface {
	New() string
}

type uuidGen struct{}

func New() Generator {
	return &uuidGen{}
}

func (u uuidGen) New() string {
	return uuid.New().String()
}
