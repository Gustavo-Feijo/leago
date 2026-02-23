package champion

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
	exampleRotation = Rotation{
		MaxNewPlayerLevel:            10,
		FreeChampionIdsForNewPlayers: []int{1, 2, 3},
		FreeChampionIds:              []int{4, 5, 6},
	}

	exampleRotationString = `
	{
		"maxNewPlayerLevel": 10,
		"freeChampionIdsForNewPlayers": [1, 2, 3],
		"freeChampionIds": [4, 5, 6]
	}
	`
)

func TestGetRotation(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult Rotation
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
			responseBody:   exampleRotationString,
			expectedResult: exampleRotation,
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

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)

			resp, err := pc.GetRotation(context.Background())

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
