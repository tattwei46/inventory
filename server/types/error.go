package types

import "github.com/pkg/errors"

var (
	BadRequest          = errors.New("BAD_REQUEST")
	InternalServerError = errors.New("INTERNAL_SERVER_ERROR")
)
