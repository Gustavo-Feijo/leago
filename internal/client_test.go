package internal

import (
	"testing"

	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHTTPClient(t *testing.T) {
	client := NewHTTPClient(string(regions.PlatformBR1), "apiKey")
	require.NotNil(t, client)
	assert.Equal(t, "apiKey", client.apiKey)
}

func TestGetURL(t *testing.T) {
	client := NewHTTPClient(string(regions.PlatformBR1), "apiKey")
	url := client.GetURL("/testapi")
	assert.Contains(t, url, string(regions.PlatformBR1))
}
