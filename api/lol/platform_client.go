package lol

import (
	"leago/api/lol/challenges"
	"leago/api/lol/champion"
	"leago/api/lol/championmastery"
	"leago/api/lol/clash"
	"leago/api/lol/league"
	"leago/api/lol/leagueexp"
	"leago/internal"
	"leago/regions"
	"log/slog"
)

type PlatformClient struct {
	Challenges      *challenges.PlatformClient
	Champion        *champion.PlatformClient
	ChampionMastery *championmastery.PlatformClient
	Clash           *clash.PlatformClient
	League          *league.PlatformClient
	LeagueExp       *leagueexp.PlatformClient
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
	}
	return c
}
