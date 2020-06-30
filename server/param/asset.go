package param

import "github.com/tattwei46/inventory/server/types"

type Asset struct {
	ID           string `json:"id" `
	Created      string `json:"created"`
	SerialNumber string `json:"serial_number" binding:"required"`
	Brand        string `json:"brand" binding:"required"`
	Model        string `json:"model" binding:"required"`
	Status       string `json:"status"`
}

type Search struct {
	ID           string      `json:"id" `
	Range        types.Range `json:"range"`
	SerialNumber string      `json:"serial_number"`
	Brand        string      `json:"brand"`
	Model        string      `json:"model"`
	Status       string      `json:"status"`
}
