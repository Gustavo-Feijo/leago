package account

import (
	"context"
	"log/slog"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedAccount = Account{
		Puuid:    "test-puuid",
		GameName: "TestPlayer",
		TagLine:  "EUW",
	}

	accountJSON = `
	{
		"puuid": "test-puuid",
		"gameName": "TestPlayer",
		"tagLine": "EUW"
	}
	`

	expectedActiveRegion = ActiveRegion{
		Puuid:       "test-puuid",
		Game:        "lol",
		ActiveShard: "euw",
	}

	activeRegionJSON = `
	{
		"puuid": "test-puuid",
		"game": "lol",
		"activeShard": "euw"
	}
	`

	expectedActiveShard = ActiveShard{
		Puuid:       "test-puuid",
		Game:        "lol",
		ActiveShard: "euw",
	}

	activeShardJSON = `
	{
		"puuid": "test-puuid",
		"game": "lol",
		"activeShard": "euw"
	}
	`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.RegionEurope), "apiKey")
	return NewRegionClient(baseClient), mockDoer
}

func TestGetByPUUID(t *testing.T) {
	tests := []struct {
		name string

		puuid string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult Account

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			puuid: "nonexistentpuuid",

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "invalid json",

			puuid: "invalidjsonpuuid",

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			puuid: "testpuuid",

			statusCode:   http.StatusOK,
			responseBody: accountJSON,

			expectedPath: "/riot/account/v1/accounts/by-puuid/testpuuid",

			expectedResult: expectedAccount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)
			resp, err := rc.GetByPUUID(context.Background(), tt.puuid)

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

func TestGetByRiotID(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, accountJSON)
	resp, err := rc.GetByRiotID(context.Background(), "TestPlayer", "NA1")

	require.NoError(t, err)

	assert.Equal(t, "/riot/account/v1/accounts/by-riot-id/TestPlayer/NA1", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedAccount, resp)
}

func TestGetActiveRegionByPUUID(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, activeRegionJSON)
	resp, err := rc.GetActiveRegionByPUUID(context.Background(), ActiveRegionLOL, "testpuuid")

	require.NoError(t, err)

	assert.Equal(t, "/riot/account/v1/region/by-game/lol/by-puuid/testpuuid", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedActiveRegion, resp)
}

func TestGetActiveShardByPUUID(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, activeShardJSON)
	resp, err := rc.GetActiveShardByPUUID(context.Background(), ActiveShardLOR, "testpuuid")

	require.NoError(t, err)

	assert.Equal(t, "/riot/account/v1/active-shards/by-game/lor/by-puuid/testpuuid", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedActiveShard, resp)
}
