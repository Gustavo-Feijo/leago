package challenges

import (
	"context"
	"fmt"
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
	expectedConfigInfo = ConfigInfo{
		ID:             1,
		LocalizedNames: map[string]map[string]string{"en_US": {"name": "Test"}},
		State:          StateEnabled,
		Tracking:       TrackingLifetime,
		StartTimestamp: 1000,
		EndTimestamp:   2000,
		Leaderboard:    true,
		Thresholds:     map[Level]float64{LevelGold: 100.0},
	}

	configInfoJSON = `{
		"id": 1,
		"localizedNames": {"en_US": {"name": "Test"}},
		"state": "ENABLED",
		"tracking": "LIFETIME",
		"startTimestamp": 1000,
		"endTimestamp": 2000,
		"leaderboard": true,
		"thresholds": {"GOLD": 100.0}
	}`

	expectedConfigList = []ConfigInfo{expectedConfigInfo}
	configListJSON     = fmt.Sprintf("[%s]", configInfoJSON)

	expectedLeaderboard = Leaderboard{
		{Puuid: "puuid-1", Value: 9.5, Position: 1},
		{Puuid: "puuid-2", Value: 8.0, Position: 2},
	}
	leaderboardJSON = `[
		{"puuid": "puuid-1", "value": 9.5, "position": 1},
		{"puuid": "puuid-2", "value": 8.0, "position": 2}
	]`

	expectedPercentileMap = PercentileMap{
		1: {LevelGold: 0.5},
		2: {LevelSilver: 0.3},
	}
	percentileMapJSON = `{
		"1": {"GOLD": 0.5},
		"2": {"SILVER": 0.3}
	}`

	expectedLevelPercentiles = LevelPercentiles{LevelGold: 0.5, LevelSilver: 0.3}
	levelPercentilesJSON     = `{"GOLD": 0.5, "SILVER": 0.3}`

	expectedPlayerInfo = PlayerInfo{
		Challenges: []PlayerChallenges{
			{
				Percentiles:    0.9,
				PlayersInLevel: 100,
				AchievedTime:   1234567890,
				Value:          42.0,
				ChallengeID:    1,
				Level:          LevelGold,
				Position:       5,
			},
		},
		Preferences: PlayerClientPreferences{
			BannerAccent:             "banner1",
			Title:                    "title1",
			ChallengeIDs:             []string{"1", "2"},
			CrestBorder:              "crest1",
			PrestigeCrestBorderLevel: 3,
		},
		TotalPoints: ChallengePoints{
			Level:      "GOLD",
			Current:    500,
			Max:        1000,
			Percentile: 0.75,
		},
		CategoryPoints: map[string]ChallengePoints{
			"COLLECTION": {Level: "SILVER", Current: 200, Max: 400, Percentile: 0.5},
		},
	}
	playerInfoJSON = `{
		"challenges": [
			{
				"percentile": 0.9,
				"playersInLevel": 100,
				"achievedTime": 1234567890,
				"value": 42.0,
				"challengeId": 1,
				"level": "GOLD",
				"position": 5
			}
		],
		"preferences": {
			"bannerAccent": "banner1",
			"title": "title1",
			"challengeIds": ["1", "2"],
			"crestBorder": "crest1",
			"prestigeCrestBorderLevel": 3
		},
		"totalPoints": {
			"level": "GOLD",
			"current": 500,
			"max": 1000,
			"percentile": 0.75
		},
		"categoryPoints": {
			"COLLECTION": {"level": "SILVER", "current": 200, "max": 400, "percentile": 0.5}
		}
	}`
)

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient), mockDoer
}

// All service calls are coupled to AuthRequest, a single test will handled the error cases, since it's basically testing how it handles the return from there.

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedResult []ConfigInfo

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
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			statusCode:   http.StatusOK,
			responseBody: configListJSON,

			expectedResult: expectedConfigList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetConfig(context.Background())

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

			assert.Equal(t, "/lol/challenges/v1/challenges/config", mockDoer.CapturedReq.URL.Path)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetConfigByID(t *testing.T) {
	tests := []struct {
		name string

		challengeID int64

		expectedPath string
	}{
		{
			name: "single digit id",

			challengeID: 1,

			expectedPath: "/lol/challenges/v1/challenges/1/config",
		},
		{
			name: "large id",

			challengeID: 123456789,

			expectedPath: "/lol/challenges/v1/challenges/123456789/config",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(http.StatusOK, configInfoJSON)
			resp, err := pc.GetConfigByID(context.Background(), tt.challengeID)

			require.NoError(t, err)

			assert.Equal(t, tt.expectedPath, mockDoer.CapturedReq.URL.Path)
			assert.Equal(t, expectedConfigInfo, resp)
		})
	}
}

func TestGetLeaderboardByChallengeIDByLevel(t *testing.T) {
	tests := []struct {
		name string

		challengeID int64
		level       TopLevel
		opts        []GetLeaderboardOption

		expectedPath     string
		expectedRawQuery string
	}{
		{
			name: "no options",

			challengeID: 1,
			level:       TopLevelMaster,
			opts:        []GetLeaderboardOption{},

			expectedPath: "/lol/challenges/v1/challenges/1/leaderboards/by-level/MASTER",
		},
		{
			name: "with limit",

			challengeID: 1,
			level:       TopLevelMaster,
			opts:        []GetLeaderboardOption{WithLimit(5)},

			expectedPath:     "/lol/challenges/v1/challenges/1/leaderboards/by-level/MASTER",
			expectedRawQuery: "limit=5",
		},
		{
			name: "grandmaster level",

			challengeID: 99,
			level:       TopLevelGrandmaster,
			opts:        []GetLeaderboardOption{},

			expectedPath: "/lol/challenges/v1/challenges/99/leaderboards/by-level/GRANDMASTER",
		},
		{
			name: "challenger level",

			challengeID: 99,
			level:       TopLevelChallenger,
			opts:        []GetLeaderboardOption{},

			expectedPath: "/lol/challenges/v1/challenges/99/leaderboards/by-level/CHALLENGER",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(http.StatusOK, leaderboardJSON)
			resp, err := pc.GetLeaderboardByChallengeIDByLevel(context.Background(), tt.challengeID, tt.level, tt.opts)

			require.NoError(t, err)

			assert.Equal(t, tt.expectedPath, mockDoer.CapturedReq.URL.Path)
			assert.Equal(t, tt.expectedRawQuery, mockDoer.CapturedReq.URL.RawQuery)
			assert.Equal(t, expectedLeaderboard, resp)
		})
	}
}

func TestGetPercentiles(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, percentileMapJSON)
	resp, err := pc.GetPercentiles(context.Background())

	require.NoError(t, err)

	assert.Equal(t, "/lol/challenges/v1/challenges/percentiles", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedPercentileMap, resp)
}

func TestGetPercentilesByChallengeID(t *testing.T) {
	tests := []struct {
		name string

		challengeID int64

		expectedPath string
	}{
		{
			name: "firstidtest",

			challengeID: 1,

			expectedPath: "/lol/challenges/v1/challenges/1/percentiles",
		},
		{
			name: "secondidtest",

			challengeID: 123456789,

			expectedPath: "/lol/challenges/v1/challenges/123456789/percentiles",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(http.StatusOK, levelPercentilesJSON)
			resp, err := pc.GetPercentilesByChallengeID(context.Background(), tt.challengeID)

			require.NoError(t, err)

			assert.Equal(t, tt.expectedPath, mockDoer.CapturedReq.URL.Path)
			assert.Equal(t, expectedLevelPercentiles, resp)
		})
	}
}

func TestGetPlayerInfoByPUUID(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, playerInfoJSON)
	resp, err := pc.GetPlayerInfoByPUUID(context.Background(), "test-puuid")

	require.NoError(t, err)

	assert.Equal(t, "/lol/challenges/v1/player-data/test-puuid", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedPlayerInfo, resp)
}
