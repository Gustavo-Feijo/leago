package leagueexp

import (
	"context"
	"leago/internal"
	"leago/internal/mock"
	"leago/regions"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLeague(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		queue          Queue
		tier           Tier
		division       Division
		httpErr        error
		responseBody   string
		expectedResult LeagueResponse
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			queue:        QueueRankedSolo,
			tier:         TierGold,
			division:     DivisionI,
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			queue:        QueueRankedSolo,
			tier:         TierGold,
			division:     DivisionI,
			statusCode:   http.StatusOK,
			responseBody: `{"shouldbearray":true}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:       "success without mini series",
			queue:      QueueRankedSolo,
			tier:       TierGold,
			division:   DivisionI,
			statusCode: http.StatusOK,
			responseBody: `[{
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
			}]`,
			expectedResult: LeagueResponse{
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
			},
			wantErr: false,
		},
		{
			name:       "success with mini series",
			queue:      QueueRankedSolo,
			tier:       TierGold,
			division:   DivisionI,
			statusCode: http.StatusOK,
			responseBody: `[{
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
			}]`,
			expectedResult: LeagueResponse{
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
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)

			baseClient := internal.NewHttpClient(
				mockDoer,
				slog.Default(),
				string(regions.PlatformBR1),
				"apiKey",
			)

			pc := NewPlatformClient(baseClient)

			resp, err := pc.GetLeague(
				context.Background(),
				tt.queue,
				tt.tier,
				tt.division,
				[]GetLeagueOption{
					WithPage(1),
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
			require.NotNil(t, resp)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}
