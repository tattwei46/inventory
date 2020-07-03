package types

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Response ResponseType

type ResponseType struct{}

func (ResponseType) NewError(err error) interface{} {
	return gin.H{"error": fmt.Sprintf("%s", err.Error())}
}

func (ResponseType) NewSuccess(message string) interface{} {
	return gin.H{"message": message}
}
