package service

import (
	"context"
	"fmt"
	"log/slog"
	"slog-sample/domain/infra_interface"
)

type Service interface {
	DoSomething(ctx context.Context)
}

type service struct {
	repository infra_interface.Repository
}

var _ Service = &service{}

func NewService(repository infra_interface.Repository) Service {
	return &service{repository}
}

func (s *service) DoSomething(ctx context.Context) {
	id := s.repository.Get(ctx)
	slog.InfoContext(ctx, fmt.Sprintf("DoSomething got %d", id))
}
