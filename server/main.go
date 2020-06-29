package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/tattwei46/inventory/framework"

	"github.com/tattwei46/inventory/api"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	framework.LoadConfigFile()

	if mr, err := initMongo(); err != nil {
		log.Fatal(err)
	} else {
		defer mr.Disconnect(context.TODO())
	}

	service := getService()
	framework.InitLogger(filepath.Join(framework.GetLogDir(), service.LogFileName), service.Name)
	logger := framework.GetLoggerInstance()
	logger.Info("logger initialized")
	router := setupRouter()

	if err := router.Run(fmt.Sprintf("%s:%s", service.Host, service.Port)); err != nil {
		logger.Error(err)
	}
}

func getService() *framework.Service {
	service := framework.GetServiceInfo(framework.INVENTORY)
	if len(service.Host) <= 0 {
		service.Host = "0.0.0.0"
		service.Port = 15701
		service.Name = "inventory"
		service.LogFileName = "inventory.log"
	}
	return service
}

func initMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(framework.GetMongoDB(framework.MONGODB).URL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return client, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func setupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(framework.GetLoggerInstance().Get(), os.Stdout)

	r := gin.Default()

	api.Base().Routes(r)
	return r
}
