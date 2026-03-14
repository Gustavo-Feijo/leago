package ratelimit

import "context"

// The ratelimiter interface will be used to define if a request to the RiotAPI should be done or not.
// Implementation details are up to the caller, so behavior constraints are very simple:
// Acquire will receive the context and a rate limit key. If error != nil, request is blocked.
type RateLimiter interface {
	Acquire(ctx context.Context, key string) error
}
