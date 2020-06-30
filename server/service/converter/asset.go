package converter

import (
	"time"

	"github.com/tattwei46/inventory/server/model"
	"github.com/tattwei46/inventory/server/param"
	"github.com/tattwei46/inventory/server/types"
)

type asset struct{}

var Asset asset

func (asset) ToModel(request param.Asset) model.Asset {
	return model.Asset{
		ID:           request.ID,
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Status:       types.Available,
		Created:      time.Now().Unix(),
	}
}

func (asset) ToParam(request *model.Asset) param.Asset {
	return param.Asset{
		ID:           request.ID,
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Status:       request.Status.String(),
		Created:      types.Format(request.Created, types.TimeZone, types.YYYYMMDD_hhmmssMST),
	}
}

func (asset) ToParams(requests []model.Asset) []param.Asset {
	var result = make([]param.Asset, 0)

	for _, m := range requests {
		result = append(result, Asset.ToParam(&m))
	}

	return result
}
