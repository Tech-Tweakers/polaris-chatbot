package environment

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

// Single is the singleton instance of the environment
type Single struct {
	ENVIRONMENT        string // nolint: golint
	APP_VERSION        string // nolint: golint
	APP_PORT           string // nolint: golint
	APP_URL            string // nolint: golint
	LOG_LEVEL          string // nolint: golint
	DEFAULT_PERSISTENT string // nolint: golint
	MODEL_PATH         string // nolint: golint
	DBNAME             string // nolint: golint
	COLLECTION_NAME    string // nolint: golint
	CONNECTION_STRING  string // nolint: golint
}

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("Error loading .env.local file")
	}
	env := GetInstance()
	env.Setup()
}

func (e *Single) Setup() {
	e.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	e.APP_VERSION = os.Getenv("APP_VERSION")
	e.APP_PORT = os.Getenv("APP_PORT")
	e.APP_URL = os.Getenv("APP_URL")
	e.LOG_LEVEL = os.Getenv("LOG_LEVEL")

	e.DEFAULT_PERSISTENT = os.Getenv("DEFAULT_PERSISTENT")
	e.DBNAME = os.Getenv("DBNAME")
	e.COLLECTION_NAME = os.Getenv("COLLECTION_NAME")
	e.CONNECTION_STRING = os.Getenv("CONNECTION_STRING")

	e.MODEL_PATH = os.Getenv("MODEL_PATH")
}

func (e *Single) IsDevelopment() bool {
	return e.ENVIRONMENT == "development"
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
			singleInstance.Setup()
		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return singleInstance
}
