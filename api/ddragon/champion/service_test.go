package champion

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
	expectedChampions = ChampionResponse{
		Type:    "champion",
		Format:  "standAloneComplex",
		Version: "16.6.1",
		Data: map[string]ChampionData{
			"Aatrox": {
				Version: "16.6.1",
				ID:      "Aatrox",
				Key:     "266",
				Name:    "Aatrox",
				Title:   "the Darkin Blade",
				Tags:    []string{"Fighter", "Tank"},
			},
		},
	}

	championsJSON = `{
		"type": "champion",
		"format": "standAloneComplex",
		"version": "16.6.1",
		"data": {
			"Aatrox": {
				"version": "16.6.1",
				"id": "Aatrox",
				"key": "266",
				"name": "Aatrox",
				"title": "the Darkin Blade",
				"blurb": "",
				"info": {"attack":0,"defense":0,"magic":0,"difficulty":0},
				"image": {"full":"","sprite":"","group":"","x":0,"y":0,"w":0,"h":0},
				"tags": ["Fighter","Tank"],
				"partype": "",
				"stats": {"hp":0,"hpperlevel":0,"mp":0,"mpperlevel":0,"movespeed":0,"armor":0,"armorperlevel":0,"spellblock":0,"spellblockperlevel":0,"attackrange":0,"hpregen":0,"hpregenperlevel":0,"mpregen":0,"mpregenperlevel":0,"crit":0,"critperlevel":0,"attackdamage":0,"attackdamageperlevel":0,"attackspeedperlevel":0,"attackspeed":0}
			}
		}
	}`

	expectedSingleChampion = SingleChampionResponse{
		Type:    "champion",
		Format:  "standAloneComplex",
		Version: "16.6.1",
		Data: map[string]FullChampionData{
			"Aatrox": {
				ID:          "Aatrox",
				Key:         "266",
				Name:        "Aatrox",
				Title:       "the Darkin Blade",
				AllyTips:    []string{},
				EnemyTips:   []string{},
				Tags:        []string{},
				Skins:       []Skin{},
				Spells:      []Spell{},
				Recommended: []any{},
			},
		},
	}

	singleChampionJSON = `{
		"type": "champion",
		"format": "standAloneComplex",
		"version": "16.6.1",
		"data": {
			"Aatrox": {
				"id": "Aatrox",
				"key": "266",
				"name": "Aatrox",
				"title": "the Darkin Blade",
				"lore": "",
				"blurb": "",
				"allytips": [],
				"enemytips": [],
				"tags": [],
				"partype": "",
				"info": {"attack":0,"defense":0,"magic":0,"difficulty":0},
				"stats": {"hp":0,"hpperlevel":0,"mp":0,"mpperlevel":0,"movespeed":0,"armor":0,"armorperlevel":0,"spellblock":0,"spellblockperlevel":0,"attackrange":0,"hpregen":0,"hpregenperlevel":0,"mpregen":0,"mpregenperlevel":0,"crit":0,"critperlevel":0,"attackdamage":0,"attackdamageperlevel":0,"attackspeedperlevel":0,"attackspeed":0},
				"image": {"full":"","sprite":"","group":"","x":0,"y":0,"w":0,"h":0},
				"skins": [],
				"spells": [],
				"passive": {"name":"","description":"","image":{"full":"","sprite":"","group":"","x":0,"y":0,"w":0,"h":0}},
				"recommended": []
			}
		}
	}`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(realms.RealmNA), "nokey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient, "16.6.1", "en_US"), mockDoer
}

func TestGetChampions(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult ChampionResponse

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
			responseBody: championsJSON,

			expectedPath: "/cdn/16.6.1/data/en_US/champion.json",

			expectedResult: expectedChampions,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)

			resp, err := rc.GetChampions(context.Background())

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

func TestGetChampionByID(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, singleChampionJSON)

	resp, err := rc.GetChampionByID(context.Background(), "Aatrox")

	require.NoError(t, err)

	assert.Equal(t, "/cdn/16.6.1/data/en_US/champion/Aatrox.json", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedSingleChampion, resp)
}
