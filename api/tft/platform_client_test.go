package tft

import (
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/require"
)

func TestNewPlatformClient(t *testing.T) {
	bc := internal.NewHTTPClient(string(regions.PlatformBR1), "apiKey")
	client := NewPlatformClient(bc)
	require.NotNil(t, client)

	require.NotNil(t, client.League)
}
