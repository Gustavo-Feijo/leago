package clash

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

func TestGetByPUUID(t *testing.T) {

	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		httpErr        error
		responseBody   string
		expectedResult PlayersResponse
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
			name:         "unmatched json",
			puuid:        "test-puuid",
			statusCode:   http.StatusOK,
			responseBody: `{"puuid":"shouldbearray"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:       "success",
			puuid:      "test-puuid",
			statusCode: http.StatusOK,
			responseBody: `[{
				"puuid":"test-puuid",
				"teamId":"456",
				"position":"TOP",
				"role":"CAPTAIN"
			}]`,
			expectedResult: PlayersResponse{
				{
					Puuid:    "test-puuid",
					TeamID:   "456",
					Position: "TOP",
					Role:     "CAPTAIN",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := &mock.MockDoer{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(strings.NewReader(tt.responseBody)),
				},
				Err: tt.httpErr,
			}

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetByPUUID(context.Background(), tt.puuid)

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
