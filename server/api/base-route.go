package api

import (
	"fmt"

	"github.com/tattwei46/inventory/framework"

	"github.com/gin-gonic/gin"
)

const contextRoot = "interface/v1"

func route(router *gin.Engine, url string) *gin.RouterGroup {
	return router.Group(fmt.Sprintf("%s/%v", contextRoot, url))
}

func Base() base {
	return base{url: "", log: framework.GetLoggerInstance()}
}

type base struct {
	url string
	log *framework.Logger
}

func (h base) Routes(router *gin.Engine) {
	handler, err := newBaseHandler()
	if err != nil {
		h.log.Error("An Error Occurred while Creating  Base Handler")
	}
	r := route(router, h.url)

	r.GET("/health", handler.health)
	r.GET("/version", handler.version)
}
