package league

import (
	"context"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	listJSON = `{
	"leagueId":"123",
	"tier":"CHALLENGER",
	"name":"Test League",
	"queue":"RANKED_TFT",
	"entries":[
		{
			"freshBlood":true,
			"wins":10,
			"inactive":false,
			"veteran":false,
			"hotStreak":true,
			"rank":"I",
			"leaguePoints":500,
			"losses":2,
			"puuid":"testpuuid"
		}
	]}`

	expectedList = List{
		LeagueID: "123",
		Tier:     "CHALLENGER",
		Name:     "Test League",
		Queue:    "RANKED_TFT",
		Entries: []Item{
			{
				FreshBlood:   true,
				Wins:         10,
				Inactive:     false,
				Veteran:      false,
				HotStreak:    true,
				Rank:         "I",
				LeaguePoints: 500,
				Losses:       2,
				Puuid:        "testpuuid",
			},
		},
	}

	entriesJSON = `[
	{
		"puuid":"testpuuid",
		"leagueId":"123",
		"queueType":"RANKED_TFT",
		"tier":"DIAMOND",
		"rank":"I",
		"leaguePoints":100,
		"wins":10,
		"losses":5
	}]`

	expectedEntries = []Entry{
		{
			Puuid:        "testpuuid",
			LeagueID:     "123",
			QueueType:    "RANKED_TFT",
			Tier:         "DIAMOND",
			Rank:         "I",
			LeaguePoints: 100,
			Wins:         10,
			Losses:       5,
		},
	}

	ratedJSON = `[
	{
		"puuid":"testpuuid",
		"ratedTier":"BLUE",
		"ratedRating":200,
		"wins":20,
		"previousUpdateLadderPosition":1
	}]`

	expectedRated = []RatedLadderEntry{
		{
			Puuid:                        "testpuuid",
			RatedTier:                    LadderBlue,
			RatedRating:                  200,
			Wins:                         20,
			PreviousUpdateLadderPosition: 1,
		},
	}
)

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.PlatformBR1), "apiKey", internal.WithHTTP(mockDoer))

	return NewPlatformClient(baseClient), mockDoer
}

func TestGetMasterLeague(t *testing.T) {
	tests := []struct {
		name string

		opts []UpperLeagueOption

		statusCode   int
		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult List

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
			responseBody: `{"invalid json}`,

			wantErr: true,
		},
		{
			name: "success no filter",

			statusCode:   http.StatusOK,
			responseBody: listJSON,

			expectedPath: "/tft/league/v1/master",

			expectedResult: expectedList,
		},
		{
			name: "success with filter",

			opts: []UpperLeagueOption{WithQueueHighElo(QueueRankedTFT)},

			statusCode:   http.StatusOK,
			responseBody: listJSON,

			expectedPath: "/tft/league/v1/master",
			expectedQuery: map[string]string{
				"queue": "RANKED_TFT",
			},

			expectedResult: expectedList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetMasterLeague(
				context.Background(),
				tt.opts,
			)

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

			query := mockDoer.CapturedReq.URL.Query()
			for k, v := range tt.expectedQuery {
				assert.Equal(t, v, query.Get(k))
			}

			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetGrandmasterLeague(t *testing.T) {
	tests := []struct {
		name string

		opts         []UpperLeagueOption
		statusCode   int
		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult List
	}{
		{
			name: "without filter",

			statusCode:   http.StatusOK,
			responseBody: listJSON,

			expectedPath: "/tft/league/v1/grandmaster",

			expectedResult: expectedList,
		},
		{
			name: "with filter",

			opts: []UpperLeagueOption{WithQueueHighElo(QueueRankedTFTDoubleUP)},

			statusCode:   http.StatusOK,
			responseBody: listJSON,

			expectedPath: "/tft/league/v1/grandmaster",
			expectedQuery: map[string]string{
				"queue": "RANKED_TFT_DOUBLE_UP",
			},

			expectedResult: expectedList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(http.StatusOK, tt.responseBody)
			resp, err := pc.GetGrandmasterLeague(
				context.Background(),
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

func TestGetChallengerLeague(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, listJSON)
	resp, err := pc.GetChallengerLeague(context.Background(), nil)

	require.NoError(t, err)

	assert.Equal(t, "/tft/league/v1/challenger", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedList, resp)
}

func TestGetLeagueByID(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, listJSON)
	resp, err := pc.GetLeagueByID(context.Background(), "123")

	require.NoError(t, err)

	assert.Equal(t, "/tft/league/v1/leagues/123", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedList, resp)
}

func TestGetLeagueEntries(t *testing.T) {
	tests := []struct {
		name string

		tier     Tier
		division Division
		opts     []LeagueOption

		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult []Entry
	}{
		{
			name: "no filters",

			tier:     TierBronze,
			division: DivisionIV,

			responseBody: entriesJSON,

			expectedPath: "/tft/league/v1/entries/BRONZE/IV",

			expectedResult: expectedEntries,
		},
		{
			name: "all filters",

			tier:     TierDiamond,
			division: DivisionI,
			opts: []LeagueOption{
				WithPage(1),
				WithQueue(QueueRankedTFT),
			},

			responseBody: entriesJSON,

			expectedPath: "/tft/league/v1/entries/DIAMOND/I",
			expectedQuery: map[string]string{
				"page":  "1",
				"queue": "RANKED_TFT",
			},

			expectedResult: expectedEntries,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(http.StatusOK, tt.responseBody)
			resp, err := pc.GetLeagueEntries(
				context.Background(),
				tt.tier,
				tt.division,
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

func TestGetLeagueEntriesByPUUID(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, entriesJSON)

	resp, err := pc.GetLeagueEntriesByPUUID(
		context.Background(),
		"testpuuid",
	)

	require.NoError(t, err)

	assert.Equal(t, "/tft/league/v1/by-puuid/testpuuid", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedEntries, resp)
}

func TestGetRatedLadder(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, ratedJSON)

	resp, err := pc.GetRatedLadder(
		context.Background(),
		LadderQueueRankedTFTDoubleUP,
	)

	require.NoError(t, err)

	assert.Equal(t, "/tft/league/v1/rated-ladders/RANKED_TFT_DOUBLE_UP/top", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedRated, resp)
}
