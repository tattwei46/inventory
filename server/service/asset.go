package service

import (
	"github.com/tattwei46/inventory/server/param"
	"github.com/tattwei46/inventory/server/repository"
	"github.com/tattwei46/inventory/server/service/converter"
)

type Asset interface {
	Add(param.Asset) error
	Get(param.Search, int, int) ([]param.Asset, error)
}

type asset struct {
	repository.Asset
}

func NewAsset() (Asset, error) {
	repo, err := repository.NewAsset()
	if err != nil {
		return nil, err
	}

	return &asset{repo}, nil
}

func (s *asset) Get(search param.Search, limit, page int) ([]param.Asset, error) {
	var result = make([]param.Asset, 0)

	filter := converter.Search.ToModel(search)
	m, err := s.Asset.Get(filter, limit, page)
	if err != nil {
		return result, err
	}

	return converter.Asset.ToParams(m), nil
}

func (s *asset) Add(request param.Asset) error {
	m := converter.Asset.ToModel(request)
	return s.Asset.Add(m)
}
