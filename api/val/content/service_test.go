package content

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
	expectedContent = Content{
		Version: "1.0",

		Characters: []ContentItem{
			{
				Name:      "Agent",
				ID:        "agent-1",
				AssetName: "agent_asset",
			},
		},

		Maps: []ContentItem{
			{
				Name:      "Ascent",
				ID:        "map-1",
				AssetName: "ascent_asset",
				AssetPath: strPtr("/Game/Maps/Ascent"),
			},
		},

		Chromas:      []ContentItem{},
		Skins:        []ContentItem{},
		SkinLevels:   []ContentItem{},
		Equips:       []ContentItem{},
		GameModes:    []ContentItem{},
		Sprays:       []ContentItem{},
		SprayLevels:  []ContentItem{},
		Charms:       []ContentItem{},
		CharmLevels:  []ContentItem{},
		PlayerCards:  []ContentItem{},
		PlayerTitles: []ContentItem{},

		Acts: []Act{
			{
				Name:     "Episode 1",
				ID:       "act-1",
				IsActive: true,
			},
		},
	}

	contentJSON = `{
		"version": "1.0",
		"characters": [
			{
				"name": "Agent",
				"id": "agent-1",
				"assetName": "agent_asset"
			}
		],
		"maps": [
			{
				"name": "Ascent",
				"id": "map-1",
				"assetName": "ascent_asset",
				"assetPath": "/Game/Maps/Ascent"
			}
		],
		"chromas": [],
		"skins": [],
		"skinLevels": [],
		"equips": [],
		"gameModes": [],
		"sprays": [],
		"sprayLevels": [],
		"charms": [],
		"charmLevels": [],
		"playerCards": [],
		"playerTitles": [],
		"acts": [
			{
				"name": "Episode 1",
				"id": "act-1",
				"isActive": true
			}
		]
	}`
)

func strPtr(s string) *string {
	return &s
}

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.RegionAmericas), "apiKey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient), mockDoer
}

func TestGetContent(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult Content

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
			responseBody: `{"invalid json,,,,::"}`,

			wantErr: true,
		},
		{
			name: "success",

			statusCode:   http.StatusOK,
			responseBody: contentJSON,

			expectedPath: "/val/content/v1/contents",

			expectedResult: expectedContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)

			resp, err := rc.GetContent(context.Background(), []ContentOption{WithLocale("pt-BR")})

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
