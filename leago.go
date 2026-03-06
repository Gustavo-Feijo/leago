package leago

import (
	"log/slog"
	"net/http"

	"github.com/Gustavo-Feijo/leago/api/lol"
	"github.com/Gustavo-Feijo/leago/api/lor"
	"github.com/Gustavo-Feijo/leago/api/riot"
	"github.com/Gustavo-Feijo/leago/api/tft"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/regions"
)

// Base client used by region and platform client.
type (
	baseClient struct {
		client internal.Doer
		logger *slog.Logger
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

// NewRegionClient returns a new client with access to the region specific APIs.
func NewRegionClient(region regions.Region, apiKey string, opts ...Option) *RegionClient {
	rc := &RegionClient{
		baseClient: newBaseClient(),
	}

	for _, opt := range opts {
		opt(rc.baseClient)
	}

	rc.Riot = riot.NewRegionClient(rc.client, rc.logger, region, apiKey)
	rc.Lor = lor.NewRegionClient(rc.client, rc.logger, region, apiKey)
	rc.Lol = lol.NewRegionClient(rc.client, rc.logger, region, apiKey)
	rc.Tft = tft.NewRegionClient(rc.client, rc.logger, region, apiKey)

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

	pc.Lol = lol.NewPlatformClient(pc.client, pc.logger, platform, apiKey)
	pc.Tft = tft.NewPlatformClient(pc.client, pc.logger, platform, apiKey)

	return pc
}

func newBaseClient() *baseClient {
	return &baseClient{
		client: http.DefaultClient,
		logger: slog.New(slog.DiscardHandler),
	}
}

// Override the default base http client.
func WithClient(doer internal.Doer) Option {
	return func(bc *baseClient) {
		bc.client = doer
	}
}

// Override the default logger with discarded output.
func WithLogger(logger *slog.Logger) Option {
	return func(bc *baseClient) {
		bc.logger = logger
	}
}
