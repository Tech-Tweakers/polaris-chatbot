package ecatrom

import (
	"polarisai/internal/domain/appcontext"

	"go.uber.org/zap"
)

type Lister interface {
	ListAll(ctx appcontext.Context) (*[]ChatPersistence, error)
}

func (l *ecatrom2000) ListAll(ctx appcontext.Context) (*[]ChatPersistence, error) {
	logger := ctx.Logger()
	logger.Info("Listing entries", zap.String("where", "listall"))

	result, err := l.repository.List()

	return result, err
}
