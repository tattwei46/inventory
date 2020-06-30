package api

import (
	"path/filepath"

	"github.com/tattwei46/inventory/framework/logger"

	"github.com/gin-gonic/gin"
)

const contextRoot = "inventory/v1"

func route(router *gin.Engine, url string) *gin.RouterGroup {
	return router.Group(filepath.Join(contextRoot, url))
}

func Base() base {
	return base{url: "", log: logger.GetLoggerInstance()}
}

type base struct {
	url string
	log *logger.Logger
}

func (h base) Routes(router *gin.Engine) {
	handler, err := newBaseHandler()
	if err != nil {
		h.log.Error("An Error Occurred while Creating Base Handler")
	}
	r := route(router, h.url)

	r.GET("/health", handler.health)
	r.GET("/version", handler.version)
}
