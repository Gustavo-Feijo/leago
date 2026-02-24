package main

import (
	"context"
	"leago"
	"leago/options"
	"leago/regions"
	"log"
	"log/slog"
	"os"
)

func main() {
	// export RIOT_API_KEY=your_key_here
	// go run examples/logging/main.go
	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		log.Fatal("RIOT_API_KEY not set")
	}

	handler := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	na1handler := handler.With("region", regions.RegionAmericas)

	client := leago.NewPlatformClient(
		regions.PlatformNA1,
		apiKey,
		leago.WithLogger(na1handler),
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := client.Lol.Champion.GetRotation(ctx); err != nil {
		// error expected due to cancelled context
	}

	if _, err := client.Lol.Champion.GetRotation(ctx, options.WithApiMethod("OverrideDefaultLogMethod")); err != nil {
		// error expected due to cancelled context
	}

}
