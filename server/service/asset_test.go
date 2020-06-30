package service_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/tattwei46/inventory/server/types"

	"github.com/stretchr/testify/assert"

	"github.com/tattwei46/inventory/server/param"

	"github.com/tattwei46/inventory/server/service"

	"github.com/tattwei46/inventory/server/framework/config"
	"github.com/tattwei46/inventory/server/framework/db"
)

const (
	dbInventory = "inventory"
	collAssets  = "assets"
)

var asset service.Asset

var idList = make([]string, 0)

var testParams = []param.Asset{
	{
		SerialNumber: "SerialNumber1",
		Brand:        "Brand1",
		Model:        "Model1",
	},
	{
		SerialNumber: "SerialNumber2",
		Brand:        "Brand2",
		Model:        "Model2",
	},
	{
		SerialNumber: "SerialNumber3",
		Brand:        "Brand3",
		Model:        "Model3",
	},
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// Load config
	config.Load()

	// Init mongo
	mr, err := db.NewMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	database := mr.Client.Database(dbInventory)
	coll := database.Collection(collAssets)

	// Delete existing collection
	if err = coll.Drop(context.TODO()); err != nil {
		fmt.Println(err.Error())
	}

	s, err := service.NewAsset()
	if err != nil {
		panic(err)
	}

	asset = s

}

func teardown() {
}

// Criteria 2.1 : Can Add if not exists
func TestAsset_Add(t *testing.T) {
	err := asset.Add(testParams)
	assert.NoError(t, err)
}

// Criteria 2.2 : Cannot add if exist same serial number, brand and model
func TestAsset_CannotAddIfExist(t *testing.T) {
	err := asset.Add(testParams)
	assert.Error(t, err)
}

func TestAsset_GetAll(t *testing.T) {
	result, err := asset.Get(param.Search{}, 0, 0)
	assert.NoError(t, err)
	for i, r := range result {
		assert.Equal(t, testParams[i].Brand, r.Brand)
		assert.Equal(t, testParams[i].Model, r.Model)
		assert.Equal(t, testParams[i].SerialNumber, r.SerialNumber)
		idList = append(idList, r.ID)
	}
}

// Criteria 1.1 : Can be filter by Serial Number
func TestAsset_FilterBySerialNumber(t *testing.T) {
	search := param.Search{
		SerialNumber: "SerialNumber1",
	}
	result, err := asset.Get(search, 0, 0)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(result))
	for _, r := range result {
		assert.Equal(t, testParams[0].Brand, r.Brand)
		assert.Equal(t, testParams[0].Model, r.Model)
		assert.Equal(t, testParams[0].SerialNumber, r.SerialNumber)
	}
}

// Criteria 1.2 : Can be filter by Brand
func TestAsset_FilterByBrand(t *testing.T) {
	search := param.Search{
		Brand: "Brand2",
	}
	result, err := asset.Get(search, 0, 0)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(result))
	for _, r := range result {
		assert.Equal(t, testParams[1].Brand, r.Brand)
		assert.Equal(t, testParams[1].Model, r.Model)
		assert.Equal(t, testParams[1].SerialNumber, r.SerialNumber)
	}
}

// Criteria 1.3 : Can be filter by Model
func TestAsset_FilterByModel(t *testing.T) {
	search := param.Search{
		Model: "Model3",
	}
	result, err := asset.Get(search, 0, 0)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(result))
	for _, r := range result {
		assert.Equal(t, testParams[2].Brand, r.Brand)
		assert.Equal(t, testParams[2].Model, r.Model)
		assert.Equal(t, testParams[2].SerialNumber, r.SerialNumber)
	}
}

// Criteria 1.4 : Can be filter by Status
func TestAsset_FilterByStatus(t *testing.T) {
	search := param.Search{
		Status: types.Available.String(),
	}
	result, err := asset.Get(search, 0, 0)
	assert.NoError(t, err)
	assert.EqualValues(t, len(testParams), len(result))
	for i, r := range result {
		assert.Equal(t, testParams[i].Brand, r.Brand)
		assert.Equal(t, testParams[i].Model, r.Model)
		assert.Equal(t, testParams[i].SerialNumber, r.SerialNumber)
	}
}

// Criteria 3 : Can delete
func TestAsset_Delete(t *testing.T) {
	for _, id := range idList {
		count, err := asset.Delete(id)
		assert.NoError(t, err)
		assert.EqualValues(t, 1, count)
	}
}
