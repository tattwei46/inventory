package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/tattwei46/inventory/framework/logger"

	"github.com/gin-gonic/gin"
	"github.com/tattwei46/inventory/api"
	"github.com/tattwei46/inventory/framework/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// 1. Load configuration from config.toml
	config.Load()

	// 2. Initialize mongodb
	if mr, err := initMongo(); err != nil {
		log.Fatal(err)
	} else {
		defer mr.Disconnect(context.TODO())
	}

	// 3. Get service info
	service := getService()

	// 4. Initialize logger
	logger.InitLogger(filepath.Join(config.GetLogDir(), service.LogFileName), service.Name)
	logger := logger.GetLoggerInstance()
	logger.Info("logger initialized")
	router := setupRouter()

	if err := router.Run(fmt.Sprintf("%s:%d", service.Host, service.Port)); err != nil {
		logger.Error(err)
	}
}

func getService() *config.Service {
	service := config.GetServiceInfo(config.INVENTORY)
	if len(service.Host) <= 0 {
		service.Host = "0.0.0.0"
		service.Port = 15701
		service.Name = "inventory"
		service.LogFileName = "inventory.log"
	}
	return service
}

func initMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.GetMongoDB(config.MONGODB).URL)
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
	gin.DefaultWriter = io.MultiWriter(logger.GetLoggerInstance().Get(), os.Stdout)

	r := gin.Default()

	api.Base().Routes(r)
	return r
}
