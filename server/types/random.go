package types

import "github.com/google/uuid"

func GetRandom() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
