// Package leago is a client library for the Riot Games API.
//
// It has a enforced separation between regions and platforms, so each client will be only accessed for it's routes.
//
// In order to instantiate the clients:
// [NewRegionClient] for region-routed APIs (match history, account lookup)
// [NewPlatformClient] for platform-routed APIs (league entries, summoners)
// [NewValRegionClient] for Valorant APIs (They use different regions).
//
// All three constructors accept functional options to change default behavior.
// [WithLimiter] to plug in a rate limiter. (Default is no-op).
// [WithLogger] for structured logging. (Default  discards in io).
// [WithClient] to override the HTTP client. (Default is http.DefaultClient).
//
// Supported games: League of Legends, TFT, Valorant and Legends of Runeterra.
package leago
