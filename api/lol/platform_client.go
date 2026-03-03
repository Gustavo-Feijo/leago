package lol

import (
	"log/slog"

	"github.com/Gustavo-Feijo/leago/api/lol/challenges"
	"github.com/Gustavo-Feijo/leago/api/lol/champion"
	"github.com/Gustavo-Feijo/leago/api/lol/championmastery"
	"github.com/Gustavo-Feijo/leago/api/lol/clash"
	"github.com/Gustavo-Feijo/leago/api/lol/league"
	"github.com/Gustavo-Feijo/leago/api/lol/leagueexp"
	"github.com/Gustavo-Feijo/leago/api/lol/status"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/regions"
)

type PlatformClient struct {
	Challenges      *challenges.PlatformClient
	Champion        *champion.PlatformClient
	ChampionMastery *championmastery.PlatformClient
	Clash           *clash.PlatformClient
	League          *league.PlatformClient
	LeagueExp       *leagueexp.PlatformClient
	Status          *status.PlatformClient
}

func NewPlatformClient(client internal.Doer, logger *slog.Logger, region regions.Platform, apiKey string) *PlatformClient {
	baseClient := internal.NewHttpClient(client, logger, string(region), apiKey)
	c := &PlatformClient{
		Challenges:      challenges.NewPlatformClient(baseClient),
		ChampionMastery: championmastery.NewPlatformClient(baseClient),
		Champion:        champion.NewPlatformClient(baseClient),
		Clash:           clash.NewPlatformClient(baseClient),
		League:          league.NewPlatformClient(baseClient),
		LeagueExp:       leagueexp.NewPlatformClient(baseClient),
		Status:          status.NewPlatformClient(baseClient),
	}
	return c
}
