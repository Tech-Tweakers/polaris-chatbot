package ecatrom

import (
	"ecatrom/internal/infrastructure/logger/logwrapper"
	"ecatrom/internal/infrastructure/queryapi"

	"github.com/go-skynet/go-llama.cpp"
)

type Loader interface {
	LoadModel() (m *llama.LLama)
	LoadLogger(logwrapper.LoggerWrapper)
}

func (l *ecatrom2000) LoadModel() (ml *llama.LLama) {
	l.logger.Info("Loading AI Model")
	ml = queryapi.LoadAiModel()
	l.aiModel = ml
	return l.aiModel
}

func (l *ecatrom2000) LoadLogger(logger logwrapper.LoggerWrapper) {
	l.logger = logger
	l.logger.Info("Loading Logger")

}
