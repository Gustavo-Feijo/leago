package tft

import (
	"github.com/Gustavo-Feijo/leago/api/tft/league"
	"github.com/Gustavo-Feijo/leago/api/tft/status"
	"github.com/Gustavo-Feijo/leago/api/tft/summoner"
	"github.com/Gustavo-Feijo/leago/internal"
)

type PlatformClient struct {
	League   *league.PlatformClient
	Status   *status.PlatformClient
	Summoner *summoner.PlatformClient
}

func NewPlatformClient(baseClient *internal.Client) *PlatformClient {
	c := &PlatformClient{
		League:   league.NewPlatformClient(baseClient),
		Status:   status.NewPlatformClient(baseClient),
		Summoner: summoner.NewPlatformClient(baseClient),
	}
	return c
}
