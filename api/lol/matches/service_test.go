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
	matchJSON = `{
		"metadata": {
			"dataVersion": "2",
			"matchId": "BR1_123",
			"participants": ["puuid1"]
		},
		"info": {
			"endOfGameResult": "GameComplete",
			"gameCreation": 1,
			"gameDuration": 1800,
			"gameEndTimestamp": 2,
			"gameId": 123,
			"gameMode": "CLASSIC",
			"gameName": "teambuilder-match-123",
			"gameStartTimestamp": 1,
			"gameType": "MATCHED_GAME",
			"gameVersion": "26.1.1",
			"mapId": 11,
			"participants": [],
			"platformId": "BR1",
			"queueId": 420,
			"teams": [],
			"tournamentCode": ""
		}
	}`

	expectedMatch = Match{
		Metadata: MatchMetadata{
			DataVersion:  "2",
			MatchID:      "BR1_123",
			Participants: []string{"puuid1"},
		},
		Info: MatchInfo{
			EndOfGameResult:    "GameComplete",
			GameCreation:       1,
			GameDuration:       1800,
			GameEndTimestamp:   2,
			GameID:             123,
			GameMode:           "CLASSIC",
			GameName:           "teambuilder-match-123",
			GameStartTimestamp: 1,
			GameType:           "MATCHED_GAME",
			GameVersion:        "26.1.1",
			MapID:              11,
			Participants:       []MatchParticipant{},
			PlatformID:         "BR1",
			QueueID:            420,
			Teams:              []Team{},
			TournamentCode:     "",
		},
	}

	replaysJSON = `{
		"total": 2,
		"matchFileURLs": [
			"https://replay1",
			"https://replay2"
		]
	}`

	expectedReplays = Replays{
		Total: 2,
		MatchFileURLs: []string{
			"https://replay1",
			"https://replay2",
		},
	}

	timelineJSON = `{
		"metadata": {
			"dataVersion": "2",
			"matchId": "BR1_123",
			"participants": ["puuid1"]
		},
		"info": {
			"endOfGameResult": "GameComplete",
			"frameInterval": 60000,
			"gameId": 123,
			"participants": [],
			"frames": []
		}
	}`

	expectedTimeline = Timeline{
		Metadata: MetadataTimeLine{
			DataVersion:  "2",
			MatchID:      "BR1_123",
			Participants: []string{"puuid1"},
		},
		Info: InfoTimeLine{
			EndOfGameResult: "GameComplete",
			FrameInterval:   60000,
			GameID:          123,
			Participants:    []ParticipantTimeLine{},
			Frames:          []FrameTimeLine{},
		},
	}
)

func newTestRegionClient(statusCode int, responseBody string, httpErr error) *RegionClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.RegionAmericas), "apiKey")

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

			expectedJson, _ := json.Marshal(tt.expectedResult)
			jsonResp, _ := json.Marshal(resp)

			assert.Equal(t, expectedJson, jsonResp)
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
					WithQueue(420),
					WithType("ranked"),
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

func TestGetReplaysByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   string
		expectedResult Replays
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
			responseBody: `{"bad json"}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   replaysJSON,
			expectedResult: expectedReplays,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := newTestRegionClient(tt.statusCode, tt.responseBody, nil)

			resp, err := rc.GetReplaysByPUUID(context.Background(), "puuid")

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

			expectedJson, _ := json.Marshal(tt.expectedResult)
			jsonResp, _ := json.Marshal(resp)

			assert.Equal(t, expectedJson, jsonResp)
		})
	}
}

func TestGetMatchTimelineByID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   string
		expectedResult Timeline
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
			responseBody: `{"bad json"}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   timelineJSON,
			expectedResult: expectedTimeline,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := newTestRegionClient(tt.statusCode, tt.responseBody, nil)

			resp, err := rc.GetMatchTimelineByID(context.Background(), "BR1_123")

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

			expectedJson, _ := json.Marshal(tt.expectedResult)
			jsonResp, _ := json.Marshal(resp)

			assert.Equal(t, expectedJson, jsonResp)
		})
	}
}
