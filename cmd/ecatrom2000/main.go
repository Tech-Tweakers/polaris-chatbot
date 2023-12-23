package main

import (
	"context"
	"ecatrom/internal/domain/ecatrom"
	"ecatrom/internal/infrastructure/api"
	"ecatrom/internal/infrastructure/database"
	"ecatrom/internal/infrastructure/environment"
	"ecatrom/internal/infrastructure/logger"
	"ecatrom/internal/infrastructure/logger/logwrapper"
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

	ecatrom2000UseCases, err := setupecatrom2000(logger)

	if err != nil {
		logger.Error("failed to setup polaris-chat", zap.Error(err))
	}

	// setupWorker(logger, ecatrom2000UseCases)
	setupApi(logger, ecatrom2000UseCases)

}

func setupecatrom2000(logger logwrapper.LoggerWrapper) (ecatrom2000UseCases ecatrom.UseCases, err error) {
	var chatValue float64 = 0000

	// For MongoDB, you can use setupMongoDB() here to switch to MongoDB
	mongodb, err := setupMongoDB()
	if err != nil {
		return nil, err
	}
	mongodbInput := &ecatrom.Input{
		Repository: mongodb,
	}
	ecatrom2000UseCases = ecatrom.New(mongodbInput)

	chatValue++
	ecatrom2000UseCases.StartChat(chatValue)

	return ecatrom2000UseCases, nil
}

func setupApi(logger logwrapper.LoggerWrapper, ecatrom2000UseCases ecatrom.UseCases) {
	input := api.Input{
		Logger:              logger,
		Ecatrom2000UseCases: ecatrom2000UseCases,
	}
	api.Start(input)
}

func setupMongoDB() (ecatrom.Repository, error) {
	env := environment.GetInstance()
	if env.DEFAULT_PERSISTENT == "false" {
		return database.NewMemoryDatabase(), nil
	}

	connectionString := "mongodb://root:examplepassword@localhost:27017/polaris?authSource=admin"

	dbName := "polaris"
	collectionName := "polaris-collection"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return database.NewMongoDB(ctx, connectionString, dbName, collectionName)
}
