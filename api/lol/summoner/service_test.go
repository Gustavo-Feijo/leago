package summoner

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
	expectedSummoner = Summoner{
		ProfileIconID: 15,
		RevisionDate:  1772728809000,
		Puuid:         "test-puuid",
		SummonerLevel: 400,
	}

	summonerJSON = `{
   "profileIconId":15,
   "revisionDate":1772728809000,
   "puuid":"test-puuid",
   "summonerLevel":400
}
	`
)

func newTestPlatformClient(statusCode int, responseBody string, httpErr error) *PlatformClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient)
}

func TestGetSummonerByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult Summoner
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
			responseBody:   summonerJSON,
			expectedResult: expectedSummoner,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetSummonerByPUUID(context.Background(), "test-puuid")

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
			assert.Equal(t, expectedSummoner, resp)
		})
	}
}
