package summoner

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
	emptyStr = ""

	expectedSummoners = SummonerResponse{
		Type:    "summoner",
		Version: "16.6.1",
		Data: map[string]Summoner{
			"SummonerFlash": {
				ID:            "SummonerFlash",
				Name:          "Flash",
				Description:   "",
				Tooltip:       "",
				MaxRank:       1,
				Cooldown:      []float64{300},
				CooldownBurn:  "300",
				Cost:          []int{0},
				CostBurn:      "0",
				DataValues:    map[string]any{},
				Effect:        [][]float64{},
				EffectBurn:    []*string{&emptyStr},
				Vars:          []any{},
				Key:           "4",
				SummonerLevel: 1,
				Modes:         []string{"CLASSIC"},
				CostType:      "",
				MaxAmmo:       "",
				Range:         []int{425},
				RangeBurn:     "425",
				Image: Image{
					Full:   "SummonerFlash.png",
					Sprite: "spell0.png",
					Group:  "spell",
					X:      0,
					Y:      0,
					W:      48,
					H:      48,
				},
				Resource: "",
			},
		},
	}

	summonersJSON = `{
		"type":"summoner",
		"version":"16.6.1",
		"data":{
			"SummonerFlash":{
				"id":"SummonerFlash",
				"name":"Flash",
				"description":"",
				"tooltip":"",
				"maxrank":1,
				"cooldown":[300],
				"cooldownBurn":"300",
				"cost":[0],
				"costBurn":"0",
				"datavalues":{},
				"effect":[],
				"effectBurn":[""],
				"vars":[],
				"key":"4",
				"summonerLevel":1,
				"modes":["CLASSIC"],
				"costType":"",
				"maxammo":"",
				"range":[425],
				"rangeBurn":"425",
				"image":{"full":"SummonerFlash.png","sprite":"spell0.png","group":"spell","x":0,"y":0,"w":48,"h":48},
				"resource":""
			}
		}
	}`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(realms.RealmNA), "nokey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient, "16.6.1", "en_US"), mockDoer
}

func TestGetSummonerSpells(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult SummonerResponse

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
			responseBody: summonersJSON,

			expectedPath: "/cdn/16.6.1/data/en_US/summoner.json",

			expectedResult: expectedSummoners,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)

			resp, err := rc.GetSummonerSpells(context.Background())

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
