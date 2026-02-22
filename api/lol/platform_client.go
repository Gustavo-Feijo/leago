package lol

import (
	"leago/api/lol/championmastery"
	"leago/internal"
	"leago/regions"
	"log/slog"
)

type PlatformClient struct {
	ChampionMastery *championmastery.PlatformClient
}

func NewPlatformClient(client internal.Doer, logger *slog.Logger, region regions.Platform, apiKey string) *PlatformClient {
	baseClient := internal.NewHttpClient(client, logger, string(region), apiKey)
	c := &PlatformClient{
		ChampionMastery: championmastery.NewPlatformClient(baseClient),
	}
	return c
}
