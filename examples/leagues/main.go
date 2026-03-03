package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Gustavo-Feijo/leago"

	"github.com/Gustavo-Feijo/leago/api/lol/leagueexp"
	"github.com/Gustavo-Feijo/leago/examples"
	"github.com/Gustavo-Feijo/leago/regions"
)

func main() {
	// export RIOT_API_KEY=your_key_here
	// go run examples/leagues/main.go
	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		log.Fatal("RIOT_API_KEY not set")
	}
	client := leago.NewPlatformClient(
		regions.PlatformNA1,
		apiKey,
	)

	for i := 1; i <= 3; i++ {
		rotation, err := client.Lol.LeagueExp.GetLeague(
			context.Background(),
			leagueexp.QueueRankedSolo,
			leagueexp.TierDiamond,
			leagueexp.DivisionI,
			[]leagueexp.GetLeagueOption{
				leagueexp.WithPage(i),
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Page ", i, ":")
		examples.PrettyPrint(rotation)
	}
}
