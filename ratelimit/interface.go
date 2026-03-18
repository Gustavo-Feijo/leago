package ratelimit

import (
	"context"
	"net/http"
)

// The ratelimiter interface will be used to define if a request to the RiotAPI should be done or not.
// Implementation details are up to the caller, so behavior constraints are very simple:
// Acquire will receive the context and a rate limit key. If error != nil, request is blocked.
type RateLimiter interface {
	Acquire(ctx context.Context, appKey string, methodKey string) error
}

// Syncer provides a interface that can be implemented to sync the rate limits from the returned headers.
// Returns all headers from the request, so the numbers of reqs during the windows can be evaluated from there.
// Headers also provide current usage, but usage as source of truth can be troublesome with paralelism.
type Syncer interface {
	Sync(ctx context.Context, headers http.Header, appKey string, methodKey string)
}

// Notifier provides a interface to handle any 429 error received.
// Returns all headers from the requests, so retry-after will be available.
// Can be used to add delays with the Retry-After.
type Notifier interface {
	NotifyTooManyRequests(ctx context.Context, headers http.Header, appKey string, methodKey string)
}
