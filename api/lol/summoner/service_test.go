package summoner

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedSummoner = Summoner{
		ProfileIconID: 15,
		RevisionDate:  internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
		Puuid:         "test-puuid",
		SummonerLevel: 400,
	}

	summonerJSON = `{
   "profileIconId":15,
   "revisionDate":1767225600000,
   "puuid":"test-puuid",
   "summonerLevel":400
}
	`
)

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.PlatformBR1), "apiKey", internal.WithHTTP(mockDoer))

	return NewPlatformClient(baseClient), mockDoer
}

func TestGetSummonerByPUUID(t *testing.T) {
	tests := []struct {
		name string

		puuid string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult Summoner

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

			puuid: "testpuuidbadjson",

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			puuid: "testpuuid",

			statusCode:   http.StatusOK,
			responseBody: summonerJSON,

			expectedPath: "/lol/summoner/v4/summoners/by-puuid/testpuuid",

			expectedResult: expectedSummoner,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetSummonerByPUUID(context.Background(), tt.puuid)

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
