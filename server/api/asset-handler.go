package api

import (
	"net/http"
	"strconv"

	"github.com/tattwei46/inventory/server/types"

	"github.com/tattwei46/inventory/server/param"

	"github.com/tattwei46/inventory/server/service"

	"github.com/tattwei46/inventory/server/framework/logger"

	"github.com/gin-gonic/gin"
)

type assetHandler struct {
	service.Asset
	log *logger.Logger
}

func newAssetHandler() (*assetHandler, error) {
	service, err := service.NewAsset()
	if err != nil {
		return nil, err
	}
	return &assetHandler{service, logger.GetInstance()}, nil
}

func (h *assetHandler) get(c *gin.Context) {
	var limit int
	var page int

	if val, err := strconv.Atoi(c.Param("limit")); err == nil && val > 0 {
		limit = val
	}
	if val, err := strconv.Atoi(c.Param("page")); err == nil && val > 0 {
		page = val
	}

	var search param.Search

	if err := c.BindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, types.Response.NewError(types.BadRequest))
		return
	}

	res, err := h.Asset.Get(search, limit, page)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, types.Response.NewError(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *assetHandler) add(c *gin.Context) {
	var request param.Asset

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, types.Response.NewError(types.BadRequest))
		return
	}

	if err := h.Asset.Add(request); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, types.Response.NewError(err))
		return
	}

	c.Status(http.StatusOK)
}
