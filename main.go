package main

import (
	"context"
	"slog-sample/domain/service"
	"slog-sample/infra/repository"
	"slog-sample/logger"
	"slog-sample/presentation/controller"
	"sync"

	"github.com/google/uuid"
)

func main() {
	c := controller.NewController(
		service.NewService(
			repository.NewRepository(),
		),
	)

	var wg sync.WaitGroup

	concurrentNum := 10
	for i := 0; i < concurrentNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			traceID := uuid.New().String()
			ctx := logger.WithTraceID(context.Background(), traceID)
			c.Handle(ctx)
		}()
	}
	wg.Wait()
}
