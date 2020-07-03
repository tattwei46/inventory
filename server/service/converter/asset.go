package converter

import (
	"time"

	"github.com/tattwei46/inventory/server/model"
	"github.com/tattwei46/inventory/server/param"
	"github.com/tattwei46/inventory/server/types"
)

type asset struct{}

var Asset asset

func (asset) ToModel(request *param.Asset) model.Asset {

	return model.Asset{
		ID:           types.GetRandom(),
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Brand:        request.Brand,
		Status:       types.Available,
		Created:      time.Now().Unix(),
	}
}

func (asset) ToParam(request *model.Asset) param.Asset {
	return param.Asset{
		ID:           request.ID,
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Brand:        request.Brand,
		Status:       request.Status.String(),
		Created:      types.Format(request.Created, types.TimeZone, types.YYYYMMDD_hhmmss),
	}
}

func (asset) ToParams(requests []model.Asset) []param.Asset {
	var result = make([]param.Asset, 0)

	for _, m := range requests {
		result = append(result, Asset.ToParam(&m))
	}

	return result
}

func (asset) ToModels(requests []param.Asset) []model.Asset {
	var result = make([]model.Asset, 0)

	for _, r := range requests {
		m := Asset.ToModel(&r)
		result = append(result, m)
	}

	return result
}
