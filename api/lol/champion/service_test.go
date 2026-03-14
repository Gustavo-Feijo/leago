package champion

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

var (
	expectedRotation = Rotation{
		MaxNewPlayerLevel:            10,
		FreeChampionIDsForNewPlayers: []int{1, 2, 3},
		FreeChampionIDs:              []int{4, 5, 6},
	}

	rotationJSON = `
	{
		"maxNewPlayerLevel": 10,
		"freeChampionIdsForNewPlayers": [1, 2, 3],
		"freeChampionIds": [4, 5, 6]
	}
	`
)

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.PlatformBR1), "apiKey", internal.WithHTTP(mockDoer))

	return NewPlatformClient(baseClient), mockDoer
}

func TestGetRotation(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedResult Rotation

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
			responseBody: rotationJSON,

			expectedResult: expectedRotation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetRotation(context.Background())

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

			assert.Equal(t, "/lol/platform/v3/champion-rotations", mockDoer.CapturedReq.URL.Path)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}
