package framework

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var config = viper.New()
var serviceMap map[Services]Service
var mongoConnectionMap map[Keys]MongoConnection
var logDir string

var serviceKeys = [...]Services{
	INVENTORY,
}

var mongoDBKeys = [...]Keys{
	MONGODB,
}

type MongoConnection struct {
	URL string
}
type Service struct {
	Name        string
	Host        string
	Port        int64
	URL         string
	LogFileName string
}

func GetLogDir() string {
	return logDir
}

func GetMongoDB(key Keys) MongoConnection {
	return mongoConnectionMap[key]
}

func LoadConfigFile() {
	c := os.Getenv("CONFIG_FILE_PATH")
	if len(c) <= 0 {
		// log.Fatalf("CONFIG_FILE_PATH environment variable is not set!!")
		c = "/var/www/golang/go-services-api/config"
		log.Printf("CONFIG_FILE_PATH environment variable is not set!! Application is setting the Default Config Path %s ", c)
	}

	config.SetConfigName("config")
	config.AddConfigPath(c)

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	fmt.Printf("Config File Path : %s\n", config.ConfigFileUsed())

	logDir = config.Get("log-directory").(string)
	readServices()
	readMongodbConnection()

	// Printing the Details Here
	printDetails()

}

func printDetails() {
	fmt.Printf("Printing Config Details \n")

	fmt.Printf("\n Service Informations")
	fmt.Printf("\n =========================\n ")

	for k, v := range serviceMap {
		fmt.Printf("%s Service Information \n", k)
		fmt.Printf("--------------------------- \n")
		fmt.Printf("Name : %s \n", v.Name)
		fmt.Printf("Host : %s \n", v.Host)
		fmt.Printf("Port : %d \n", v.Port)
		fmt.Printf("Log File Name : %s \n", v.LogFileName)
	}
	fmt.Printf("\n =========================\n ")

	fmt.Printf("Mongo Database Details \n")
	for k, v := range mongoConnectionMap {
		fmt.Printf("%s Database Information \n", k)
		fmt.Printf("--------------------------- \n")
		fmt.Printf("URL : xxxxxx%s \n", v.URL[len(v.URL)-20:])
	}
	fmt.Printf("\n =========================\n ")

}

func readServices() {
	serviceMap = make(map[Services]Service)
	for _, key := range serviceKeys {
		service, err := getServiceInfo(key.String())
		if err == nil {
			serviceMap[key] = *service
		}
	}
}

func GetServiceInfo(service Services) *Service {
	if val, ok := serviceMap[service]; ok {
		return &val
	}
	return &Service{}
}

func getServiceInfo(key string) (*Service, error) {
	service := config.GetStringMap(key)
	if len(service) == 0 {
		return nil, fmt.Errorf("service info '%s' not found", key)
	}

	name := service["name"].(string)
	host := service["host"].(string)
	port := service["port"].(int64)
	url := ""
	if service["url"] != nil {
		url = service["url"].(string)
	}

	logFileName := fmt.Sprintf("%s.log", name)
	if service["log-file-name"] != nil {
		logFileName = service["log-file-name"].(string)
	}
	return &Service{name, host, port, url, logFileName}, nil
}

func readMongodbConnection() {
	mongoConnectionMap = make(map[Keys]MongoConnection)
	for _, key := range mongoDBKeys {
		con, err := getMongoConnection(key.String())
		if err == nil {
			mongoConnectionMap[key] = *con
		}
	}
}

func getMongoConnection(key string) (*MongoConnection, error) {
	k := fmt.Sprintf("%s.url", key)
	con := config.Get(k)

	if con == nil {
		return nil, fmt.Errorf("%s Not found", key)
	}

	conString := con.(string)
	return &MongoConnection{conString}, nil

}
