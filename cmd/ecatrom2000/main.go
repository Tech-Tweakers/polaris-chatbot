package main

import (
	"ecatrom/internal/domain/ecatrom"
	"ecatrom/internal/infrastructure/api"
	"ecatrom/internal/infrastructure/aws"
	"ecatrom/internal/infrastructure/database"
	"ecatrom/internal/infrastructure/environment"
	"ecatrom/internal/infrastructure/logger"
	"ecatrom/internal/infrastructure/logger/logwrapper"

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
		zap.String("AWS_REGION", env.AWS_REGION),
		zap.String("AWS_ENDPOINT", env.AWS_ENDPOINT),
		zap.String("AWS_PROFILE", env.AWS_PROFILE),
		zap.String("DYNAMO_AWS_ENDPOINT", env.DYNAMO_AWS_ENDPOINT),
		zap.String("DYNAMO_TABLE_NAME", env.DYNAMO_TABLE_NAME),
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

	dynamodb, err := setupDynamoDB()
	if err != nil {
		return nil, err
	}

	memdbInput := &ecatrom.Input{
		Repository: dynamodb,
	}
	ecatrom2000UseCases = ecatrom.New(memdbInput)

	chatValue++
	ecatrom2000UseCases.StartChat(chatValue)

	return ecatrom2000UseCases, nil
}

func setupDynamoDB() (ecatrom.Repository, error) {
	env := environment.GetInstance()
	if env.DEFAULT_PERSISTENT == "false" {
		return database.NewMemoryDatabase(), nil
	}

	awsRegion := env.AWS_REGION
	awsEndpoint := env.DYNAMO_AWS_ENDPOINT
	table := env.DYNAMO_TABLE_NAME
	cfg, err := aws.EndpointResolverWithOptionsFunc(awsEndpoint, awsRegion)
	if err != nil {
		return nil, err
	}
	return database.NewDynamoDB(cfg, table), nil
}

func setupApi(logger logwrapper.LoggerWrapper, ecatrom2000UseCases ecatrom.UseCases) {
	input := api.Input{
		Logger:              logger,
		Ecatrom2000UseCases: ecatrom2000UseCases,
	}
	api.Start(input)
}
