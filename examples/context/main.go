package main

import (
	"context"
	"leago"
	"leago/examples"
	"leago/regions"
	"log"
	"os"
	"time"
)

func main() {
	// export RIOT_API_KEY=your_key_here
	// go run examples/context/main.go
	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		log.Fatal("RIOT_API_KEY not set")
	}

	client := leago.NewPlatformClient(
		regions.PlatformNA1,
		apiKey,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	rotation, err := client.Lol.Champion.GetRotation(ctx)
	if err != nil {
		log.Fatal(err)
	}

	examples.PrettyPrint(rotation)
}
