package ddragon

import (
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/api/ddragon/realms"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	realmJSON = `{
		"n":{
			"item":"16.6.1",
			"rune":"7.23.1",
			"mastery":"7.23.1",
			"summoner":"16.6.1",
			"champion":"16.6.1",
			"profileicon":"16.6.1",
			"map":"16.6.1",
			"language":"16.6.1",
			"sticker":"16.6.1"
		},
		"v":"99.1.1",
		"l":"ko_KR",
		"cdn":"https://ddragon.leagueoflegends.com/cdn",
		"dd":"99.1.1",
		"lg":"99.1.1",
		"css":"99.1.1",
		"profileiconmax":28,
		"store":null
	}`
)

func newTestRegionClient(statusCode int, responseBody string) (*internal.Client, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	client := internal.NewHTTPClient(string(realms.RealmNA), "nokey", internal.WithHTTP(mockDoer))
	return client, mockDoer
}

func TestNewRegionClient(t *testing.T) {
	bc, _ := newTestRegionClient(http.StatusOK, realmJSON)

	client, _ := NewRegionClient(bc, realms.RealmBR, "16.6.1", "pt_BR")

	require.NotNil(t, client)

	assert.Equal(t, "16.6.1", client.version)
	assert.Equal(t, "pt_BR", client.locale)

	require.NotNil(t, client.Champion)
	require.NotNil(t, client.Item)
	require.NotNil(t, client.Language)
	require.NotNil(t, client.ProfileIcon)
	require.NotNil(t, client.Realms)
	require.NotNil(t, client.Summoner)
	require.NotNil(t, client.Version)
}

func TestNewRegionClientWithBootstrap(t *testing.T) {
	bc, mockDoer := newTestRegionClient(http.StatusOK, realmJSON)

	client, _ := NewRegionClient(bc, realms.RealmKR, "", "")

	require.NotNil(t, client)

	assert.Equal(t, "99.1.1", client.version)
	assert.Equal(t, "ko_KR", client.locale)

	assert.Equal(t, "/realms/kr.json", mockDoer.CapturedReq.URL.Path)

	require.NotNil(t, client.Champion)
	require.NotNil(t, client.Item)
	require.NotNil(t, client.Language)
	require.NotNil(t, client.ProfileIcon)
	require.NotNil(t, client.Realms)
	require.NotNil(t, client.Summoner)
	require.NotNil(t, client.Version)
}

func TestNewRegionClientBootstrapErrorFallback(t *testing.T) {
	bc, _ := newTestRegionClient(http.StatusNotFound, `{"status":{"status_code":404}}`)

	client, err := NewRegionClient(bc, realms.RealmKR, "", "")

	require.Error(t, err)
	require.Nil(t, client)
}

func TestNewRegionClientManualOverride(t *testing.T) {
	bc, _ := newTestRegionClient(http.StatusOK, realmJSON)

	client, _ := NewRegionClient(bc, realms.RealmKR, "1.2.3", "fr_FR")

	require.NotNil(t, client)

	assert.Equal(t, "1.2.3", client.version)
	assert.Equal(t, "fr_FR", client.locale)
}

func TestBootstrap(t *testing.T) {
	bc, _ := newTestRegionClient(http.StatusOK, realmJSON)

	client := &RegionClient{
		version: "old",
		locale:  "old",
		Realms:  realms.NewRegionClient(bc),
	}

	err := client.bootstrap(realms.RealmKR)
	require.NoError(t, err)

	assert.Equal(t, "99.1.1", client.version)
	assert.Equal(t, "ko_KR", client.locale)
}

func TestBootstrapError(t *testing.T) {
	bc, _ := newTestRegionClient(http.StatusNotFound, `{"status":{"status_code":404}}`)

	client := &RegionClient{
		version: "old",
		locale:  "old",
		Realms:  realms.NewRegionClient(bc),
	}

	err := client.bootstrap(realms.RealmKR)
	require.Error(t, err)

	assert.Equal(t, "old", client.version)
	assert.Equal(t, "old", client.locale)
}
