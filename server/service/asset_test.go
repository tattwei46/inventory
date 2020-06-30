package service_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

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

func TestAsset_Add(t *testing.T) {
	err := asset.Add(testParams)
	assert.NoError(t, err)
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

func TestAsset_GetOne(t *testing.T) {
	search := param.Search{
		Brand: "Brand1",
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

func TestAsset_Delete(t *testing.T) {
	for _, id := range idList {
		count, err := asset.Delete(id)
		assert.NoError(t, err)
		assert.EqualValues(t, 1, count)
	}
}
