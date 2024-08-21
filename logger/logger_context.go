package logger

import "context"

type contextKey string

const (
	traceIDKey contextKey = "traceID"
)

var extraContextKeys = []contextKey{traceIDKey}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}
