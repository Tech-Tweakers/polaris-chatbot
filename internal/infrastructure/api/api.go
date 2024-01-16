package api

import (
	"polarisai/internal/domain/polaris"
	"polarisai/internal/infrastructure/api/middlewares"
	"polarisai/internal/infrastructure/api/routes"
	"polarisai/internal/infrastructure/environment"
	"polarisai/internal/infrastructure/logger/logwrapper"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Input struct {
	Logger              logwrapper.LoggerWrapper
	Ecatrom2000UseCases polaris.UseCases
}

func Start(input Input) {

	env := environment.GetInstance()
	r := gin.New()

	logger := input.Logger
	logger.Info("Starting Polaris API")

	applicationPort := resolvePort()
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.ContextMiddleware())
	r.Use(middlewares.TraceMiddleware())
	r.Use(middlewares.Logger(logger))
	if !environment.GetInstance().IsDevelopment() {
		r.Use(middlewares.Recovery(true))
	}
	// r.Use(middlewares.Metrics(metricService)

	routes.MakeHealthRoute(r)
	routes.MakeMetricRoute(r)
	routes.MakeEntriesRoute(r, input.Ecatrom2000UseCases)

	enableSSL := env.ENABLE_SSL

	if enableSSL == "true" {
		// If SSL is enabled, use HTTPS
		if err := r.RunTLS(applicationPort, "ssl-certs/polaris.crt", "ssl-certs/polaris.key"); err != nil {
			logger.Fatal("failed to start server with TLS", zap.Error(err))
		}
	} else {
		// If SSL is disabled, use HTTP
		if err := r.Run(applicationPort); err != nil {
			logger.Fatal("failed to start server", zap.Error(err))
		}
	}

}

func resolvePort() string {
	const CHAR string = ":"
	env := environment.GetInstance()
	port := env.APP_PORT

	// Check if SSL should be enabled
	enableSSL := env.ENABLE_SSL // You need to define this environment variable

	if enableSSL == "true" {
		return CHAR + port // SSL enabled
	} else {
		firstChar := port[:0]
		if firstChar != CHAR {
			port = CHAR + "9001"
		}
		return port // SSL disabled
	}
}
