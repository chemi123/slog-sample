package controller

import (
	"context"
	"log/slog"
	"slog-sample/domain/service"
)

type Controller interface {
	Handle(ctx context.Context)
}

type controller struct {
	service service.Service
}

var _ Controller = &controller{}

func NewController(service service.Service) Controller {
	return &controller{
		service,
	}
}

func (c *controller) Handle(ctx context.Context) {
	slog.InfoContext(ctx, "Handle")
	c.service.DoSomething(ctx)
}
