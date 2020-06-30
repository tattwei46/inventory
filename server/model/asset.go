package model

import "github.com/tattwei46/inventory/server/types"

type Asset struct {
	ID           string       `bson:"id"`
	SerialNumber string       `bson:"serial_number"`
	Model        string       `bson:"model"`
	Status       types.Status `bson:"status"`
	Created      int64        `bson:"created"`
}
