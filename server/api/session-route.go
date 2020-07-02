package api

import (
	"github.com/tattwei46/inventory/server/framework/logger"

	"github.com/gin-gonic/gin"
)

func Session() session {
	return session{url: "sessions", log: logger.GetInstance()}
}

type session struct {
	url string
	log *logger.Logger
}

func (h session) Routes(router *gin.Engine) {
	handler := newSessionHandler()
	r := route(router, h.url)

	r.POST("/", handler.new)
}
