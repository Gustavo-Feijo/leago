package item

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
	expectedItems = ItemResponse{
		Type:    "item",
		Version: "16.6.1",
		Basic: BasicItem{
			Name:        "Basic",
			Description: "Basic item",
			Colloq:      "",
			Plaintext:   "",
			Consumed:    false,
			Stacks:      1,
			Depth:       1,
			From:        []string{},
			Into:        []string{},
			InStore:     true,
			Stats:       ItemStats{},
			Tags:        []string{},
			Maps:        map[string]bool{"11": true},
		},
		Data: map[string]Item{
			"1001": {
				Name:        "Boots",
				Description: "Slightly increases Movement Speed",
				Colloq:      "",
				Plaintext:   "Speed",
				Into:        []string{},
				From:        []string{},
				Image: Image{
					Full:   "1001.png",
					Sprite: "item0.png",
					Group:  "item",
					X:      0,
					Y:      0,
					W:      48,
					H:      48,
				},
				Gold: Gold{
					Base:        300,
					Total:       300,
					Sell:        210,
					Purchasable: true,
				},
				Tags:  []string{"Boots"},
				Maps:  map[string]bool{"11": true},
				Stats: map[string]float64{"FlatMovementSpeedMod": 25},
			},
		},
		Groups: []ItemGroup{},
		Tree:   []ItemTree{},
	}

	itemsJSON = `{
		"type": "item",
		"version": "16.6.1",
		"basic": {
			"name": "Basic",
			"rune": {"isrune": false, "tier": 0, "type": ""},
			"gold": {"base": 0, "total": 0, "sell": 0, "purchasable": false},
			"group": "",
			"description": "Basic item",
			"colloq": "",
			"plaintext": "",
			"consumed": false,
			"stacks": 1,
			"depth": 1,
			"consumeOnFull": false,
			"from": [],
			"into": [],
			"specialRecipe": 0,
			"inStore": true,
			"hideFromAll": false,
			"requiredChampion": "",
			"requiredAlly": "",
			"stats": {},
			"tags": [],
			"maps": {"11": true}
		},
		"data": {
			"1001": {
				"name": "Boots",
				"description": "Slightly increases Movement Speed",
				"colloq": "",
				"plaintext": "Speed",
				"into": [],
				"from": [],
				"image": {"full":"1001.png","sprite":"item0.png","group":"item","x":0,"y":0,"w":48,"h":48},
				"gold": {"base":300,"total":300,"sell":210,"purchasable":true},
				"tags": ["Boots"],
				"maps": {"11": true},
				"stats": {"FlatMovementSpeedMod": 25}
			}
		},
		"groups": [],
		"tree": []
	}`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(realms.RealmNA), "nokey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient, "16.6.1", "en_US"), mockDoer
}

func TestGetItems(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult ItemResponse

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
			responseBody: itemsJSON,

			expectedPath: "/cdn/16.6.1/data/en_US/item.json",

			expectedResult: expectedItems,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)

			resp, err := rc.GetItems(context.Background())

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
