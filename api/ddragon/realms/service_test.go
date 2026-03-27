package realms

import (
	"context"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	storeURL = "https://store.riotgames.com"

	expectedRealm = DDragonResponse{
		N: DDragonVersions{
			Item:        "16.6.1",
			Rune:        "7.23.1",
			Mastery:     "7.23.1",
			Summoner:    "16.6.1",
			Champion:    "16.6.1",
			ProfileIcon: "16.6.1",
			Map:         "16.6.1",
			Language:    "16.6.1",
			Sticker:     "16.6.1",
		},
		V:              "16.6.1",
		L:              "en_US",
		CDN:            "https://ddragon.leagueoflegends.com/cdn",
		DD:             "16.6.1",
		LG:             "16.6.1",
		CSS:            "16.6.1",
		ProfileIconMax: 28,
		Store:          &storeURL,
	}

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
		"v":"16.6.1",
		"l":"en_US",
		"cdn":"https://ddragon.leagueoflegends.com/cdn",
		"dd":"16.6.1",
		"lg":"16.6.1",
		"css":"16.6.1",
		"profileiconmax":28,
		"store":"https://store.riotgames.com"
	}`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient("na1", "nokey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient), mockDoer
}

func TestGetRealm(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult DDragonResponse

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "invalid json",

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json`,

			wantErr: true,
		},
		{
			name: "success",

			statusCode:   http.StatusOK,
			responseBody: realmJSON,

			expectedPath: "/realms/na.json",

			expectedResult: expectedRealm,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)

			resp, err := rc.GetRealm(context.Background(), RealmNA)

			if tt.wantErr {
				require.Error(t, err)
				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}
				return
			}

			require.NoError(t, err)

			assert.Equal(t, tt.expectedPath, mockDoer.CapturedReq.URL.Path)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}
