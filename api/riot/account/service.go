package account

import (
	"context"
	"fmt"
	"leago/internal"
	"net/http"
	"net/url"
)

type (
	Account struct {
		Puuid    string `json:"puuid"`
		GameName string `json:"gameName"`
		TagLine  string `json:"tagLine"`
	}

	ActiveRegion struct {
		Puuid       string `json:"puuid"`
		Game        string `json:"game"`
		ActiveShard string `json:"activeShard"`
	}

	ActiveShard struct {
		Puuid       string `json:"puuid"`
		Game        string `json:"game"`
		ActiveShard string `json:"activeShard"`
	}
)

// GetActiveRegion returns the user active region by their puuid and game.
func (rc *RegionClient) GetActiveRegionByPUUID(ctx context.Context, game, puuid string) (ActiveRegion, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/region/by-game/%s/by-puuid/%s",
		url.PathEscape(game),
		url.PathEscape(puuid),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[ActiveRegion](ctx, rc.client.Http, rc.client.ApiKey, uri, http.MethodGet, nil)
}

// GetActiveShard returns the user active shard by their puuid and game.
func (rc *RegionClient) GetActiveShardByPUUID(ctx context.Context, game, puuid string) (ActiveShard, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/active-shards/by-game/%s/by-puuid/%s",
		url.PathEscape(game),
		url.PathEscape(puuid),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[ActiveShard](ctx, rc.client.Http, rc.client.ApiKey, uri, http.MethodGet, nil)
}

// GetByPUUID returns the user account by their puuid.
func (rc *RegionClient) GetByPUUID(ctx context.Context, puuid string) (Account, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/accounts/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Account](ctx, rc.client.Http, rc.client.ApiKey, uri, http.MethodGet, nil)
}

// GetByRiotID returns the user account by their gamename and tagline.
func (rc *RegionClient) GetByRiotID(ctx context.Context, gameName, tagLine string) (Account, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/accounts/by-riot-id/%s/%s",
		url.PathEscape(gameName),
		url.PathEscape(tagLine),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Account](ctx, rc.client.Http, rc.client.ApiKey, uri, http.MethodGet, nil)
}
