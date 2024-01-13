package logger

import (
	"encoding/json"
	"log"

	"polarisai/internal/infrastructure/environment"

	"go.uber.org/zap"
)

// New create logger
func New() (*zap.Logger, func()) {
	env := environment.GetInstance()
	var cfg zap.Config
	rawJSON := []byte(`{
			"level": "` + env.LOG_LEVEL + `",
			"encoding": "json",
			"outputPaths": ["stdout"],
			"errorOutputPaths": ["stderr"],
			"encoderConfig": {
			  "messageKey": "message",
			  "levelKey": "level",
			  "levelEncoder": "lowercase"
			}
		  }`)

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		log.Fatal(err.Error())
	}
	logger, err := cfg.Build()

	if err != nil {
		log.Fatal(err.Error())
	}

	undo := zap.ReplaceGlobals(logger)

	return logger, func() {
		/* #nosec */
		logger.Sync()
		undo()
	}
}
