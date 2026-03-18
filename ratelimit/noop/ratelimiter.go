package noop

import (
	"context"

	"github.com/Gustavo-Feijo/leago/ratelimit"
)

func NewNoopLimiter() ratelimit.RateLimiter {
	return &noopLimiter{}
}

type noopLimiter struct{}

func (l *noopLimiter) Acquire(ctx context.Context, appKey, methodKey string) error {
	return nil
}
