package profileicon

import (
	"context"
	"fmt"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetProfileIcons = "DDProfileIcon.GetProfileIcons"
)

// GetProfileIcons returns a list of all summonerspells.
func (rc *RegionClient) GetProfileIcons(
	ctx context.Context,
	opts ...options.PublicOption,
) (ProfileIconResponse, error) {
	endpoint := fmt.Sprintf("/cdn/%s/data/%s/profileicon.json", rc.version, rc.language)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetProfileIcons),
	}

	uri := rc.client.GetDDragonURL(endpoint)
	return internal.Request[ProfileIconResponse](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
