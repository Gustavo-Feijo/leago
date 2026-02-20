package championmastery

import (
	"context"
	"leago/internal"
	"net/http"
)

type ChampionRotation struct {
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	FreeChampionIds              []int `json:"freeChampionIds"`
}

// GetRotation returns the current free champion rotation.
func (pc *PlatformClient) GetRotation(ctx context.Context, puuid string) (ChampionRotation, error) {
	endpoint := "/lol/platform/v3/champion-rotations"

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[ChampionRotation](ctx, pc.client.Http, pc.client.ApiKey, uri, http.MethodGet, nil)
}
