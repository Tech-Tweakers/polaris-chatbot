package main

import (
	"context"
	"polarisai/internal/domain/polaris"
	"polarisai/internal/infrastructure/api"
	"polarisai/internal/infrastructure/database"
	"polarisai/internal/infrastructure/environment"
	"polarisai/internal/infrastructure/logger"
	"polarisai/internal/infrastructure/logger/logwrapper"
	"time"

	"go.uber.org/zap"
)

func main() {

	env := environment.GetInstance()
	zaplogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zaplogger}).Version(env.APP_VERSION)
	logger.Info("Starting Polaris Chat API")

	logger.Info("env",
		zap.String("LOG_LEVEL", env.LOG_LEVEL),
		zap.String("DEFAULT_PERSISTENT", env.DEFAULT_PERSISTENT),
		zap.String("APP_PORT", env.APP_PORT),
		zap.String("ENVIRONMENT", env.ENVIRONMENT),
		zap.String("APP_VERSION", env.APP_VERSION),
		zap.String("APP_URL", env.APP_URL),
	)

	polarisUseCases, err := setupPolaris(logger)

	var chatValue float64 = 0000
	polarisUseCases.StartChat(chatValue)

	if err != nil {
		logger.Error("failed to setup polaris-chat", zap.Error(err))
	}

	// setupWorker(logger, ecatrom2000UseCases)
	setupApi(logger, polarisUseCases)

}

func setupPolaris(logger logwrapper.LoggerWrapper) (polarisUseCases polaris.UseCases, err error) {
	var chatValue float64 = 0000

	mongodb, err := setupMongoDB()
	if err != nil {
		return nil, err
	}
	mongodbInput := &polaris.Input{
		Repository: mongodb,
	}
	polarisUseCases = polaris.New(mongodbInput)

	chatValue++

	return polarisUseCases, nil
}

func setupApi(logger logwrapper.LoggerWrapper, ecatrom2000UseCases polaris.UseCases) {
	input := api.Input{
		Logger:              logger,
		Ecatrom2000UseCases: ecatrom2000UseCases,
	}
	api.Start(input)
}

func setupMongoDB() (polaris.Repository, error) {
	env := environment.GetInstance()
	if env.DEFAULT_PERSISTENT == "false" {
		return database.NewMemoryDatabase(), nil
	}

	connectionString := env.CONNECTION_STRING
	dbName := env.DBNAME
	collectionName := env.COLLECTION_NAME

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return database.NewMongoDB(ctx, connectionString, dbName, collectionName)
}
