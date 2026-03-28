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

	ddClient, err := leago.NewDDragonClient(
		realms.RealmNA,
		[]leago.DDragonOption{
			// Optional overrides
			leago.WithVersion("16.6.1"),
			leago.WithLocale("en_US"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	champion, err := ddClient.DDragon.Champion.GetChampionByID(context.Background(), "Aatrox")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(champion)
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

#### Type-safe Client
The API is split into different clients, so it has ```RegionClient```, ```PlatformClient```, ```ValRegionClient``` and ```DDragonClient``` instead of a single unified client.

This approach is intentional, the Riot API has separations for routing, like regions (Americas, Asia, Europe and Sea), platforms (NA1, BR1, KR1, etc), Valorant regions (Latam, NA) and DDragon Realms (PBE, PH). Using this approach each client only can access the endpoints for its routing, making it consistent for the type system.

#### Rate Limiting
The Ratelimiter is a interface, not concrete implementation.
Three interfaces are provided for implementation (```RateLimiter```, ```Sync``` and ```Notifier```), you can implement for whatever your use case needs.
A single-process app can use the included in-memory limiter. 
A distributed system can plug in a Redis based implementation without touching the client. 