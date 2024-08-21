package repository

import (
	"context"
	"log/slog"
	"slog-sample/domain/infra_interface"
)

type repository struct {
}

var _ infra_interface.Repository = &repository{}

func NewRepository() infra_interface.Repository {
	return &repository{}
}

func (r *repository) Get(ctx context.Context) int {
	slog.InfoContext(ctx, "Get")
	return 10
}
