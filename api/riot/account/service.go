package account

import (
	"context"
	"fmt"
	"leago/internal"
	"leago/options"
	"net/url"
)

const (
	MethodGetActiveRegionByPUUID = "Account.GetActiveRegionByPUUID"
	MethodGetActiveShardByPUUID  = "Account.GetActiveShardByPUUID"
	MethodGetByPUUID             = "Account.GetByPUUID"
	MethodGetByRiotID            = "Account.GetByRiotID"
)

// GetActiveRegion returns the user active region by their puuid and game.
func (rc *RegionClient) GetActiveRegionByPUUID(ctx context.Context, game ActiveRegionGame, puuid string, opts ...options.PublicOption) (ActiveRegion, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/region/by-game/%s/by-puuid/%s",
		game,
		url.PathEscape(puuid),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[ActiveRegion](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(
			[]internal.RequestOption{
				internal.WithApiMethod(MethodGetActiveRegionByPUUID),
			}, opts,
		)...,
	)
}

// GetActiveShard returns the user active shard by their puuid and game.
func (rc *RegionClient) GetActiveShardByPUUID(ctx context.Context, game ActiveShardGame, puuid string, opts ...options.PublicOption) (ActiveShard, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/active-shards/by-game/%s/by-puuid/%s",
		game,
		url.PathEscape(puuid),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[ActiveShard](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(
			[]internal.RequestOption{
				internal.WithApiMethod(MethodGetActiveShardByPUUID),
			}, opts,
		)...,
	)
}

// GetByPUUID returns the user account by their puuid.
func (rc *RegionClient) GetByPUUID(ctx context.Context, puuid string, opts ...options.PublicOption) (Account, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/accounts/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Account](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(
			[]internal.RequestOption{
				internal.WithApiMethod(MethodGetByPUUID),
			}, opts,
		)...,
	)
}

// GetByRiotID returns the user account by their gamename and tagline.
func (rc *RegionClient) GetByRiotID(ctx context.Context, gameName, tagLine string, opts ...options.PublicOption) (Account, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/accounts/by-riot-id/%s/%s",
		url.PathEscape(gameName),
		url.PathEscape(tagLine),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Account](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(
			[]internal.RequestOption{
				internal.WithApiMethod(MethodGetByRiotID),
			},
			opts,
		)...,
	)
}
