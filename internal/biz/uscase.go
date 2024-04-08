package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPassportUseCase)

type HelloRepo interface {
}

type HelloUseCase struct {
	repo HelloRepo
}

func NewPassportUseCase(repo HelloRepo) *HelloUseCase {
	return &HelloUseCase{
		repo: repo,
	}
}

func (uc *HelloUseCase) Hello(ctx context.Context, message string) (string, error) {
	return "hello" + message, nil
}
