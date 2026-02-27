package challenges

import (
	"context"
	"fmt"
	"leago/internal"
	"leago/options"
	"net/url"
)

const (
	MethodGetConfig                          = "Challenges.GetConfig"
	MethodGetConfigByID                      = "Challenges.GetConfigByID"
	MethodGetLeaderboardByChallengeIDByLevel = "Challenges.GetLeaderboardByChallengeIDByLevel"
	MethodGetPercentiles                     = "Challenges.GetPercentiles"
	MethodGetPercentilesByChallengeID        = "Challenges.GetPercentilesByChallengeID"
	MethodGetPlayerInfoByPUUID               = "Challenges.GetPlayerInfoByPUUID"
)

// GetConfig returns the challenges config information.
func (pc *PlatformClient) GetConfig(
	ctx context.Context,
	opts ...options.PublicOption,
) ([]ConfigInfo, error) {
	endpoint := "/lol/challenges/v1/challenges/config"

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetConfig),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[[]ConfigInfo](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetConfigByID returns a given challeng config information.
func (pc *PlatformClient) GetConfigByID(
	ctx context.Context,
	challengeID int64,
	opts ...options.PublicOption,
) (ConfigInfo, error) {
	endpoint := fmt.Sprintf(
		"/lol/challenges/v1/challenges/%d/config",
		challengeID,
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetConfigByID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[ConfigInfo](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetLeaderboardByChallengeIDByLevel returns the top players for a given challengeID and level.
func (pc *PlatformClient) GetLeaderboardByChallengeIDByLevel(
	ctx context.Context,
	challengeID int64,
	level TopLevel,
	endpointOpts []GetLeaderboardOption,
	opts ...options.PublicOption,
) (Leaderboard, error) {
	endpoint := fmt.Sprintf(
		"/lol/challenges/v1/challenges/%d/leaderboards/by-level/%s",
		challengeID,
		level,
	)

	defaultOpts := append(
		[]internal.RequestOption{internal.WithApiMethod(MethodGetLeaderboardByChallengeIDByLevel)},
		getLeaderboardOptionsToRequestOptions(endpointOpts)...,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Leaderboard](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetPercentiles returns a map of percentiles of players that had achieved a given challenge level.
func (pc *PlatformClient) GetPercentiles(
	ctx context.Context,
	opts ...options.PublicOption,
) (PercentileMap, error) {
	endpoint := "/lol/challenges/v1/challenges/percentiles"

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetPercentiles),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[PercentileMap](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetPercentilesByChallengeID returns the level percentile distribution for a given challengeID.
func (pc *PlatformClient) GetPercentilesByChallengeID(
	ctx context.Context,
	challengeID int64,
	opts ...options.PublicOption,
) (LevelPercentiles, error) {
	endpoint := fmt.Sprintf(
		"/lol/challenges/v1/challenges/%d/percentiles",
		challengeID,
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetPercentilesByChallengeID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[LevelPercentiles](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetPlayerInfoByPUUID returns the player progress in the challenges.
func (pc *PlatformClient) GetPlayerInfoByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) (PlayerInfo, error) {
	endpoint := fmt.Sprintf(
		"/lol/challenges/v1/player-data/%s",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetPlayerInfoByPUUID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[PlayerInfo](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
