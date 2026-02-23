package lol

import (
	"leago/regions"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPlatformClient(t *testing.T) {
	client := NewPlatformClient(http.DefaultClient, slog.Default(), regions.PlatformBR1, "apiKey")
	require.NotNil(t, client)

	require.NotNil(t, client.Champion)
	require.NotNil(t, client.ChampionMastery)
	require.NotNil(t, client.Clash)
}
