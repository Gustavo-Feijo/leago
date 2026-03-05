package league

import (
	"context"
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
		"wreviousUpdateLadderPosition":1
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

func newTestPlatformClient(statusCode int, responseBody string, httpErr error) *PlatformClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient)
}

func TestGetMasterLeague(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult List
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
			responseBody: `{"invalid json}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   listJSON,
			expectedResult: expectedList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetMasterLeague(
				context.Background(),
				[]UpperLeagueOption{
					WithQueueHighElo(QueueRankedTFT),
				},
			)

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

func TestGetGrandmasterLeague(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult List
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
			responseBody: `{"invalid json}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   listJSON,
			expectedResult: expectedList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetGrandmasterLeague(context.Background(), nil)

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

func TestGetChallengerLeague(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult List
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
			responseBody: `{"invalid json}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   listJSON,
			expectedResult: expectedList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetChallengerLeague(context.Background(), nil)

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

func TestGetLeagueByID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult List
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
			responseBody: `{"invalid json}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   listJSON,
			expectedResult: expectedList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetLeagueByID(context.Background(), "123")

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

func TestGetLeagueEntries(t *testing.T) {

	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult []Entry
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
			responseBody: `{"invalid json}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   entriesJSON,
			expectedResult: expectedEntries,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetLeagueEntries(
				context.Background(),
				TierDiamond,
				DivisionI,
				[]LeagueOption{
					WithPage(1),
					WithQueue(QueueRankedTFT),
				},
			)

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

func TestGetLeagueEntriesByPUUID(t *testing.T) {

	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult []Entry
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
			responseBody: `{"invalid json}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   entriesJSON,
			expectedResult: expectedEntries,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)

			resp, err := pc.GetLeagueEntriesByPUUID(
				context.Background(),
				"testpuuid",
			)

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

func TestGetRatedLadder(t *testing.T) {

	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult []RatedLadderEntry
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
			responseBody: `{"invalid json}`,
			wantErr:      true,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   ratedJSON,
			expectedResult: expectedRated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)

			resp, err := pc.GetRatedLadder(
				context.Background(),
				LadderQueueRankedTFTDoubleUP,
			)

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
