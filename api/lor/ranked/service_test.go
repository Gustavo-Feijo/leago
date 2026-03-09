package ranked

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
	expectedLeaderboard = Leaderboard{
		Players: []Player{
			{
				Name: "test",
				Rank: 1,
				LP:   100.0,
			},
		},
	}

	leaderboardJSON = `{
   "players":[
      {
         "name":"test",
         "rank":1,
         "lp":100.0
      }
   ]
}`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.RegionAmericas), "apiKey")
	return NewRegionClient(baseClient), mockDoer
}

func TestGetLeaderboard(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedPath string

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
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			statusCode:   http.StatusOK,
			responseBody: leaderboardJSON,

			expectedPath: "/lor/ranked/v1/leaderboards",

			expectedResult: expectedLeaderboard,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)
			resp, err := rc.GetLeaderboards(context.Background())

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
