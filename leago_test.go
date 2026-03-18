package leago_test

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago"
	memorylimiter "github.com/Gustavo-Feijo/leago/ratelimit/memory"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/require"
)

func TestNewRegionClient(t *testing.T) {
	client := leago.NewRegionClient(
		regions.RegionAmericas,
		"key",
		leago.WithClient(http.DefaultClient),
		leago.WithLogger(slog.Default()),
		leago.WithLimiter(memorylimiter.NewMemoryLimiter()),
	)
	require.NotNil(t, client)
}

func TestNewPlatformClient(t *testing.T) {
	client := leago.NewPlatformClient(
		regions.PlatformBR1,
		"ApiKey",
		leago.WithClient(http.DefaultClient),
		leago.WithLogger(slog.Default()),
		leago.WithLimiter(memorylimiter.NewMemoryLimiter()),
	)
	require.NotNil(t, client)
}
