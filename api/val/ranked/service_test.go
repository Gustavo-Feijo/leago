package ranked

import (
	"context"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func strPtr(s string) *string {
	return &s
}

var (
	expectedLeaderboard = Leaderboard{
		Shard:        "na",
		ActID:        "test-act",
		TotalPlayers: 1,
		Players: []Player{
			{
				PUUID:           strPtr("test-puuid"),
				GameName:        strPtr("TestUser"),
				TagLine:         strPtr("NA1"),
				LeaderboardRank: 1,
				RankedRating:    100,
				NumberOfWins:    10,
			},
		},
	}
	leaderboardJSON = `
	{
  "shard": "na",
  "actId": "test-act",
  "totalPlayers": 1,
  "players": [
    {
      "puuid": "test-puuid",
      "gameName": "TestUser",
      "tagLine": "NA1",
      "leaderboardRank": 1,
      "rankedRating": 100,
      "numberOfWins": 10
    }
  ]
}	`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.ValRegionBR), "apiKey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient), mockDoer
}

func TestGetLeaderboard(t *testing.T) {
	tests := []struct {
		name string

		actID string
		opts  []RankedOption

		statusCode   int
		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult Leaderboard

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
			responseBody: `{"invalid json}`,

			wantErr: true,
		},
		{
			name: "success no filter",

			actID: "act1",

			statusCode:   http.StatusOK,
			responseBody: leaderboardJSON,

			expectedPath: "/val/ranked/v1/leaderboards/by-act/act1",

			expectedResult: expectedLeaderboard,
		},
		{
			name: "success with filter",

			actID: "act2",
			opts: []RankedOption{
				WithSize(200),
				WithStartIndex(0),
			},

			statusCode:   http.StatusOK,
			responseBody: leaderboardJSON,

			expectedPath: "/val/ranked/v1/leaderboards/by-act/act2",
			expectedQuery: map[string]string{
				"size":       "200",
				"startIndex": "0",
			},

			expectedResult: expectedLeaderboard,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)
			resp, err := rc.GetLeaderboard(
				context.Background(),
				tt.actID,
				tt.opts,
			)

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

			query := mockDoer.CapturedReq.URL.Query()
			for k, v := range tt.expectedQuery {
				assert.Equal(t, v, query.Get(k))
			}

			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}
