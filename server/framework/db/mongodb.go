package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/tattwei46/inventory/server/framework/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB MongoDB

type MongoDB struct {
	*mongo.Client
}

func NewMongoDB() (MongoDB, error) {
	clientOptions := options.Client().ApplyURI(config.GetMongoDB(config.MONGODB).URL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return mongoDB, err
	}

	mongoDB.Client = client
	fmt.Println("Connected to MongoDB!")
	return mongoDB, nil
}

func GetMongoDB() (MongoDB, error) {
	if mongoDB == (MongoDB{}) {
		return mongoDB, errors.New("mongodb not initialized")
	}
	return mongoDB, nil
}
