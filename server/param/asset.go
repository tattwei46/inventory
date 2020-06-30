package param

type Asset struct {
	ID           string `json:"id" `
	Created      string `json:"created"`
	SerialNumber string `json:"serial_number" binding:"required"`
	Brand        string `json:"brand" binding:"required"`
	Model        string `json:"model" binding:"required"`
	Status       string `json:status`
}
