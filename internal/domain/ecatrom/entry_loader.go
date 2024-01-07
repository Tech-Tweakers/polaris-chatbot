package ecatrom

import (
	"ecatrom/internal/infrastructure/queryapi"

	"github.com/go-skynet/go-llama.cpp"
)

type Loader interface {
	LoadModel(kind string) (m *llama.LLama)
}

func (l *ecatrom2000) LoadModel(kind string) (ml *llama.LLama) {
	ml = queryapi.LoadAiModel(kind)
	l.aiModel = ml
	return l.aiModel
}
