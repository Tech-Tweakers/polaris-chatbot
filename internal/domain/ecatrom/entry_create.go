package ecatrom

import (
	"fmt"
	"polarisai/internal/domain/appcontext"
	"polarisai/internal/infrastructure/queryapi"
	"polarisai/internal/infrastructure/structx"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Creator interface {
	Create(ctx appcontext.Context, ecatromEntity structx.Messages, kind string) (*ChatPersistence, error)
}

func (l *ecatrom2000) Create(ctx appcontext.Context, ecatromEntity structx.Messages, kind string) (*ChatPersistence, error) {

	logger := ctx.Logger()

	if len(ecatromEntity) == 0 {
		return nil, DomainErrorFactory(BadRequest, "ecatromEntity must have at least one element")
	}

	logger.Info("Question from user", zap.String("content", ecatromEntity[0].Content), zap.String("where", "create"))

	var chatPersistenceUser ChatPersistence

	llamaInstTag := "[INST] " + ecatromEntity[0].Content + " [/INST]"

	l.LastEntryID++
	chatPersistenceUser.EntryID = l.LastEntryID
	chatPersistenceUser.ChatID = ecatromEntity[0].ChatID
	chatPersistenceUser.CreatedAt = time.Now()
	chatPersistenceUser.Role = ecatromEntity[0].Role
	chatPersistenceUser.Content = llamaInstTag

	if chatPersistenceUser.EntryID == 0000 {
		return nil, DomainErrorFactory(BadRequest, "entryID is required")
	}

	_, err := l.repository.Insert(chatPersistenceUser)
	if err != nil {
		logger.Error("error creating chat question from user", zap.Error(err), zap.String("where", "create"))
		return nil, err
	}

	logger.Info("Querying AI", zap.String("where", "create"))

	l.LastEntryID++

	dbData, _ := l.repository.List()
	chatPersistenceValues := *dbData

	chatPersistenceToSummary := func(persistence ChatPersistence) ChatSummary {
		if persistence.ChatID == ecatromEntity[0].ChatID || persistence.ChatID == "0000" {
			return ChatSummary{Role: persistence.Role, Content: persistence.Content}
		} else {
			return ChatSummary{}
		}
	}

	chatSummaries := make([]ChatSummary, len(chatPersistenceValues))
	for i, chatPersistence := range chatPersistenceValues {
		chatSummaries[i] = chatPersistenceToSummary(chatPersistence)
	}
	chatSumToString := fmt.Sprintf("%v", chatSummaries)
	chatSumToString = strings.Replace(chatSumToString, "Role:", "", -1)
	chatSumToString = strings.Replace(chatSumToString, "Content:", "", -1)
	chatSumToString = strings.Replace(chatSumToString, "{", "", -1)
	chatSumToString = strings.Replace(chatSumToString, "}", "", -1)

	fmt.Println(chatSumToString)

	l.LoadModel(kind)

	aiReply := queryapi.SendMessage(chatSumToString, l.aiModel, kind)
	aiReply = strings.Replace(aiReply, "assistant:", "", -1)
	chatPersistenceAi := ChatPersistence{
		CreatedAt: time.Now(),
		EntryID:   l.LastEntryID,
		ChatID:    ecatromEntity[0].ChatID,
		Role:      "assistant:",
		Content:   aiReply,
	}
	_, err = l.repository.Insert(chatPersistenceAi)
	if err != nil {
		logger.Error("error creating chat question from user", zap.Error(err), zap.String("where", "create"))
		return nil, err
	}

	return &chatPersistenceAi, nil
}
