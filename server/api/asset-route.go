package api

import (
	"github.com/tattwei46/inventory/server/framework/logger"

	"github.com/gin-gonic/gin"
)

func Asset() asset {
	return asset{url: "assets", log: logger.GetInstance()}
}

type asset struct {
	url string
	log *logger.Logger
}

func (h asset) Routes(router *gin.Engine) {
	handler, err := newAssetHandler()
	if err != nil {
		h.log.Error("an error occurred when creating asset handler")
	}
	r := route(router, h.url)

	r.GET("/", handler.get)
	r.POST("/", handler.add)
}
