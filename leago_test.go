package leago_test

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago"
	"github.com/Gustavo-Feijo/leago/api/ddragon/realms"
	"github.com/Gustavo-Feijo/leago/internal/mock"
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
	require.NotNil(t, client.Lol)
	require.NotNil(t, client.Lor)
	require.NotNil(t, client.Riot)
	require.NotNil(t, client.Tft)
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
	require.NotNil(t, client.Lol)
	require.NotNil(t, client.Tft)
}

func TestNewValRegionClient(t *testing.T) {
	client := leago.NewValRegionClient(
		regions.ValRegionBR,
		"ApiKey",
		leago.WithClient(http.DefaultClient),
		leago.WithLogger(slog.Default()),
		leago.WithLimiter(memorylimiter.NewMemoryLimiter()),
	)
	require.NotNil(t, client)
	require.NotNil(t, client.Val)
}

func TestNewDDragonClient(t *testing.T) {
	mockDoer := mock.NewDefaultDoer(http.StatusOK, "okay")
	client, err := leago.NewDDragonClient(
		realms.RealmBR,
		[]leago.DDragonOption{
			leago.WithLocale("en_US"),
			leago.WithVersion("16.6.1"),
		},
		leago.WithClient(mockDoer),
		leago.WithLogger(slog.Default()),
		leago.WithLimiter(memorylimiter.NewMemoryLimiter()),
	)
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestNewDDragonClientError(t *testing.T) {
	mockDoer := mock.NewDefaultDoer(http.StatusBadRequest, "badrequest")
	client, err := leago.NewDDragonClient(
		realms.RealmBR,
		[]leago.DDragonOption{
			// No locale and version, must trigger the request.
		},
		leago.WithClient(mockDoer),
		leago.WithLogger(slog.Default()),
		leago.WithLimiter(memorylimiter.NewMemoryLimiter()),
	)

	require.Nil(t, client)
	require.Error(t, err)
}
