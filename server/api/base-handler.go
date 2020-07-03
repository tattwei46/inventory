package api

import (
	"fmt"
	"net/http"

	"github.com/tattwei46/inventory/server/framework/logger"

	"github.com/gin-gonic/gin"
)

var (
	BuildVersion = ""
	BuildTime    = ""
)

type baseHandler struct {
	log *logger.Logger
}

func newBaseHandler() (*baseHandler, error) {
	return &baseHandler{log: logger.GetInstance()}, nil
}

func (h *baseHandler) health(c *gin.Context) {
	h.log.Debug("Health endpoint is called")
	c.JSON(http.StatusOK, gin.H{"message": "ping successful"})
}

func (h *baseHandler) version(c *gin.Context) {
	version := fmt.Sprintf("%s-%s", BuildVersion, BuildTime)
	c.JSON(http.StatusOK, version)
}
