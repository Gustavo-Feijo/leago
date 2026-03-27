package item

import (
	"context"
	"fmt"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetItems = "DDItem.GetItems"
)

// GetItems returns a list of all items.
func (rc *RegionClient) GetItems(
	ctx context.Context,
	opts ...options.PublicOption,
) (ItemResponse, error) {
	endpoint := fmt.Sprintf("/cdn/%s/data/%s/item.json", rc.version, rc.language)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetItems),
	}

	uri := rc.client.GetDDragonURL(endpoint)
	return internal.Request[ItemResponse](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
