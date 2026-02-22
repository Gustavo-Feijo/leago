package leago_test

import (
	"leago"
	"leago/regions"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRegionClient(t *testing.T) {
	client := leago.NewRegionClient(
		regions.RegionAmericas,
		"key",
		leago.WithClient(http.DefaultClient),
		leago.WithLogger(slog.Default()),
	)
	require.NotNil(t, client)
}

func TestNewPlatformClient(t *testing.T) {
	client := leago.NewPlatformClient(
		regions.PlatformBR1,
		"ApiKey",
		leago.WithClient(http.DefaultClient),
		leago.WithLogger(slog.Default()),
	)
	require.NotNil(t, client)
}
