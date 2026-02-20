package leago

import (
	"leago/api/lol"
	"leago/api/riot"
	"leago/internal"
	"leago/regions"
	"net/http"
)

// Base client used by region and platform client.
type (
	baseClient struct {
		client internal.Doer
	}

	Option func(*baseClient)

	// RegionClient provides access to all region related APIs.
	RegionClient struct {
		*baseClient
		Riot *riot.RegionClient
	}

	// PlatformClient provides access to all platform related APIs.
	PlatformClient struct {
		*baseClient
		Lol *lol.PlatformClient
	}
)

// NewRegionClient returns a new client with access to the region specific APIs.
func NewRegionClient(region regions.Region, apiKey string, opts ...Option) *RegionClient {
	rc := &RegionClient{
		baseClient: &baseClient{
			client: http.DefaultClient,
		},
	}

	for _, opt := range opts {
		opt(rc.baseClient)
	}

	rc.Riot = riot.NewRegionClient(rc.client, region, apiKey)

	return rc
}

// NewPlatformClient returns a new client with access to the platform specific APIs.
func NewPlatformClient(platform regions.Platform, apiKey string, opts ...Option) *PlatformClient {
	pc := &PlatformClient{
		baseClient: &baseClient{
			client: http.DefaultClient,
		},
	}

	for _, opt := range opts {
		opt(pc.baseClient)
	}

	pc.Lol = lol.NewPlatformClient(pc.client, platform, apiKey)

	return pc
}

// Override the default base http client.
func WithClient(doer internal.Doer) Option {
	return func(bc *baseClient) {
		bc.client = doer
	}
}
