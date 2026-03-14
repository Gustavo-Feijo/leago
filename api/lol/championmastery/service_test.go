package championmastery

import (
	"context"
	"fmt"
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
	expectedMastery = Mastery{
		Puuid:                        "test-puuid",
		ChampionPointsUntilNextLevel: 100,
		ChestGranted:                 true,
		ChampionID:                   266,
		LastPlayTime:                 internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
		ChampionLevel:                7,
		ChampionPoints:               123456,
		ChampionPointsSinceLastLevel: 5000,
		MarkRequiredForNextLevel:     2,
		ChampionSeasonMilestone:      1,
		NextSeasonMilestone: NextSeasonMilestones{
			RequireGradeCounts: map[string]int{
				"S": 1,
			},
			RewardMarks: 1,
			Bonus:       false,
			RewardConfig: RewardConfig{
				RewardValue:   "CHEST",
				RewardType:    "HEXTECH",
				MaximumReward: 1,
			},
		},
		TokensEarned:    2,
		MilestoneGrades: []string{"S", "A"},
	}

	masteryJSON = `
	{
			"puuid": "test-puuid",
			"championPointsUntilNextLevel": 100,
			"chestGranted": true,
			"championId": 266,
			"lastPlayTime": 1767225600000,
			"championLevel": 7,
			"championPoints": 123456,
			"championPointsSinceLastLevel": 5000,
			"markRequiredForNextLevel": 2,
			"championSeasonMilestone": 1,
			"nextSeasonMilestone": {
				"requireGradeCounts": {
					"S": 1
				},
				"rewardMarks": 1,
				"bonus": false,
				"rewardConfig": {
					"rewardValue": "CHEST",
					"rewardType": "HEXTECH",
					"maximumReward": 1
				}
			},
			"tokensEarned": 2,
			"milestoneGrades": ["S", "A"]
		}
	`

	expectedMasteries = []Mastery{expectedMastery}
	masteriesJSON     = fmt.Sprintf("[%s]", masteryJSON)
)

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient), mockDoer
}

func TestGetByPUUID(t *testing.T) {
	tests := []struct {
		name string

		puuid string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult []Mastery

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			puuid: "nonexistentpuuid",

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			expectedPath: "/lol/champion-mastery/v4/champion-masteries/by-puuid/nonexistentpuuid",

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "invalid json",

			puuid: "test-puuid",

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,

			expectedPath: "/lol/champion-mastery/v4/champion-masteries/by-puuid/test-puuid",

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			puuid: "test-puuid",

			statusCode:   http.StatusOK,
			responseBody: masteriesJSON,

			expectedPath: "/lol/champion-mastery/v4/champion-masteries/by-puuid/test-puuid",

			expectedResult: expectedMasteries,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetByPUUID(context.Background(), tt.puuid)

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

func TestGetByPUUIDTop(t *testing.T) {
	tests := []struct {
		name string

		puuid string
		opts  []GetByPUUIDTopOption

		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult []Mastery
	}{
		{
			name: "no filters",

			puuid: "test-puuid",

			responseBody: masteriesJSON,

			expectedPath: "/lol/champion-mastery/v4/champion-masteries/by-puuid/test-puuid/top",

			expectedResult: expectedMasteries,
		},
		{
			name: "with count",

			puuid: "test-puuid",

			opts: []GetByPUUIDTopOption{
				WithCount(5),
			},

			responseBody: masteriesJSON,

			expectedPath: "/lol/champion-mastery/v4/champion-masteries/by-puuid/test-puuid/top",
			expectedQuery: map[string]string{
				"count": "5",
			},

			expectedResult: expectedMasteries,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(http.StatusOK, tt.responseBody)
			resp, err := pc.GetByPUUIDTop(context.Background(), tt.puuid, tt.opts)

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

func TestGetByPUUIDByChampion(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, masteryJSON)
	resp, err := pc.GetByPUUIDByChampion(context.Background(), "test-puuid", 266)

	require.NoError(t, err)

	assert.Equal(t, "/lol/champion-mastery/v4/champion-masteries/by-puuid/test-puuid/by-champion/266", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedMastery, resp)
}

func TestGetScoreByPUUID(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, "15")
	resp, err := pc.GetScoreByPUUID(context.Background(), "test-puuid")

	require.NoError(t, err)

	assert.Equal(t, "/lol/champion-mastery/v4/scores/by-puuid/test-puuid", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, MasteryScore(15), resp)
}
