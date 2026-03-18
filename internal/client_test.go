package internal

import (
	"log/slog"
	"testing"

	"github.com/Gustavo-Feijo/leago/internal/mock"
	memorylimiter "github.com/Gustavo-Feijo/leago/ratelimit/memory"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHTTPClient(t *testing.T) {
	client := NewHTTPClient(string(regions.PlatformBR1), "apiKey")
	require.NotNil(t, client)
	assert.Equal(t, "apiKey", client.apiKey)
}

func TestNewHTTPClientWithParams(t *testing.T) {
	mockDoer := &mock.Doer{}
	limiter := memorylimiter.NewMemoryLimiter()
	logger := slog.New(slog.DiscardHandler)
	client := NewHTTPClient(string(regions.PlatformBR1), "apiKey",
		WithHTTP(mockDoer),
		WithLimiter(limiter),
		WithLogger(logger),
	)
	require.NotNil(t, client)
	assert.Equal(t, mockDoer, client.HTTP)
	assert.Equal(t, limiter, client.limiter)
	assert.Equal(t, logger, client.Logger)
}
func TestGetURL(t *testing.T) {
	client := NewHTTPClient(string(regions.PlatformBR1), "apiKey")
	url := client.GetURL("/testapi")
	assert.Contains(t, url, string(regions.PlatformBR1))
}
