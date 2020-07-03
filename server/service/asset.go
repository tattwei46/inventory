package service

import (
	"fmt"

	"github.com/tattwei46/inventory/server/param"
	"github.com/tattwei46/inventory/server/repository"
	"github.com/tattwei46/inventory/server/service/converter"
	"github.com/tattwei46/inventory/server/types"
)

type Asset interface {
	Add(param.Asset) error
	Get(param.Search, int, int) ([]param.Asset, error)
	Delete(id string) (int64, error)
	Update(id string, update param.Asset) error
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

func (s *asset) Add(requests param.Asset) error {
	m := converter.Asset.ToModel(&requests)
	return s.Asset.Add(m)
}

func (s *asset) Delete(id string) (int64, error) {
	return s.Asset.Delete(id)
}

func (s *asset) Update(id string, update param.Asset) error {
	toUpdate := converter.Update.ToModel(&update)

	fmt.Println(toUpdate)

	search := param.Search{
		Range: types.Range{
			From: update.Created,
			To:   update.Created,
		},
		SerialNumber: update.SerialNumber,
		Brand:        update.Brand,
		Model:        update.Model,
		Status:       update.Status,
	}

	res, err := s.Get(search, 1, 1)
	if err != nil {
		return fmt.Errorf("an error occured when check item exist before update %v", err)
	}

	if len(res) > 0 {
		return types.DuplicatedItem
	}

	return s.Asset.Update(id, toUpdate)
}
