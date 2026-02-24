# Leago

Leago is a simple client for the Riot APIs, providing clean API access to the Riot API and Data Dragon.

## Example usage
```go
func main() {
	apiKey := "RGAPI-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

	rClient := leago.NewRegionClient(
		regions.RegionAmericas,
		apiKey,
	)

	ctx := context.TODO()
	account, err := rClient.Riot.Account.GetByRiotID(ctx, "GameName", "TagLine")
	if err != nil {
		panic(err)
	}

	pClient := leago.NewPlatformClient(
		regions.PlatformNA1,
		apiKey,
	)

	mastery, err := pClient.Lol.ChampionMastery.GetByPUUID(ctx, account.Puuid)
	if err != nil {
		panic(err)
	}

	fmt.Println(mastery)
}
```

More usage examples can be found and executed inside ```examples/```.

## Decisions
It works with multiple client instances, with each client being coupled to its region or platform.

The goal is to add reliable completion and separation between the clients, that way a Platform Client (NA1 for example) can't be used with Region specific endpoints.

This project comes from some challenges I faced while developing hobby projects using the Riot API, the goal here is to provide a working client for it's access.
