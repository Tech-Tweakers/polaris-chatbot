package polaris

import (
	"polarisai/internal/infrastructure/environment"
	"time"
)

type Starter interface {
	StartChat(chatValue float64) error
}

func (l *polaris) StartChat(chatValue float64) error {

	env := environment.GetInstance()

	l.LastEntryID++
	initInstruction := env.AI_SYSTEM_INSTRUCTION

	initStruct := ChatPersistence{
		CreatedAt: time.Now(),
		ChatID:    l.ChatID,
		EntryID:   l.LastEntryID,
		Role:      "system:",
		Content:   initInstruction,
	}

	l.repository.Insert(initStruct)

	return nil
}
