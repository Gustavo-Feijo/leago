package leago_test

import (
	"leago"
	"leago/regions"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRegionClient(t *testing.T) {
	client := leago.NewRegionClient(
		regions.RegionAmericas,
		"key",
		leago.WithClient(http.DefaultClient),
	)
	require.NotNil(t, client)
}

func TestNewPlatformClient(t *testing.T) {
	client := leago.NewPlatformClient(
		regions.RegionBR1,
		"key",
		leago.WithClient(http.DefaultClient),
	)
	require.NotNil(t, client)
}
