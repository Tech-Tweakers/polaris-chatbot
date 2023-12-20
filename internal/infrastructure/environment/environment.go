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
	ENVIRONMENT         string // nolint: golint
	APP_VERSION         string // nolint: golint
	APP_PORT            string // nolint: golint
	APP_URL             string // nolint: golint
	AWS_ENDPOINT        string // nolint: golint
	AWS_REGION          string // nolint: golint
	AWS_PROFILE         string // nolint: golint
	LOG_LEVEL           string // nolint: golint
	DYNAMO_AWS_ENDPOINT string // nolint: golint
	DYNAMO_TABLE_NAME   string // nolint: golint
	DEFAULT_PERSISTENT  string // nolint: golint
	MODEL_PATH          string // nolint: golint
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

	e.AWS_REGION = os.Getenv("AWS_REGION")
	e.AWS_ENDPOINT = os.Getenv("AWS_ENDPOINT")
	e.AWS_PROFILE = os.Getenv("AWS_PROFILE")

	e.LOG_LEVEL = os.Getenv("LOG_LEVEL")

	e.DEFAULT_PERSISTENT = os.Getenv("DEFAULT_PERSISTENT")

	e.DYNAMO_AWS_ENDPOINT = os.Getenv("DYNAMO_AWS_ENDPOINT")
	e.DYNAMO_TABLE_NAME = os.Getenv("DYNAMO_TABLE_NAME")

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
