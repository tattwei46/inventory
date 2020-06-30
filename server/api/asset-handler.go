package api

import (
	"fmt"
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

func (h *assetHandler) search(c *gin.Context) {
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
	var requests []param.Asset

	if err := c.BindJSON(&requests); err != nil {
		c.JSON(http.StatusBadRequest, types.Response.NewError(types.BadRequest))
		return
	}

	err := h.Asset.Add(requests)

	if err == types.DuplicatedItem {
		c.JSON(http.StatusBadRequest, types.Response.NewError(err))
		return
	}

	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, types.Response.NewError(err))
		return
	}

	c.Status(http.StatusOK)
}

func (h *assetHandler) get(c *gin.Context) {
	res, err := h.Asset.Get(param.Search{}, 0, 0)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, types.Response.NewError(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *assetHandler) delete(c *gin.Context) {
	id := c.Param("id")
	count, err := h.Asset.Delete(id)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, types.Response.NewError(err))
		return
	}
	c.JSON(http.StatusOK, types.Response.NewSuccess(fmt.Sprintf("Deleted %d", count)))
}

func (h *assetHandler) update(c *gin.Context) {
	id := c.Param("id")

	var update param.Asset
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, types.Response.NewError(types.BadRequest))
		return
	}

	err := h.Asset.Update(id, update)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, types.Response.NewError(err))
		return
	}
	c.JSON(http.StatusOK, types.Response.NewSuccess("updated"))
}
