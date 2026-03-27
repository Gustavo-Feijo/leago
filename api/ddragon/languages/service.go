package languages

import (
	"context"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetLanguages = "DDLanguages.GetLanguages"
)

// GetLanguages returns a list of all available languages.
func (rc *RegionClient) GetLanguages(
	ctx context.Context,
	opts ...options.PublicOption,
) ([]string, error) {
	endpoint := "/cdn/languages.json"

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetLanguages),
	}

	uri := rc.client.GetDDragonURL(endpoint)
	return internal.Request[[]string](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
