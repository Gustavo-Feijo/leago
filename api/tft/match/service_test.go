package match

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
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

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.RegionAmericas), "apiKey")

	return NewRegionClient(baseClient), mockDoer
}

func TestGetMatchByID(t *testing.T) {
	tests := []struct {
		name string

		matchID string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult Match

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			matchID: "nonexistent",

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "invalid json",

			matchID: "badjsonmatch",

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"}`,

			wantErr: true,
		},
		{
			name: "success",

			matchID: "BR1_123",

			statusCode:   http.StatusOK,
			responseBody: matchJSON,

			expectedPath: "/tft/match/v1/matches/BR1_123",

			expectedResult: expectedMatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)

			resp, err := rc.GetMatchByID(context.Background(), "BR1_123")

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

func TestGetMatchesByPUUID(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC).Unix()

	tests := []struct {
		name string

		puuid string
		opts  []GetMatchesByPUUIDOption

		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult []string
	}{
		{
			name: "no filters",

			puuid: "testpuuidnofilters",

			responseBody: `["BR1_1","BR1_2"]`,

			expectedPath: "/tft/match/v1/matches/by-puuid/testpuuidnofilters/ids",

			expectedResult: []string{"BR1_1", "BR1_2"},
		},
		{
			name: "all filters",

			puuid: "testpuuidfiltered",
			opts: []GetMatchesByPUUIDOption{
				WithStart(0),
				WithCount(20),
				WithStartTime(start),
				WithEndTime(end),
			},

			responseBody: `["BR1_1"]`,

			expectedPath: "/tft/match/v1/matches/by-puuid/testpuuidfiltered/ids",
			expectedQuery: map[string]string{
				"start":     "0",
				"count":     "20",
				"startTime": strconv.FormatInt(start, 10),
				"endTime":   strconv.FormatInt(end, 10),
			},

			expectedResult: []string{"BR1_1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(http.StatusOK, tt.responseBody)

			resp, err := rc.GetMatchesByPUUID(
				context.Background(),
				tt.puuid,
				tt.opts,
			)

			require.NoError(t, err)

			assert.Equal(t, tt.expectedPath, mockDoer.CapturedReq.URL.Path)

			query := mockDoer.CapturedReq.URL.Query()
			for k, v := range tt.expectedQuery {
				assert.Equal(t, v, query.Get(k))
			}

			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}
