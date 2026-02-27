package challenges

import (
	"context"
	"fmt"
	"leago/internal"
	"leago/internal/mock"
	"leago/regions"
	"log/slog"
	"net/http"
	"testing"

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
			ChallengeIds:             []string{"1", "2"},
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

func newTestPlatformClient(statusCode int, responseBody string, httpErr error) *PlatformClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient)
}

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult []ConfigInfo
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
			responseBody:   configListJSON,
			expectedResult: expectedConfigList,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetConfig(context.Background())

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetConfigByID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult ConfigInfo
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
			responseBody:   configInfoJSON,
			expectedResult: expectedConfigInfo,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetConfigByID(context.Background(), 1)

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetLeaderboardByChallengeIDByLevel(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult Leaderboard
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
			responseBody:   leaderboardJSON,
			expectedResult: expectedLeaderboard,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetLeaderboardByChallengeIDByLevel(context.Background(), 1, TopLevelMaster, []GetLeaderboardOption{WithLimit(5)})

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetPercentiles(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult PercentileMap
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
			responseBody:   percentileMapJSON,
			expectedResult: expectedPercentileMap,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetPercentiles(context.Background())

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetPercentilesByChallengeID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult LevelPercentiles
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
			responseBody:   levelPercentilesJSON,
			expectedResult: expectedLevelPercentiles,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetPercentilesByChallengeID(context.Background(), 1)

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetPlayerInfoByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult PlayerInfo
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
			responseBody:   playerInfoJSON,
			expectedResult: expectedPlayerInfo,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetPlayerInfoByPUUID(context.Background(), "test-puuid")

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
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}
