package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func ParseID(str string) (ID, error) {
	id, err := uuid.Parse(str)

	return ID(id), err
}
