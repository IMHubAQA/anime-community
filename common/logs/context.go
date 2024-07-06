package logs

import (
	"context"

	"github.com/google/uuid"
)

type traceIDKey struct{}

func NewTraceContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	uuid := uuid.New().String()
	ctx = context.WithValue(ctx, traceIDKey{}, uuid)
	return ctx
}
