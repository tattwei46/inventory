package repository_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/tattwei46/inventory/server/types"

	"github.com/tattwei46/inventory/server/framework/config"

	"github.com/stretchr/testify/assert"

	"github.com/tattwei46/inventory/server/model"

	"github.com/tattwei46/inventory/server/framework/db"

	"github.com/tattwei46/inventory/server/repository"
)

const (
	dbInventory = "inventory"
	collAssets  = "assets"
)

var mgoAsset repository.Asset
var mgoClient db.MongoDB
var timeNow = time.Now().Unix()
var timeThen = timeNow + 3600

var testModels = []model.Asset{
	{
		ID:           "0001",
		SerialNumber: "SerialNumber1",
		Model:        "Model1",
		Brand:        "Brand1",
		Status:       1,
		Created:      timeNow + 120,
	},
	{
		ID:           "0002",
		SerialNumber: "SerialNumber2",
		Model:        "Model2",
		Brand:        "Brand2",
		Status:       1,
		Created:      timeNow + 60,
	},
	{
		ID:           "0003",
		SerialNumber: "SerialNumber3",
		Model:        "Model3",
		Brand:        "Brand3",
		Status:       1,
		Created:      timeNow,
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

	mgoClient = mr

	database := mr.Client.Database(dbInventory)
	coll := database.Collection(collAssets)

	// Delete existing collection
	if err = coll.Drop(context.TODO()); err != nil {
		fmt.Println(err.Error())
	}

	// Init repository
	repo, err := repository.NewAsset()
	if err != nil {
		panic(err)
	}
	mgoAsset = repo

}

func teardown() {
	if err := mgoClient.Client.Disconnect(context.TODO()); err != nil {
		fmt.Println(err.Error())
	}
}

func TestAsset_Add(t *testing.T) {
	var request = make([]model.Asset, 0)
	request = append(request, testModels...)

	err := mgoAsset.Add(request)
	assert.NoError(t, err)
}

func TestAsset_GetAll(t *testing.T) {
	result, err := mgoAsset.Get(model.Search{}, 0, 0)
	assert.NoError(t, err)
	assert.EqualValues(t, len(testModels), len(result))
	for i, r := range result {
		assert.Equal(t, testModels[i], r)
	}
}

// Criteria 1.5 : Can be filter by Date Bought
func TestAsset_GetOne(t *testing.T) {
	search := model.Search{
		Range: types.RangeUnix{
			From: timeNow,
			To:   timeThen,
		},
	}

	result, err := mgoAsset.Get(search, 0, 0)
	assert.NoError(t, err)
	assert.EqualValues(t, len(testModels), len(result))
	for i, r := range result {
		assert.Equal(t, r, testModels[i])
	}
}

func TestAsset_Delete(t *testing.T) {
	for _, r := range testModels {
		count, err := mgoAsset.Delete(r.ID)
		assert.NoError(t, err)
		assert.EqualValues(t, 1, count)
	}

	// search one should get empty
	search := model.Search{
		Brand: testModels[1].Brand,
	}

	result, err := mgoAsset.Get(search, 0, 0)
	assert.EqualValues(t, 0, len(result))
	assert.NoError(t, err)
}
