package mock

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockLimiter struct {
	mock.Mock
}

func (m *MockLimiter) Acquire(ctx context.Context, appKey, methodKey string) error {
	args := m.Called(ctx, appKey, methodKey)
	return args.Error(0)
}

func (m *MockLimiter) Sync(ctx context.Context, headers http.Header, appKey, methodKey string) {
	m.Called(ctx, headers, appKey, methodKey)
}

func (m *MockLimiter) NotifyTooManyRequests(ctx context.Context, headers http.Header, appKey, methodKey string) {
	m.Called(ctx, headers, appKey, methodKey)
}
