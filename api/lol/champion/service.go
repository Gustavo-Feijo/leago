package champion

import (
	"context"
	"leago/internal"
)

type ChampionRotation struct {
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	FreeChampionIds              []int `json:"freeChampionIds"`
}

// GetRotation returns the current free champion rotation.
func (pc *PlatformClient) GetRotation(ctx context.Context) (ChampionRotation, error) {
	endpoint := "/lol/platform/v3/champion-rotations"

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[ChampionRotation](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod("Champion.GetRotation"),
	)
}
