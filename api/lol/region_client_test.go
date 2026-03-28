package lol

import (
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/require"
)

func TestNewRegionClient(t *testing.T) {
	bc := internal.NewHTTPClient(string(regions.RegionAmericas), "apiKey")
	client := NewRegionClient(bc)
	require.NotNil(t, client)

	require.NotNil(t, client.Match)
	require.NotNil(t, client.Tournament)
}
