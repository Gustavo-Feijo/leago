package main

import (
	"context"
	"log"
	"os"

	"github.com/Gustavo-Feijo/leago"

	"github.com/Gustavo-Feijo/leago/examples"
	"github.com/Gustavo-Feijo/leago/regions"
)

func main() {
	// export RIOT_API_KEY=your_key_here
	// go run examples/valorant/main.go
	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		log.Fatal("RIOT_API_KEY not set")
	}
	client := leago.NewValRegionClient(
		regions.ValRegionBR,
		apiKey,
	)

	content, err := client.Val.Content.GetContent(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	examples.PrettyPrint(content)
}
