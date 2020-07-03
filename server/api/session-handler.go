package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
}

func newSessionHandler() *sessionHandler {
	return &sessionHandler{}
}

func (h *sessionHandler) new(c *gin.Context) {
	c.Status(http.StatusOK)
}
