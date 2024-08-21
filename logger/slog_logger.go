package logger

import (
	"context"
	"log/slog"
	"os"
)

type customHandler struct {
	handler slog.Handler
}

var _ slog.Handler = &customHandler{}

func init() {
	logger := slog.New(
		&customHandler{
			slog.NewJSONHandler(
				os.Stdout, nil,
			),
		},
	)

	slog.SetDefault(logger)
}

func (h *customHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, key := range extraContextKeys {
		if value := ctx.Value(key); value != nil {
			record.AddAttrs(
				slog.Attr{
					Key:   string(key),
					Value: slog.AnyValue(value),
				},
			)
		}
	}
	return h.handler.Handle(ctx, record)
}

func (h *customHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &customHandler{
		handler: h.handler.WithAttrs(attrs),
	}
}

func (h *customHandler) WithGroup(name string) slog.Handler {
	return &customHandler{
		handler: h.handler.WithGroup(name),
	}
}
