package ecatrom

import (
	"ecatrom/internal/infrastructure/logger/logwrapper"

	"github.com/go-skynet/go-llama.cpp"
)

type Input struct {
	Repository Repository
}

type ecatrom2000 struct {
	repository  Repository
	LastEntryID float64
	aiModel     *llama.LLama
	logger      logwrapper.LoggerWrapper
}

type UseCases interface {
	Creator
	Lister
	Loader
	Starter
}

func New(input *Input) UseCases {
	return &ecatrom2000{
		repository:  input.Repository,
		LastEntryID: 0000,
		logger:      nil,
	}
}
