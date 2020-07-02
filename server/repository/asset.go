package repository

import (
	"context"
	"errors"

	"github.com/tattwei46/inventory/server/types"

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
	Delete(id string) (int64, error)
	Update(id string, update model.Asset) error
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

	var filter = bson.D{}

	if !search.Range.IsEmpty() {
		filter = append(filter, bson.E{Key: "created", Value: bson.M{"$gte": search.Range.From}})
		filter = append(filter, bson.E{Key: "created", Value: bson.M{"$lte": search.Range.To}})
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

	return result, nil
}

// TODO : CHANGE TO ADD MANY
func (r *asset) Add(req model.Asset) error {
	// Find if item is duplicated
	search := model.Search{
		Brand:        req.Brand,
		SerialNumber: req.SerialNumber,
		Model:        req.Model,
	}

	found, err := r.Get(search, 0, 0)
	if err != nil {
		return err
	}

	if len(found) > 0 {
		return types.DuplicatedItem
	}

	result, err := r.coll.InsertOne(context.TODO(), req)
	if err != nil {
		return err
	}

	if _, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("failed to get oid")
	}
	return nil
}

func (r *asset) Delete(id string) (int64, error) {

	if len(id) <= 0 {
		return 0, nil
	}

	f := bson.D{
		{"id", id},
	}

	res, err := r.coll.DeleteOne(context.TODO(), f)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func (r *asset) Update(id string, update model.Asset) error {
	// Search if item exists in db
	search := model.Search{
		ID: id,
	}
	found, err := r.Get(search, 0, 0)
	if err != nil {
		return err
	}
	if len(found) <= 0 {
		return types.NoItem
	}

	filter := bson.D{{"id", id}}
	var condition bson.D

	if update.Created > 0 {
		condition = append(condition, bson.E{Key: "created", Value: update.Created})
	}

	if len(update.Brand) > 0 {
		condition = append(condition, bson.E{Key: "brand", Value: update.Brand})
	}

	if len(update.Model) > 0 {
		condition = append(condition, bson.E{Key: "model", Value: update.Model})
	}

	if len(update.SerialNumber) > 0 {
		condition = append(condition, bson.E{Key: "serial_number", Value: update.SerialNumber})
	}

	if update.Status > 0 {
		condition = append(condition, bson.E{Key: "status", Value: update.Status})
	}

	updateCond := bson.D{
		{"$set", condition},
	}

	_, err = r.coll.UpdateOne(context.TODO(), filter, updateCond)

	if err != nil {
		return err
	}

	return nil

}
