package match

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedMatch = Match{
		Metadata: MatchMetadata{
			DataVersion:  "1",
			MatchID:      "BR1_123",
			Participants: []string{"puuid1"},
		},
		Info: MatchInfo{
			GameCreation: 1234567890,
			GameID:       123,
			GameDatetime: 1234567890,
			GameLength:   1800,
			GameVersion:  "14.1",
			MapID:        1,
			Participants: []MatchParticipant{
				{
					Companion: Companion{
						ContentID: "companion",
						ItemID:    1,
						SkinID:    1,
						Species:   "pet",
					},
					GoldLeft:             0,
					LastRound:            1,
					Level:                1,
					Placement:            1,
					PlayersEliminated:    0,
					Puuid:                "puuid1",
					RiotIDGameName:       "player",
					RiotIDTagline:        "BR1",
					TimeEliminated:       0,
					TotalDamageToPlayers: 0,
					Traits:               []Trait{},
					Units:                []Unit{},
					Win:                  true,
				},
			},
			QueueID:        1100,
			TFTGameType:    "standard",
			TFTSetCoreName: "set",
			TFTSetNumber:   1,
		},
	}

	matchJSON = `{
   "metadata":{
      "data_version":"1",
      "match_id":"BR1_123",
      "participants":[
         "puuid1"
      ]
   },
   "info":{
      "gameCreation":1234567890,
      "gameId":123,
      "game_datetime":1234567890,
      "game_length":1800,
      "game_version":"14.1",
      "mapId":1,
      "participants":[
         {
            "companion":{
               "content_ID":"companion",
               "item_ID":1,
               "skin_ID":1,
               "species":"pet"
            },
            "gold_left":0,
            "last_round":1,
            "level":1,
            "placement":1,
            "players_eliminated":0,
            "puuid":"puuid1",
            "riotIdGameName":"player",
            "riotIdTagline":"BR1",
            "time_eliminated":0,
            "total_damage_to_players":0,
            "traits":[
               
            ],
            "units":[
               
            ],
            "win":true
         }
      ],
      "queue_id":1100,
      "tft_game_":"standard",
      "tft_set_core_name":"set",
      "tft_set_number":1
   }
}`
)

func newTestRegionClient(statusCode int, responseBody string, httpErr error) *RegionClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.RegionAmericas), "apiKey")

	return NewRegionClient(baseClient)
}

func TestGetMatchByID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult Match
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
			responseBody: `{"invalid json,,,,::"}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   matchJSON,
			expectedResult: expectedMatch,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := newTestRegionClient(tt.statusCode, tt.responseBody, tt.httpErr)

			resp, err := rc.GetMatchByID(context.Background(), "BR1_123")

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

			expectedJSON, _ := json.Marshal(tt.expectedResult)
			jsonResp, _ := json.Marshal(resp)

			assert.Equal(t, expectedJSON, jsonResp)
		})
	}
}

func TestGetMatchesByPUUID(t *testing.T) {
	tests := []struct {
		name         string
		statusCode   int
		responseBody string
		wantErr      bool
	}{
		{
			name:         "success",
			statusCode:   http.StatusOK,
			responseBody: `["BR1_1","BR1_2"]`,
			wantErr:      false,
		},
		{
			name:         "invalid json",
			statusCode:   http.StatusOK,
			responseBody: `{"bad"}`,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := newTestRegionClient(tt.statusCode, tt.responseBody, nil)

			start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
			end := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC).Unix()
			resp, err := rc.GetMatchesByPUUID(
				context.Background(),
				"puuid",
				[]GetMatchesByPUUIDOption{
					WithStart(0),
					WithCount(20),
					WithStartTime(start),
					WithEndTime(end),
				},
			)

			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			assert.Equal(t, []string{"BR1_1", "BR1_2"}, resp)
		})
	}
}
