package noop

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoopLimiter(t *testing.T) {
	limiter := NewNoopLimiter()
	require.NotNil(t, limiter)
	assert.Nil(t, limiter.Acquire(context.Background(), "key", "methodKey"))
}
