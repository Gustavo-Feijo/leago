package lol

import (
	"leago/api/lol/championmastery"
	"leago/internal"
	"leago/regions"
)

type PlatformClient struct {
	ChampionMastery *championmastery.PlatformClient
}

func NewPlatformClient(client internal.Doer, region regions.Platform, apiKey string) *PlatformClient {
	baseClient := internal.NewHttpClient(client, string(region), apiKey)
	c := &PlatformClient{
		ChampionMastery: championmastery.NewPlatformClient(baseClient),
	}
	return c
}
