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

	certFile := "./ssl-certs/polaris.crt"
	keyFile := "./ssl-certs/polaris.key"
	if err := r.RunTLS(applicationPort, certFile, keyFile); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}

func resolvePort() string {
	const CHAR string = ":"
	env := environment.GetInstance()
	port := env.APP_PORT
	firstChar := port[:0]
	if firstChar != CHAR {
		port = CHAR + "9001"
	}
	return port
}
