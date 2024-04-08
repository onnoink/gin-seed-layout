package service

import (
	"context"
	"github.com/google/wire"
	v1 "github.com/onnoink/gin-seed-layout/api/helloworld/v1"
	"github.com/onnoink/gin-seed-layout/internal/biz"
	"github.com/onnoink/gin-seed-layout/internal/server"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewService)

type Service struct {
	uc *biz.HelloUseCase
}

var _ server.Service = (*Service)(nil)

func NewService(uc *biz.HelloUseCase) *Service {
	return &Service{
		uc: uc,
	}
}

type AuthorizationWechatCodeReply struct {
}

func (s *Service) Hello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	rMsg, err := s.uc.Hello(ctx, req.Message)
	if err != nil {
		return nil, err
	}

	return &v1.HelloReply{
		ReplyMessage: rMsg,
	}, nil
}
