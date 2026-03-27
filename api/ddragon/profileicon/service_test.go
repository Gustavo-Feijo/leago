package profileicon

import (
	"context"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/api/ddragon/realms"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedProfileIcons = ProfileIconResponse{
		Type:    "profileicon",
		Version: "16.6.1",
		Data: map[string]ProfileIcon{
			"0": {
				ID: 0,
				Image: Image{
					Full:   "0.png",
					Sprite: "profileicon0.png",
					Group:  "profileicon",
					X:      0,
					Y:      0,
					W:      48,
					H:      48,
				},
			},
			"50": {
				ID: 50,
				Image: Image{
					Full:   "50.png",
					Sprite: "profileicon0.png",
					Group:  "profileicon",
					X:      48,
					Y:      0,
					W:      48,
					H:      48,
				},
			},
		},
	}

	profileIconsJSON = `{
		"type":"profileicon",
		"version":"16.6.1",
		"data":{
			"0":{
				"id":0,
				"image":{"full":"0.png","sprite":"profileicon0.png","group":"profileicon","x":0,"y":0,"w":48,"h":48}
			},
			"50":{
				"id":"50",
				"image":{"full":"50.png","sprite":"profileicon0.png","group":"profileicon","x":48,"y":0,"w":48,"h":48}
			}
		}
	}`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(realms.RealmNA), "nokey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient, "16.6.1", "en_US"), mockDoer
}

func TestGetProfileIcons(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult ProfileIconResponse

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
			responseBody: `{"invalid json`,

			wantErr: true,
		},
		{
			name: "success",

			statusCode:   http.StatusOK,
			responseBody: profileIconsJSON,

			expectedPath: "/cdn/16.6.1/data/en_US/profileicon.json",

			expectedResult: expectedProfileIcons,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)

			resp, err := rc.GetProfileIcons(context.Background())

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
