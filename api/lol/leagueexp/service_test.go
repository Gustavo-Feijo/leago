package leagueexp

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
)

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.PlatformBR1), "apiKey", internal.WithHTTP(mockDoer))

	return NewPlatformClient(baseClient), mockDoer
}

func TestGetLeague(t *testing.T) {
	tests := []struct {
		name string

		queue    Queue
		tier     Tier
		division Division
		opts     []GetLeagueOption

		statusCode   int
		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult []Entry

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			queue:    QueueRankedSolo,
			tier:     TierGold,
			division: DivisionI,

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "unmatched json",

			queue:    QueueRankedSolo,
			tier:     TierGold,
			division: DivisionI,

			statusCode:   http.StatusOK,
			responseBody: `{"shouldbearray":true}`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "without mini series",

			queue:    QueueRankedSolo,
			tier:     TierGold,
			division: DivisionI,

			statusCode:   http.StatusOK,
			responseBody: leagueWithoutMiniseriesJSON,

			expectedPath: "/lol/league-exp/v4/entries/RANKED_SOLO_5x5/GOLD/I",

			expectedResult: expectedLeagueWithoutMiniseries,
		},
		{
			name: "with mini series",

			queue:    QueueRankedFlexSR,
			tier:     TierDiamond,
			division: DivisionIV,
			opts:     []GetLeagueOption{WithPage(1)},

			statusCode:   http.StatusOK,
			responseBody: leagueWithMiniseriesJSON,

			expectedPath: "/lol/league-exp/v4/entries/RANKED_FLEX_SR/DIAMOND/IV",
			expectedQuery: map[string]string{
				"page": "1",
			},

			expectedResult: expectedLeagueWithMiniseries,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetLeague(
				context.Background(),
				tt.queue,
				tt.tier,
				tt.division,
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
