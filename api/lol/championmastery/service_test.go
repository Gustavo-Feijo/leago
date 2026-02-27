package championmastery

import (
	"context"
	"fmt"
	"leago/internal"
	"leago/internal/mock"
	"leago/regions"
	"log/slog"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedMastery = Mastery{
		Puuid:                        "test-puuid",
		ChampionPointsUntilNextLevel: 100,
		ChestGranted:                 true,
		ChampionID:                   266,
		LastPlayTime:                 1700000000000,
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
			"lastPlayTime": 1700000000000,
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

	expectedMasteries = MasteryList{expectedMastery}
	masteriesJSON     = fmt.Sprintf("[%s]", masteryJSON)
)

func TestGetByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		httpErr        error
		responseBody   string
		expectedResult MasteryList
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
			responseBody:   masteriesJSON,
			expectedResult: expectedMasteries,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)

			resp, err := pc.GetByPUUID(context.Background(), tt.puuid)

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

func TestGetByPUUIDTop(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		httpErr        error
		responseBody   string
		expectedResult MasteryList
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
			responseBody:   masteriesJSON,
			expectedResult: expectedMasteries,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)

			defaultCount := 5
			resp, err := pc.GetByPUUIDTop(context.Background(), tt.puuid, []GetByPUUIDTopOption{WithCount(defaultCount)})

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

			countParam := mockDoer.CapturedReq.URL.Query().Get("count")
			val, paramErr := strconv.ParseInt(countParam, 10, 0)
			assert.Nil(t, paramErr)
			assert.Equal(t, int(val), defaultCount)

			require.NotNil(t, resp)
			assert.Equal(t, resp, tt.expectedResult)
		})
	}
}

func TestGetByPUUIDByChampion(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		championId     int64
		httpErr        error
		responseBody   string
		expectedResult Mastery
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			puuid:        "test-puuid",
			championId:   12,
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			puuid:        "test-puuid",
			championId:   125,
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			puuid:          "test-puuid",
			championId:     266,
			statusCode:     http.StatusOK,
			responseBody:   masteryJSON,
			expectedResult: expectedMastery,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)

			resp, err := pc.GetByPUUIDByChampion(context.Background(), tt.puuid, tt.championId)

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

func TestGetScoreByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		httpErr        error
		responseBody   string
		expectedResult MasteryScore
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
			responseBody:   "15",
			expectedResult: 15,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)

			baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)

			resp, err := pc.GetScoreByPUUID(context.Background(), tt.puuid)

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
