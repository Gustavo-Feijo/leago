package status

import "leago/internal"

type PlatformClient struct {
	client *internal.Client
}

func NewPlatformClient(base *internal.Client) *PlatformClient {
	return &PlatformClient{
		base,
	}
}
