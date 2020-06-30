package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tattwei46/inventory/server/framework/db"
	"github.com/tattwei46/inventory/server/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbInventory = "inventory"
	collAssets  = "assets"
)

type Asset interface {
	Get(search model.Search, limit, offset int) ([]model.Asset, error)
	Add(model.Asset) error
}

type asset struct {
	mgo  *mongo.Client
	db   *mongo.Database
	coll *mongo.Collection
}

func NewAsset() (Asset, error) {
	mr, err := db.GetMongoDB()
	if err != nil {
		return nil, err
	}

	database := mr.Client.Database(dbInventory)
	coll := database.Collection(collAssets)

	return &asset{mr.Client, database, coll}, nil
}

func (r *asset) Get(search model.Search, limit, offset int) ([]model.Asset, error) {
	var result = make([]model.Asset, 0)

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created", -1}}) // sort by created descending

	if limit > 0 && offset > 0 { // set pagination
		findOptions.SetLimit(int64(limit))
		findOptions.SetSkip(int64((offset - 1) * limit))
	}

	// TODO : ADD FILTER
	var filter = bson.D{}

	if !search.Range.IsEmpty() {
		filter = append(filter, bson.E{Key: "created", Value: bson.M{"$gte": search.Range.From}})
	}

	if search.Status > 0 {
		filter = append(filter, bson.E{Key: "status", Value: search.Status})
	}

	if len(search.Model) > 0 {
		filter = append(filter, bson.E{Key: "model", Value: search.Model})
	}

	if len(search.Brand) > 0 {
		filter = append(filter, bson.E{Key: "brand", Value: search.Brand})
	}

	if len(search.SerialNumber) > 0 {
		filter = append(filter, bson.E{Key: "serial_number", Value: search.SerialNumber})
	}

	// Get Result
	cur, err := r.coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return result, err
	}

	for cur.Next(context.TODO()) {
		var asset model.Asset
		if err := cur.Decode(&asset); err != nil {
			return result, err
		}

		result = append(result, asset)
	}

	fmt.Println(result)

	fmt.Println(filter)

	return result, nil
}

// TODO : CHANGE TO ADD MANY
func (r *asset) Add(request model.Asset) error {
	result, err := r.coll.InsertOne(context.TODO(), request)
	if err != nil {
		return err
	}

	if _, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("failed to get oid")
	}

	return nil
}
