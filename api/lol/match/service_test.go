package match

import (
	"context"
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
			DataVersion:  "2",
			MatchID:      "BR1_123",
			Participants: []string{"puuid1"},
		},
		Info: MatchInfo{
			EndOfGameResult:    "GameComplete",
			GameCreation:       internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
			GameDuration:       internal.SecondsDuration{Duration: time.Minute * 30},
			GameEndTimestamp:   internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 30, 0, 0, time.UTC)},
			GameID:             123,
			GameMode:           "CLASSIC",
			GameName:           "teambuilder-match-123",
			GameStartTimestamp: internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
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

	matchJSON = `{
		"metadata": {
			"dataVersion": "2",
			"matchId": "BR1_123",
			"participants": ["puuid1"]
		},
		"info": {
			"endOfGameResult": "GameComplete",
			"gameCreation": 1767225600000,
			"gameDuration": 1800,
			"gameEndTimestamp": 1767227400000,
			"gameId": 123,
			"gameMode": "CLASSIC",
			"gameName": "teambuilder-match-123",
			"gameStartTimestamp": 1767225600000,
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

	expectedReplays = Replays{
		Total: 2,
		MatchFileURLs: []string{
			"https://replay1",
			"https://replay2",
		},
	}

	replaysJSON = `{
		"total": 2,
		"matchFileURLs": [
			"https://replay1",
			"https://replay2"
		]
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
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.RegionAmericas), "apiKey", internal.WithHTTP(mockDoer))

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

			matchID: "badreturn",

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"}`,

			wantErr: true,
		},
		{
			name: "success",

			matchID: "BR1_123",

			statusCode:   http.StatusOK,
			responseBody: matchJSON,

			expectedPath:   "/lol/match/v5/matches/BR1_123",
			expectedResult: expectedMatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)
			resp, err := rc.GetMatchByID(context.Background(), tt.matchID)

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
	}{
		{
			name: "no filter",

			puuid: "nofilterpuuid",

			responseBody: `["BR1_1","BR1_2"]`,

			expectedPath: "/lol/match/v5/matches/by-puuid/nofilterpuuid/ids",
		},
		{
			name: "all filters",

			puuid: "allfilterspuuid",
			opts: []GetMatchesByPUUIDOption{
				WithStart(0),
				WithCount(20),
				WithQueue(420),
				WithType("ranked"),
				WithStartTime(start),
				WithEndTime(end),
			},

			responseBody: `["BR1_1","BR1_2"]`,

			expectedPath: "/lol/match/v5/matches/by-puuid/allfilterspuuid/ids",
			expectedQuery: map[string]string{
				"start":     "0",
				"count":     "20",
				"queue":     "420",
				"type":      "ranked",
				"startTime": strconv.FormatInt(start, 10),
				"endTime":   strconv.FormatInt(end, 10),
			},
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

			assert.Equal(t, []string{"BR1_1", "BR1_2"}, resp)
		})
	}
}

func TestGetReplaysByPUUID(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, replaysJSON)
	resp, err := rc.GetReplaysByPUUID(context.Background(), "testpuuid")

	require.NoError(t, err)

	assert.Equal(t, "/lol/match/v5/matches/by-puuid/testpuuid/replays", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedReplays, resp)
}

func TestGetMatchTimelineByID(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, timelineJSON)
	resp, err := rc.GetMatchTimelineByID(context.Background(), "BR1_123")

	require.NoError(t, err)

	assert.Equal(t, "/lol/match/v5/matches/BR1_123/timeline", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedTimeline, resp)
}
