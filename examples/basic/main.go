package main

import (
	"context"
	"leago"
	"leago/examples"
	"leago/regions"
	"log"
	"os"
)

func main() {
	// export RIOT_API_KEY=your_key_here
	// go run examples/basic/main.go
	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		log.Fatal("RIOT_API_KEY not set")
	}
	client := leago.NewPlatformClient(
		regions.PlatformNA1,
		apiKey,
	)

	rotation, err := client.Lol.Champion.GetRotation(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	examples.PrettyPrint(rotation)
}
