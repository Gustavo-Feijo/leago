package ranked

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestRegionClient(statusCode int, responseBody string, httpErr error) *RegionClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.RegionAmericas), "apiKey")
	return NewRegionClient(baseClient)
}

var (
	leaderboardJSON = `{
   "players":[
      {
         "name":"test",
         "rank":1,
         "lp":100.0
      }
   ]
}`

	expectedLeaderboard = Leaderboard{
		Players: []Player{
			{
				Name: "test",
				Rank: 1,
				LP:   100.0,
			},
		},
	}
)

func TestGetLeaderboard(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult Leaderboard
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   leaderboardJSON,
			expectedResult: expectedLeaderboard,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := newTestRegionClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := rc.GetLeaderboards(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err)
				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}
				return
			}

			require.Nil(t, err)
			require.NotNil(t, resp)

			// Marshal both to not run into timezone problems.
			expectedJson, _ := json.Marshal(tt.expectedResult)
			jsonResp, _ := json.Marshal(resp)

			assert.Equal(t, expectedJson, jsonResp)
		})
	}
}
