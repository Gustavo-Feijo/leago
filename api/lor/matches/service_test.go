package matches

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
	expectedMatchList = []string{
		"matchTest1",
		"matchTest2",
	}

	matchListJSON = `["matchTest1","matchTest2"]`

	expectedMatch = Match{
		Metadata: Metadata{
			DataVersion:  "2",
			MatchID:      "testID",
			Participants: []string{"testpuuid"},
		},
		Info: Info{
			GameMode:         GameModeThePathOfChampions,
			GameType:         GameType(""),
			GameStartTimeUTC: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			GameVersion:      "live-green-6-07-3",
			GameFormat:       GameFormatStandard,
			Players: []Player{
				{
					PUUID:    "testpuuid",
					DeckID:   "",
					DeckCode: "",
					Factions: []string{
						"faction_BandleCity_Name",
						"faction_Bard_Name",
						"faction_Demacia_Name",
						"faction_Freljord_Name",
						"faction_MtTargon_Name",
					},
					GameOutcome: "win",
					OrderOfPlay: 1,
				},
			},
			TotalTurnCount: 19,
		},
	}

	matchJSON = `{
    "metadata": {
        "data_version": "2",
        "match_id": "testID",
        "participants": [
            "testpuuid"
        ]
    },
    "info": {
        "game_mode": "ThePathOfChampions",
        "game_type": "",
        "game_start_time_utc": "2026-01-01T00:00:00.0000000+00:00",
        "game_version": "live-green-6-07-3",
        "game_format": "standard",
        "players": [
            {
                "puuid": "testpuuid",
                "deck_id": "",
                "deck_code": "",
                "factions": [
                    "faction_BandleCity_Name",
                    "faction_Bard_Name",
                    "faction_Demacia_Name",
                    "faction_Freljord_Name",
                    "faction_MtTargon_Name"
                ],
                "game_outcome": "win",
                "order_of_play": 1
            }
        ],
        "total_turn_count": 19
    }
}`
)

func newTestRegionClient(statusCode int, responseBody string, httpErr error) *RegionClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.RegionAmericas), "apiKey")
	return NewRegionClient(baseClient)
}

func TestGetMatchesByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		httpErr        error
		responseBody   string
		expectedResult []string
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			puuid:        "test-puuid",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			puuid:        "test-puuid",
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			puuid:          "test-puuid",
			statusCode:     http.StatusOK,
			responseBody:   matchListJSON,
			expectedResult: expectedMatchList,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := newTestRegionClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := rc.GetMatchesByPUUID(context.Background(), tt.puuid)

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
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetMatchByID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		matchID        string
		httpErr        error
		responseBody   string
		expectedResult Match
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			matchID:      "test-matchidnotfound",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			matchID:      "test-matchid",
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			matchID:        "test-matchid",
			statusCode:     http.StatusOK,
			responseBody:   matchJSON,
			expectedResult: expectedMatch,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := newTestRegionClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := rc.GetMatchByID(context.Background(), tt.matchID)

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

			// Marshal both to not run into timezone problems.
			expectedJSON, _ := json.Marshal(tt.expectedResult)
			jsonResp, _ := json.Marshal(resp)

			assert.Equal(t, expectedJSON, jsonResp)
		})
	}
}
