package spectator

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedCurrentGame = CurrentGameInfo{
		GameID:            123456789,
		GameType:          "MATCHED_GAME",
		GameStartTime:     1700000000000,
		MapID:             11,
		GameLength:        600,
		PlatformID:        "BR1",
		GameMode:          "CLASSIC",
		GameQueueConfigID: 420,
		BannedChampions: []BannedChampion{
			{
				PickTurn:   1,
				ChampionID: 157,
				TeamID:     100,
			},
		},
		Observers: Observer{
			EncryptionKey: "test-encryption-key",
		},
		Participants: []CurrentGameParticipant{
			{
				ChampionID:    157,
				ProfileIconID: 1234,
				Bot:           false,
				TeamID:        100,
				PUUID:         new("test-puuid"),
				Spell1ID:      4,
				Spell2ID:      14,
				Perks: Perks{
					PerkIDs:      []int64{8005, 9111},
					PerkStyle:    8000,
					PerkSubStyle: 8100,
				},
				GameCustomizationObjects: []GameCustomizationObject{
					{
						Category: "championSkin",
						Content:  "skin-id-1",
					},
				},
			},
		},
	}

	currentGameJSON = `{
		"gameId": 123456789,
		"gameType": "MATCHED_GAME",
		"gameStartTime": 1700000000000,
		"mapId": 11,
		"gameLength": 600,
		"platformId": "BR1",
		"gameMode": "CLASSIC",
		"gameQueueConfigId": 420,
		"bannedChampions": [
			{
				"pickTurn": 1,
				"championId": 157,
				"teamId": 100
			}
		],
		"observers": {
			"encryptionKey": "test-encryption-key"
		},
		"participants": [
			{
				"championId": 157,
				"perks": {
					"perkIds": [8005, 9111],
					"perkStyle": 8000,
					"perkSubStyle": 8100
				},
				"profileIconId": 1234,
				"bot": false,
				"teamId": 100,
				"puuid": "test-puuid",
				"spell1Id": 4,
				"spell2Id": 14,
				"gameCustomizationObjects": [
					{
						"category": "championSkin",
						"content": "skin-id-1"
					}
				]
			}
		]
	}`
)

func newTestPlatformClient(statusCode int, responseBody string, httpErr error) *PlatformClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient)
}

func TestGetGameByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult CurrentGameInfo
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
			responseBody:   currentGameJSON,
			expectedResult: expectedCurrentGame,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetGameByPUUID(context.Background(), "test-puuid")

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

			// Normalize through JSON to avoid pointer comparison issues.
			expectedJSON, _ := json.Marshal(tt.expectedResult)
			respJSON, _ := json.Marshal(resp)

			assert.Equal(t, expectedJSON, respJSON)
		})
	}
}
