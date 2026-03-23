package leago

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/Gustavo-Feijo/leago/api/lol"
	"github.com/Gustavo-Feijo/leago/api/lor"
	"github.com/Gustavo-Feijo/leago/api/riot"
	"github.com/Gustavo-Feijo/leago/api/tft"
	"github.com/Gustavo-Feijo/leago/api/val"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/ratelimit"
	nooprl "github.com/Gustavo-Feijo/leago/ratelimit/noop"
	"github.com/Gustavo-Feijo/leago/regions"
)

// Base client used by region and platform client.
type (
	baseClient struct {
		client  internal.Doer
		limiter ratelimit.RateLimiter
		logger  *slog.Logger
	}

	Option func(*baseClient)
)

// RegionClient provides access to all region related APIs.
type RegionClient struct {
	*baseClient
	Riot *riot.RegionClient
	Lor  *lor.RegionClient
	Lol  *lol.RegionClient
	Tft  *tft.RegionClient
}

// PlatformClient provides access to all platform related APIs.
type PlatformClient struct {
	*baseClient
	Lol *lol.PlatformClient
	Tft *tft.PlatformClient
}

// ValRegionClient provides access to all valorant region related APIs.
// Separated from RegionClient due to different region separation.
type ValRegionClient struct {
	*baseClient
	Val *val.RegionClient
}

// NewRegionClient returns a new client with access to the region specific APIs.
func NewRegionClient(region regions.Region, apiKey string, opts ...Option) *RegionClient {
	rc := &RegionClient{
		baseClient: newBaseClient(),
	}

	for _, opt := range opts {
		opt(rc.baseClient)
	}

	baseClient := internal.NewHTTPClient(
		string(region),
		apiKey,
		internal.WithHTTP(rc.client),
		internal.WithLimiter(rc.limiter),
		internal.WithLogger(rc.logger),
	)
	rc.Riot = riot.NewRegionClient(baseClient)
	rc.Lor = lor.NewRegionClient(baseClient)
	rc.Lol = lol.NewRegionClient(baseClient)
	rc.Tft = tft.NewRegionClient(baseClient)

	return rc
}

// NewPlatformClient returns a new client with access to the platform specific APIs.
func NewPlatformClient(platform regions.Platform, apiKey string, opts ...Option) *PlatformClient {
	pc := &PlatformClient{
		baseClient: newBaseClient(),
	}

	for _, opt := range opts {
		opt(pc.baseClient)
	}

	baseClient := internal.NewHTTPClient(
		string(platform),
		apiKey,
		internal.WithHTTP(pc.client),
		internal.WithLimiter(pc.limiter),
		internal.WithLogger(pc.logger),
	)
	pc.Lol = lol.NewPlatformClient(baseClient)
	pc.Tft = tft.NewPlatformClient(baseClient)

	return pc
}

// NewValRegionClient returns a new client with access to the valorant region specific APIs.
func NewValRegionClient(region regions.ValRegion, apiKey string, opts ...Option) *ValRegionClient {
	rc := &ValRegionClient{
		baseClient: newBaseClient(),
	}

	for _, opt := range opts {
		opt(rc.baseClient)
	}

	baseClient := internal.NewHTTPClient(
		string(region),
		apiKey,
		internal.WithHTTP(rc.client),
		internal.WithLimiter(rc.limiter),
		internal.WithLogger(rc.logger),
	)

	rc.Val = val.NewRegionClient(baseClient)

	return rc
}

func newBaseClient() *baseClient {
	return &baseClient{
		client:  http.DefaultClient,
		limiter: nooprl.NewNoopLimiter(),
		logger:  slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

// Override the default base http client.
func WithClient(doer internal.Doer) Option {
	return func(bc *baseClient) {
		bc.client = doer
	}
}

// Override the default noop limiter.
func WithLimiter(limiter ratelimit.RateLimiter) Option {
	return func(bc *baseClient) {
		bc.limiter = limiter
	}
}

// Override the default logger with discarded output.
func WithLogger(logger *slog.Logger) Option {
	return func(bc *baseClient) {
		bc.logger = logger
	}
}
