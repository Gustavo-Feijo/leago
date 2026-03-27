# Leago

[![Go Reference](https://pkg.go.dev/badge/github.com/Gustavo-Feijo/leago.svg)](https://pkg.go.dev/github.com/Gustavo-Feijo/leago) 
[![CI](https://github.com/Gustavo-Feijo/leago/actions/workflows/ci.yml/badge.svg)](https://github.com/Gustavo-Feijo/leago/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/Gustavo-Feijo/leago/branch/main/graph/badge.svg?token=QDCH0O72Y8)](https://codecov.io/github/Gustavo-Feijo/leago) 
[![Go Report Card](https://goreportcard.com/badge/github.com/Gustavo-Feijo/leago)](https://goreportcard.com/report/github.com/Gustavo-Feijo/leago)

Leago is a simple client for the Riot APIs, providing clean API access to the Riot API and Data Dragon.

## Installation

```bash
go get github.com/Gustavo-Feijo/leago
```

## Example usage
```go
func main() {
	apiKey := "RGAPI-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

	// Create a regional client (For account endpoint).
	rClient := leago.NewRegionClient(
		regions.RegionAmericas,
		apiKey,
	)

	ctx := context.TODO()

	// Fetch account by Riot ID.
	account, err := rClient.Riot.Account.GetByRiotID(ctx, "GameName", "TagLine")
	if err != nil {
		panic(err)
	}

	// Create a platform client (For game specific endpoint).
	pClient := leago.NewPlatformClient(
		regions.PlatformNA1,
		apiKey,
	)

	// Fetch champion mastery for the PUUID.
	mastery, err := pClient.Lol.ChampionMastery.GetByPUUID(ctx, account.Puuid)
	if err != nil {
		panic(err)
	}

	fmt.Println(mastery)

	// Create a Valorant regional client.
	valClient := leago.NewValRegionClient(
		regions.ValRegionBR,
		apiKey,
	)

	content, err := valClient.Val.Content.GetContent(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
}
```

More usage examples can be found and executed inside ```examples/```.

## Ratelimiting
Riot APIs enforce strict rate limits. A Ratelimiting interface is provided in order to plug in your own rate limiting strategy. (Or use the provided in-memory option).

The interfaces that can be implemented to work with the client are present in ```ratelimit/interface.go```.

Three interfaces can be implemented for the Ratelimiter:
- RateLimiter: Acquire(ctx context.Context, appKey string, methodKey string) error
  - Simple function to acquire a rate limit, if any error is returned the request will error.
- Syncer: Sync(ctx context.Context, headers http.Header, appKey string, methodKey string)
  - After finishing a request its headers will be passed to the limiter, so it can be used to Sync the limits.
  - Recommended to Sync only the limits bounds, since the usage can be wrong due to ongoing requests. (Can be used to limits if running purely synchronously).
- Notifier: NotifyTooManyRequests(ctx context.Context, headers http.Header, appKey string, methodKey string)
  - If a request receive a 429 call this method passing the headers so the Retry-After header can be used to adjust limits.
  
A simple in-memory implementation with all three interfaces can be found in ```ratelimit/memory/ratelimit.go```.

Example usage:
```go
func main() {
	apiKey := "RGAPI-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	ratelimiter := memorylimiter.NewMemoryLimiter(
		// Simple methods to reduce the limits usage/delay limit resets. 
		memorylimiter.WithLimitSafetyMargin(0.7),
		memorylimiter.WithIntervalSafetyMargin(time.Second),
	)

	client := leago.NewPlatformClient(
		regions.PlatformNA1,
		apiKey,
		leago.WithLimiter(ratelimiter),
	)

	// The in-memory implementation has syncing to get the actual limits, so making a first request initialize them.
	rotation, err := client.Lol.Champion.GetRotation(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
```

## Decisions
It works with multiple client instances, with each client being coupled to its region or platform.

The goal is to add reliable completion and separation between the clients, that way a Platform Client (NA1 for example) can't be used with Region specific endpoints.

This project comes from some challenges I faced while developing hobby projects using the Riot API, the goal here is to provide a working client for its access.
