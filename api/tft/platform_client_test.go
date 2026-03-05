package tft

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/require"
)

func TestNewPlatformClient(t *testing.T) {
	client := NewPlatformClient(http.DefaultClient, slog.Default(), regions.PlatformBR1, "apiKey")
	require.NotNil(t, client)

	require.NotNil(t, client.League)
}
