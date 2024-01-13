package ecatrom

import (
	llama "github.com/go-skynet/go-llama.cpp"
)

type Input struct {
	Repository Repository
}

type ecatrom2000 struct {
	repository  Repository
	LastEntryID float64
	ChatID      string
	aiModel     *llama.LLama
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
		ChatID:      "0000",
		LastEntryID: 0000,
	}
}
