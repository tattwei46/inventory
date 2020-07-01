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

// Criteria 3.1: Can update brand
func TestAsset_UpdateBrand(t *testing.T) {

	// Get existing ID
	search := param.Search{
		SerialNumber: testParams[1].SerialNumber, // SerialNumber1,
	}
	result, _ := asset.Get(search, 1, 1)

	id := result[0].ID

	// Update existing ID
	update := param.Asset{
		Brand: "BrandNew2",
	}

	err := asset.Update(id, update)
	assert.NoError(t, err)

	// Get updated item
	result, _ = asset.Get(search, 1, 1)
	assert.Equal(t, update.Brand, result[0].Brand)

	// Update existing ID
	update = param.Asset{
		Brand: testParams[1].Brand,
	}

	asset.Update(id, update)

}

// Criteria 3.2: Can update model
func TestAsset_UpdateModel(t *testing.T) {

	// Get existing ID
	search := param.Search{
		SerialNumber: testParams[1].SerialNumber, // SerialNumber1,
	}
	result, _ := asset.Get(search, 1, 1)

	id := result[0].ID

	// Update existing ID
	update := param.Asset{
		Model: "ModelNew2",
	}

	err := asset.Update(id, update)
	assert.NoError(t, err)

	// Get updated item
	result, _ = asset.Get(search, 1, 1)
	assert.Equal(t, update.Model, result[0].Model)

	// Update existing ID
	update = param.Asset{
		Model: testParams[1].Model,
	}

	asset.Update(id, update)

}

// Criteria 3.3: Can update serial number
func TestAsset_UpdateSerialNumber(t *testing.T) {

	// Get existing ID
	search := param.Search{
		SerialNumber: testParams[1].SerialNumber, // SerialNumber2,
	}
	result, _ := asset.Get(search, 1, 1)

	id := result[0].ID

	// Update existing ID
	update := param.Asset{
		SerialNumber: "SerialNumberNew2",
	}

	err := asset.Update(id, update)
	assert.NoError(t, err)

	search = param.Search{
		SerialNumber: update.SerialNumber, // SerialNumber2,
	}
	// Get updated item
	result, _ = asset.Get(search, 1, 1)
	assert.Equal(t, update.SerialNumber, result[0].SerialNumber)

	// Update existing ID
	update = param.Asset{
		SerialNumber: testParams[1].SerialNumber,
	}

	asset.Update(id, update)

}

// Criteria 3.4: Can update status
func TestAsset_UpdateStatus(t *testing.T) {

	// Get existing ID
	search := param.Search{
		SerialNumber: testParams[1].SerialNumber, // SerialNumber1,
	}
	result, _ := asset.Get(search, 1, 1)

	id := result[0].ID

	// Update existing ID
	update := param.Asset{
		Status: types.NotAvailable.String(),
	}

	err := asset.Update(id, update)
	assert.NoError(t, err)

	// Get updated item
	result, _ = asset.Get(search, 1, 1)
	assert.Equal(t, update.Status, result[0].Status)

	// Update existing ID
	update = param.Asset{
		Status: testParams[1].Status,
	}

	asset.Update(id, update)

}

// Criteria 3.5: Can update date bought
func TestAsset_UpdateDateBought(t *testing.T) {

	// Get existing ID
	search := param.Search{
		SerialNumber: testParams[1].SerialNumber, // SerialNumber1,
	}
	result, _ := asset.Get(search, 1, 1)

	id := result[0].ID
	created := result[0].Created

	// Update existing ID
	update := param.Asset{
		Created: "2020-07-01 00:00:00",
	}

	err := asset.Update(id, update)
	assert.NoError(t, err)

	// Get updated item
	result, _ = asset.Get(search, 1, 1)
	assert.Equal(t, update.Created, result[0].Created)

	// Update existing ID
	update = param.Asset{
		Created: created,
	}

	asset.Update(id, update)

}

// Criteria 3 : Can delete
func TestAsset_Delete(t *testing.T) {
	for _, id := range idList {
		count, err := asset.Delete(id)
		assert.NoError(t, err)
		assert.EqualValues(t, 1, count)
	}
}
