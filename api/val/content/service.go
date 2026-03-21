package content

import (
	"context"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetContent = "ValContent.GetContent"
)

// GetContent returns all static game content (agents, maps, skins, etc).
func (rc *RegionClient) GetContent(
	ctx context.Context,
	endpointOpts []ContentOption,
	opts ...options.PublicOption,
) (Content, error) {
	endpoint := "/val/content/v1/contents"

	defaultOpts := append(
		[]internal.RequestOption{internal.WithAPIMethod(MethodGetContent)},
		contentOptionsToRequestOptions(endpointOpts)...,
	)

	uri := rc.client.GetURL(endpoint)

	return internal.AuthRequest[Content](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
