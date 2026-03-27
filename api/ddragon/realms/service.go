package realms

import (
	"context"
	"fmt"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetRealm = "DDRealms.GetRealm"
)

// GetRealm returns the information for a given realm.
func (rc *RegionClient) GetRealm(
	ctx context.Context,
	realm Realm,
	opts ...options.PublicOption,
) (DDragonResponse, error) {
	endpoint := fmt.Sprintf("/realms/%s.json", realm)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetRealm),
	}

	uri := rc.client.GetDDragonURL(endpoint)
	return internal.Request[DDragonResponse](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
