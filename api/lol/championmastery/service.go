package championmastery

import (
	"context"
	"fmt"
	"leago/internal"
	"leago/options"
	"net/url"
)

const (
	MethodGetByPUUID           = "ChampionMastery.GetByPUUID"
	MethodGetByPUUIDTop        = "ChampionMastery.GetByPUUIDTop"
	MethodGetByPUUIDByChampion = "ChampionMastery.GetByPUUIDByChampion"
	MethodGetScoreByPUUID      = "ChampionMastery.GetScoreByPUUID"
)

// GetByPUUID returns the player champion mastery information got by their PUUID.
func (pc *PlatformClient) GetByPUUID(ctx context.Context, puuid string, opts ...options.PublicOption) (MasteryList, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/champion-masteries/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)

	return internal.AuthRequest[MasteryList](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(
			[]internal.RequestOption{
				internal.WithApiMethod(MethodGetByPUUID),
			},
			opts,
		)...,
	)
}

// GetByPUUIDTop returns the top X player champion mastery information got by their PUUID.
func (pc *PlatformClient) GetByPUUIDTop(ctx context.Context, puuid string, endpointOpts []GetByPUUIDTopOption, opts ...options.PublicOption) (MasteryList, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/top",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)

	// Adds endpoint specific options, like count.
	defaultOpts := append(
		[]internal.RequestOption{internal.WithApiMethod(MethodGetByPUUIDTop)},
		puuidTopOptionsToRequestOptions(endpointOpts)...,
	)

	return internal.AuthRequest[MasteryList](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetByPUUIDByChampion returns the player champion mastery information for a given champion got by their PUUID.
func (pc *PlatformClient) GetByPUUIDByChampion(ctx context.Context, puuid string, championID int64, opts ...options.PublicOption) (Mastery, error) {
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
		options.MergeOptions(
			[]internal.RequestOption{
				internal.WithApiMethod(MethodGetByPUUIDByChampion),
			},
			opts,
		)...,
	)
}

// GetScoreByPUUID returns a player total champion mastery score (Sum of individual champion mastery levels).
func (pc *PlatformClient) GetScoreByPUUID(ctx context.Context, puuid string, opts ...options.PublicOption) (MasteryScore, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/scores/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)

	return internal.AuthRequest[MasteryScore](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(
			[]internal.RequestOption{
				internal.WithApiMethod(MethodGetScoreByPUUID),
			},
			opts,
		)...,
	)
}
