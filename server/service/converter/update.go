package converter

import (
	"github.com/tattwei46/inventory/server/model"
	"github.com/tattwei46/inventory/server/param"
	"github.com/tattwei46/inventory/server/types"
)

type update struct{}

var Update update

func (update) ToModel(request *param.Asset) model.Asset {

	return model.Asset{
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Brand:        request.Brand,
		Status:       types.GetStatus(request.Status),
		Created:      types.ToUnix(request.Created),
	}
}
