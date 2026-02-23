package riot

import (
	"leago/regions"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPlatformClient(t *testing.T) {
	client := NewRegionClient(http.DefaultClient, slog.Default(), regions.RegionAmericas, "apiKey")
	require.NotNil(t, client)

	require.NotNil(t, client.Account)
}
