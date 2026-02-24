package championmastery

import (
	"context"
	"fmt"
	"leago/internal"
	"net/url"
	"strconv"
)

const (
	MethodGetByPUUID           = "ChampionMastery.GetByPUUID"
	MethodGetByPUUIDTop        = "ChampionMastery.GetByPUUIDTop"
	MethodGetByPUUIDByChampion = "ChampionMastery.GetByPUUIDByChampion"
	MethodGetScoreByPUUID      = "ChampionMastery.GetScoreByPUUID"
)

// GetByPUUID returns the player champion mastery information got by their PUUID.
func (pc *PlatformClient) GetByPUUID(ctx context.Context, puuid string) (MasteryList, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/champion-masteries/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[MasteryList](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod(MethodGetByPUUID),
	)
}

// GetByPUUIDTop returns the top X player champion mastery information got by their PUUID.
func (pc *PlatformClient) GetByPUUIDTop(ctx context.Context, puuid string, count int) (MasteryList, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/top",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)

	params := make(map[string]string)
	params["count"] = strconv.Itoa(count)

	return internal.AuthRequest[MasteryList](
		ctx,
		pc.client,
		uri,
		internal.WithParams(params),
		internal.WithApiMethod(MethodGetByPUUIDTop))
}

// GetByPUUIDByChampion returns the player champion mastery information for a given champion got by their PUUID.
func (pc *PlatformClient) GetByPUUIDByChampion(ctx context.Context, puuid string, championID int64) (Mastery, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%d",
		url.PathEscape(puuid),
		championID,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Mastery](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod(MethodGetByPUUIDByChampion),
	)
}

// GetScoreByPUUID returns a player total champion mastery score (Sum of individual champion mastery levels).
func (pc *PlatformClient) GetScoreByPUUID(ctx context.Context, puuid string) (MasteryScore, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/scores/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[MasteryScore](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod(MethodGetScoreByPUUID),
	)
}
