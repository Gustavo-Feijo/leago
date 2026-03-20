package lol

import (
	"github.com/Gustavo-Feijo/leago/api/lol/challenges"
	"github.com/Gustavo-Feijo/leago/api/lol/champion"
	"github.com/Gustavo-Feijo/leago/api/lol/championmastery"
	"github.com/Gustavo-Feijo/leago/api/lol/clash"
	"github.com/Gustavo-Feijo/leago/api/lol/league"
	"github.com/Gustavo-Feijo/leago/api/lol/leagueexp"
	"github.com/Gustavo-Feijo/leago/api/lol/spectator"
	"github.com/Gustavo-Feijo/leago/api/lol/status"
	"github.com/Gustavo-Feijo/leago/api/lol/summoner"
	"github.com/Gustavo-Feijo/leago/internal"
)

type PlatformClient struct {
	Challenges      *challenges.PlatformClient
	Champion        *champion.PlatformClient
	ChampionMastery *championmastery.PlatformClient
	Clash           *clash.PlatformClient
	League          *league.PlatformClient
	LeagueExp       *leagueexp.PlatformClient
	Spectator       *spectator.PlatformClient
	Status          *status.PlatformClient
	Summoner        *summoner.PlatformClient
}

func NewPlatformClient(baseClient *internal.Client) *PlatformClient {
	c := &PlatformClient{
		Challenges:      challenges.NewPlatformClient(baseClient),
		Champion:        champion.NewPlatformClient(baseClient),
		ChampionMastery: championmastery.NewPlatformClient(baseClient),
		Clash:           clash.NewPlatformClient(baseClient),
		League:          league.NewPlatformClient(baseClient),
		LeagueExp:       leagueexp.NewPlatformClient(baseClient),
		Spectator:       spectator.NewPlatformClient(baseClient),
		Status:          status.NewPlatformClient(baseClient),
		Summoner:        summoner.NewPlatformClient(baseClient),
	}
	return c
}
