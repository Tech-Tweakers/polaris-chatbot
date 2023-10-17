package ecatrom

import (
	"ecatrom/internal/domain/appcontext"
	"ecatrom/internal/infrastructure/queryapi"
	"ecatrom/internal/infrastructure/structx"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Creator interface {
	Create(ctx appcontext.Context, ecatromEntity structx.Messages) (*ChatPersistence, error)
}

func (l *ecatrom2000) Create(ctx appcontext.Context, ecatromEntity structx.Messages) (*ChatPersistence, error) {

	logger := ctx.Logger()

	logger.Info("Question from user", zap.String("content", ecatromEntity[0].Content), zap.String("where", "create"))

	var chatPersistenceUser ChatPersistence

	l.LastEntryID++
	chatPersistenceUser.EntryID = l.LastEntryID

	chatPersistenceUser.CreatedAt = time.Now()
	chatPersistenceUser.Role = ecatromEntity[0].Role
	chatPersistenceUser.Content = ecatromEntity[0].Content

	if chatPersistenceUser.EntryID == 0000 {
		return nil, DomainErrorFactory(BadRequest, "entryID is required")
	}

	_, err := l.repository.Insert(chatPersistenceUser)
	if err != nil {
		logger.Error("error creating chat question from user", zap.Error(err), zap.String("where", "create"))
		return nil, err
	}

	// CALL MSGS PARSER TO AI

	dbData, _ := l.repository.List()

	logger.Info("Querying AI", zap.String("where", "create"))

	l.LastEntryID++

	chatPersistenceValues := *dbData
	chatPersistenceToSummary := func(persistence ChatPersistence) ChatSummary {
		return ChatSummary{Role: persistence.Role, Content: persistence.Content}
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

	l.LoadModel()

	aiReply := queryapi.SendMessage(chatSumToString, l.aiModel)
	aiReply = strings.Replace(aiReply, "assistant:", "", -1)
	chatPersistenceAi := ChatPersistence{
		CreatedAt: time.Now(),
		EntryID:   l.LastEntryID,
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
