package account

import (
	"context"
	"io"
	"leago/internal"
	"leago/internal/mock"
	"leago/regions"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	exampleAccount = Account{
		Puuid:    "test-puuid",
		GameName: "TestPlayer",
		TagLine:  "EUW",
	}

	exampleAccountString = `
	{
		"puuid": "test-puuid",
		"gameName": "TestPlayer",
		"tagLine": "EUW"
	}
	`

	exampleActiveRegion = ActiveRegion{
		Puuid:       "test-puuid",
		Game:        "lol",
		ActiveShard: "euw",
	}

	exampleActiveRegionString = `
	{
		"puuid": "test-puuid",
		"game": "lol",
		"activeShard": "euw"
	}
	`

	exampleActiveShard = ActiveShard{
		Puuid:       "test-puuid",
		Game:        "lol",
		ActiveShard: "euw",
	}

	exampleActiveShardString = `
	{
		"puuid": "test-puuid",
		"game": "lol",
		"activeShard": "euw"
	}
	`
)

func TestGetByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		httpErr        error
		responseBody   string
		expectedResult Account
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			puuid:        "test-puuid",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			puuid:        "test-puuid",
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			puuid:          "test-puuid",
			statusCode:     http.StatusOK,
			responseBody:   exampleAccountString,
			expectedResult: exampleAccount,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := &mock.Doer{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(strings.NewReader(tt.responseBody)),
				},
				Err: tt.httpErr,
			}

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.RegionEurope), "apiKey")
			rc := NewRegionClient(baseClient)

			resp, err := rc.GetByPUUID(context.Background(), tt.puuid)

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetByRiotID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		gameName       string
		tagLine        string
		httpErr        error
		responseBody   string
		expectedResult Account
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			gameName:     "TestPlayer",
			tagLine:      "EUW",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			gameName:     "TestPlayer",
			tagLine:      "EUW",
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			gameName:       "TestPlayer",
			tagLine:        "EUW",
			statusCode:     http.StatusOK,
			responseBody:   exampleAccountString,
			expectedResult: exampleAccount,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := &mock.Doer{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(strings.NewReader(tt.responseBody)),
				},
				Err: tt.httpErr,
			}

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.RegionEurope), "apiKey")
			rc := NewRegionClient(baseClient)

			resp, err := rc.GetByRiotID(context.Background(), tt.gameName, tt.tagLine)

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetActiveRegionByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		game           string
		httpErr        error
		responseBody   string
		expectedResult ActiveRegion
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			puuid:        "test-puuid",
			game:         "lol",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			puuid:        "test-puuid",
			game:         "lol",
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			puuid:          "test-puuid",
			game:           "lol",
			statusCode:     http.StatusOK,
			responseBody:   exampleActiveRegionString,
			expectedResult: exampleActiveRegion,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := &mock.Doer{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(strings.NewReader(tt.responseBody)),
				},
				Err: tt.httpErr,
			}

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.RegionEurope), "apiKey")
			rc := NewRegionClient(baseClient)

			resp, err := rc.GetActiveRegionByPUUID(context.Background(), tt.game, tt.puuid)

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetActiveShardByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		game           string
		httpErr        error
		responseBody   string
		expectedResult ActiveShard
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			puuid:        "test-puuid",
			game:         "lol",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			puuid:        "test-puuid",
			game:         "lol",
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			puuid:          "test-puuid",
			game:           "lol",
			statusCode:     http.StatusOK,
			responseBody:   exampleActiveShardString,
			expectedResult: exampleActiveShard,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := &mock.Doer{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(strings.NewReader(tt.responseBody)),
				},
				Err: tt.httpErr,
			}

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.RegionEurope), "apiKey")
			rc := NewRegionClient(baseClient)

			resp, err := rc.GetActiveShardByPUUID(context.Background(), tt.game, tt.puuid)

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}
