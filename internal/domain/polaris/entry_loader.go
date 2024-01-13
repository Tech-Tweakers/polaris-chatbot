package polaris

import (
	"polarisai/internal/infrastructure/queryapi"

	llama "github.com/go-skynet/go-llama.cpp"
)

type Loader interface {
	LoadModel(kind string) (m *llama.LLama)
}

func (l *polaris) LoadModel(kind string) (ml *llama.LLama) {
	ml = queryapi.LoadAiModel(kind)
	l.aiModel = ml
	return l.aiModel
}
