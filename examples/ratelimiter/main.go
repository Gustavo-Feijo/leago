package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Gustavo-Feijo/leago"

	"github.com/Gustavo-Feijo/leago/examples"
	memorylimiter "github.com/Gustavo-Feijo/leago/ratelimit/memory"
	"github.com/Gustavo-Feijo/leago/regions"
)

func main() {
	// export RIOT_API_KEY=your_key_here
	// go run examples/ratelimiter/main.go
	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		log.Fatal("RIOT_API_KEY not set")
	}

	// Ratelimiter implementation (You can provide your own if it satisfies the interface).
	ratelimiter := memorylimiter.NewMemoryLimiter(
		// Limit margin to avoid 429 due to clock skew.
		memorylimiter.WithLimitSafetyMargin(0.7),

		// Add buffer to interval resets to avoid 429.
		// Below it just make it wait one more second.
		memorylimiter.WithIntervalSafetyMargin(time.Second),
	)

	client := leago.NewPlatformClient(
		regions.PlatformNA1,
		apiKey,
		leago.WithLimiter(ratelimiter),
	)

	// The default implementation has syncing to get the actual limits, so making a first request just to get it.
	rotation, err := client.Lol.Champion.GetRotation(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := range 50 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, err := client.Lol.Champion.GetRotation(context.Background())
			if err != nil {
				log.Println(err)
			}
			log.Println("Made request ", i)
		}(i)
	}
	wg.Wait()

	examples.PrettyPrint(rotation)
}
