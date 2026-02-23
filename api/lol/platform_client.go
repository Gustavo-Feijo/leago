package lol

import (
	"leago/api/lol/champion"
	"leago/api/lol/championmastery"
	"leago/api/lol/clash"
	"leago/internal"
	"leago/regions"
	"log/slog"
)

type PlatformClient struct {
	Champion        *champion.PlatformClient
	ChampionMastery *championmastery.PlatformClient
	Clash           *clash.PlatformClient
}

func NewPlatformClient(client internal.Doer, logger *slog.Logger, region regions.Platform, apiKey string) *PlatformClient {
	baseClient := internal.NewHttpClient(client, logger, string(region), apiKey)
	c := &PlatformClient{
		ChampionMastery: championmastery.NewPlatformClient(baseClient),
		Champion:        champion.NewPlatformClient(baseClient),
		Clash:           clash.NewPlatformClient(baseClient),
	}
	return c
}
