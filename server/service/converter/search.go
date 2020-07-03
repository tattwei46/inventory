package converter

import (
	"github.com/tattwei46/inventory/server/model"
	"github.com/tattwei46/inventory/server/param"
	"github.com/tattwei46/inventory/server/types"
)

type search struct{}

var Search search

func (search) ToModel(request param.Search) model.Search {
	return model.Search{
		ID:           request.ID,
		Range:        request.Range.ToUnix(),
		Brand:        request.Brand,
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Status:       types.GetStatus(request.Status),
	}
}
