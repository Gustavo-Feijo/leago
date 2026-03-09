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

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient), mockDoer
}

var (
	expectedChallengerLeague = RawLeague{
		LeagueID: "league-1",
		Entries: []RawEloEntry{
			{
				PUUID:        "puuid-1",
				Rank:         "I",
				LeaguePoints: 1000,
				Wins:         200,
				Losses:       100,
				HotStreak:    true,
				Veteran:      true,
				FreshBlood:   false,
				Inactive:     false,
				MiniSeries:   nil,
			},
		},
		Tier:  "CHALLENGER",
		Name:  "test league",
		Queue: "RANKED_SOLO_5x5",
	}

	challengerJSON = `{
		"leagueId":"league-1",
		"entries":[{
			"puuid":"puuid-1",
			"rank":"I",
			"leaguePoints":1000,
			"wins":200,
			"losses":100,
			"hotStreak":true,
			"veteran":true,
			"freshBlood":false,
			"inactive":false
		}],
		"tier":"CHALLENGER",
		"name":"test league",
		"queue":"RANKED_SOLO_5x5"
	}`

	expectedGrandmasterLeague = RawLeague{
		LeagueID: "league-2",
		Entries: []RawEloEntry{
			{
				PUUID:        "puuid-2",
				Rank:         "I",
				LeaguePoints: 500,
				Wins:         150,
				Losses:       80,
				HotStreak:    false,
				Veteran:      false,
				FreshBlood:   true,
				Inactive:     false,
				MiniSeries:   nil,
			},
		},
		Tier:  "GRANDMASTER",
		Name:  "grandmaster league",
		Queue: "RANKED_FLEX_SR",
	}

	grandmasterJSON = `{
		"leagueId":"league-2",
		"entries":[{
			"puuid":"puuid-2",
			"rank":"I",
			"leaguePoints":500,
			"wins":150,
			"losses":80,
			"hotStreak":false,
			"veteran":false,
			"freshBlood":true,
			"inactive":false
		}],
		"tier":"GRANDMASTER",
		"name":"grandmaster league",
		"queue":"RANKED_FLEX_SR"
	}`

	expectedMasterLeague = RawLeague{
		LeagueID: "league-3",
		Entries: []RawEloEntry{
			{
				PUUID:        "puuid-3",
				Rank:         "I",
				LeaguePoints: 100,
				Wins:         50,
				Losses:       40,
				HotStreak:    false,
				Veteran:      true,
				FreshBlood:   false,
				Inactive:     false,
				MiniSeries:   nil,
			}},
		Tier:  "MASTER",
		Name:  "master league",
		Queue: "RANKED_SOLO_5x5",
	}

	masterJSON = `{
		"leagueId":"league-3",
		"entries":[{
			"puuid":"puuid-3",
			"rank":"I",
			"leaguePoints":100,
			"wins":50,
			"losses":40,
			"hotStreak":false,
			"veteran":true,
			"freshBlood":false,
			"inactive":false
		}],
		"tier":"MASTER",
		"name":"master league",
		"queue":"RANKED_SOLO_5x5"
	}`

	expectedLeagueWithoutMiniseries = []Entry{
		{
			LeagueID:     "league-1",
			SummonerID:   "summoner-1",
			PUUID:        "puuid-1",
			QueueType:    "RANKED_SOLO_5x5",
			Tier:         "GOLD",
			Rank:         "I",
			LeaguePoints: 75,
			Wins:         10,
			Losses:       5,
			HotStreak:    true,
			Veteran:      false,
			FreshBlood:   true,
			Inactive:     false,
			MiniSeries:   nil,
		},
	}

	leagueWithoutMiniseriesJSON = `[{
		"leagueId":"league-1",
		"summonerId":"summoner-1",
		"puuid":"puuid-1",
		"queueType":"RANKED_SOLO_5x5",
		"tier":"GOLD",
		"rank":"I",
		"leaguePoints":75,
		"wins":10,
		"losses":5,
		"hotStreak":true,
		"veteran":false,
		"freshBlood":true,
		"inactive":false
	}]`

	expectedLeagueWithMiniseries = []Entry{
		{
			LeagueID:     "league-2",
			SummonerID:   "summoner-2",
			PUUID:        "puuid-2",
			QueueType:    "RANKED_SOLO_5x5",
			Tier:         "GOLD",
			Rank:         "I",
			LeaguePoints: 100,
			Wins:         20,
			Losses:       10,
			HotStreak:    false,
			Veteran:      true,
			FreshBlood:   false,
			Inactive:     false,
			MiniSeries: &MiniSeries{
				Losses:   1,
				Progress: "WLN",
				Target:   3,
				Wins:     1,
			},
		},
	}

	leagueWithMiniseriesJSON = `[{
		"leagueId":"league-2",
		"summonerId":"summoner-2",
		"puuid":"puuid-2",
		"queueType":"RANKED_SOLO_5x5",
		"tier":"GOLD",
		"rank":"I",
		"leaguePoints":100,
		"wins":20,
		"losses":10,
		"hotStreak":false,
		"veteran":true,
		"freshBlood":false,
		"inactive":false,
		"miniSeries":{
			"losses":1,
			"progress":"WLN",
			"target":3,
			"wins":1
		}
	}]`

	expectedPlayerEntry = []Entry{
		{
			LeagueID:     "league-1",
			SummonerID:   "summoner-1",
			PUUID:        "puuid-1",
			QueueType:    "RANKED_SOLO_5x5",
			Tier:         "GOLD",
			Rank:         "II",
			LeaguePoints: 50,
			Wins:         30,
			Losses:       20,
			HotStreak:    false,
			Veteran:      false,
			FreshBlood:   false,
			Inactive:     false,
			MiniSeries:   nil,
		},
	}

	playerEntryJSON = `[{
		"leagueId":"league-1",
		"summonerId":"summoner-1",
		"puuid":"puuid-1",
		"queueType":"RANKED_SOLO_5x5",
		"tier":"GOLD",
		"rank":"II",
		"leaguePoints":50,
		"wins":30,
		"losses":20,
		"hotStreak":false,
		"veteran":false,
		"freshBlood":false,
		"inactive":false
	}]`

	expectedRawLeague = RawLeague{
		LeagueID: "league-1",
		Tier:     "SILVER",
		Name:     "Silver League",
		Queue:    "RANKED_SOLO_5x5",
		Entries: []RawEloEntry{
			{
				PUUID:        "puuid-1",
				Rank:         "I",
				LeaguePoints: 100,
				Wins:         40,
				Losses:       35,
				HotStreak:    true,
				Veteran:      false,
				FreshBlood:   false,
				Inactive:     false,
				MiniSeries: &MiniSeries{
					Losses:   0,
					Progress: "WNN",
					Target:   3,
					Wins:     1,
				},
			},
		},
	}

	rawLeagueJSON = `{
    	"leagueId":"league-1",
    	"tier":"SILVER",
    	"name":"Silver League",
    	"queue":"RANKED_SOLO_5x5",
    	"entries":[{
    	    "puuid":"puuid-1",
    	    "rank":"I",
    	    "leaguePoints":100,
    	    "wins":40,
    	    "losses":35,
    	    "hotStreak":true,
    	    "veteran":false,
    	    "freshBlood":false,
    	    "inactive":false,
    	    "miniSeries":{
    	        "losses":0,
    	        "progress":"WNN",
    	        "target":3,
    	        "wins":1
    	    }
    	}]
   	}`
)

func TestGetChallengerLeague(t *testing.T) {
	tests := []struct {
		name string

		queue Queue

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult RawLeague

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			queue: QueueRankedSolo,

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "unmatched json",

			queue: QueueRankedSolo,

			statusCode:   http.StatusOK,
			responseBody: `["invalid"]`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			queue: QueueRankedSolo,

			statusCode:   http.StatusOK,
			responseBody: challengerJSON,

			expectedPath: "/lol/league/v4/challengerleagues/by-queue/RANKED_SOLO_5x5",

			expectedResult: expectedChallengerLeague,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetChallengerLeague(context.Background(), tt.queue)

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

func TestGetGrandmasterLeague(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, grandmasterJSON)
	resp, err := pc.GetGrandmasterLeague(context.Background(), QueueRankedFlexSR)

	require.NoError(t, err)

	assert.Equal(t, "/lol/league/v4/grandmasterleagues/by-queue/RANKED_FLEX_SR", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedGrandmasterLeague, resp)
}

func TestGetMasterLeague(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, masterJSON)
	resp, err := pc.GetMasterLeague(context.Background(), QueueRankedSolo)

	require.NoError(t, err)

	assert.Equal(t, "/lol/league/v4/masterleagues/by-queue/RANKED_SOLO_5x5", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedMasterLeague, resp)
}

func TestGetLeagueEntries(t *testing.T) {
	tests := []struct {
		name string

		queue    Queue
		tier     Tier
		division Division
		opts     []GetLeagueOption

		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult []Entry
	}{
		{
			name: "without mini series",

			queue:    QueueRankedSolo,
			tier:     TierGold,
			division: DivisionI,

			responseBody: leagueWithoutMiniseriesJSON,

			expectedPath: "/lol/league/v4/entries/RANKED_SOLO_5x5/GOLD/I",

			expectedResult: expectedLeagueWithoutMiniseries,
		},
		{
			name: "with mini series",

			queue:    QueueRankedFlexSR,
			tier:     TierDiamond,
			division: DivisionIV,
			opts:     []GetLeagueOption{WithPage(5)},

			responseBody: leagueWithMiniseriesJSON,

			expectedPath: "/lol/league/v4/entries/RANKED_FLEX_SR/DIAMOND/IV",
			expectedQuery: map[string]string{
				"page": "5",
			},

			expectedResult: expectedLeagueWithMiniseries,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(http.StatusOK, tt.responseBody)
			resp, err := pc.GetLeagueEntries(
				context.Background(),
				tt.queue,
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
	pc, mockDoer := newTestPlatformClient(http.StatusOK, playerEntryJSON)
	resp, err := pc.GetLeagueEntriesByPUUID(context.Background(), "puuid-1")

	require.NoError(t, err)

	assert.Equal(t, "/lol/league/v4/entries/by-puuid/puuid-1", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedPlayerEntry, resp)
}

func TestGetLeagueByID(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, rawLeagueJSON)
	resp, err := pc.GetLeagueByID(context.Background(), "league-1")

	require.NoError(t, err)

	assert.Equal(t, "/lol/league/v4/leagues/league-1", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedRawLeague, resp)
}
