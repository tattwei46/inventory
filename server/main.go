package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/tattwei46/inventory/server/framework/db"

	"github.com/tattwei46/inventory/server/framework/logger"

	"github.com/gin-gonic/gin"
	"github.com/tattwei46/inventory/server/api"
	"github.com/tattwei46/inventory/server/framework/config"
)

func main() {

	// 1. Load configuration from config.toml
	config.Load()

	// 2. Initialize mongodb
	if mr, err := db.NewMongoDB(); err != nil {
		log.Fatal(err)
	} else {
		defer mr.Client.Disconnect(context.TODO())
	}

	// 3. Get service info
	service := getService()

	// 4. Initialize logger
	logger.InitLogger(filepath.Join(config.GetLogDir(), service.LogFileName), service.Name)
	logger := logger.GetInstance()
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

func setupRouter() *gin.Engine {
	gin.DisableConsoleColor()

	// output gin log into file and console
	gin.DefaultWriter = io.MultiWriter(logger.GetInstance().Get(), os.Stdout)

	r := gin.Default()

	api.Base().Routes(r)

	api.Asset().Routes(r)
	return r
}
