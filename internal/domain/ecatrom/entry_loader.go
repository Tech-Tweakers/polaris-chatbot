package ecatrom

import (
	"ecatrom/internal/infrastructure/queryapi"

	"github.com/go-skynet/go-llama.cpp"
)

type Loader interface {
	LoadModel() (m *llama.LLama)
}

func (l *ecatrom2000) LoadModel() (ml *llama.LLama) {
	ml = queryapi.LoadAiModel()
	l.aiModel = ml
	return l.aiModel
}
