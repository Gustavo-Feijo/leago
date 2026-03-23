package internal

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Gustavo-Feijo/leago/ratelimit"
)

const apiTokenHeader = "X-Riot-Token" // #nosec Header name, not credential

// AuthRequest calls Request with a automatically attached client API Key.
// It's the standard entrypoint for most of the Riot API Calls.
func AuthRequest[T any](ctx context.Context, client *Client, uri string, opts ...RequestOption) (T, error) {
	return Request[T](
		ctx,
		client,
		uri,
		append(opts, withAPIKey(client.apiKey))...,
	)
}

// Request builds and executes an HTTP request.
// Is able to handler rate limiting, limits syncing from headers and notifying limiter on 429 (Depends on limiter implementation).
// Use AuthRequest instead if a API key is required.
func Request[T any](ctx context.Context, client *Client, uri string, opts ...RequestOption) (T, error) {
	var ro requestOptions
	for _, o := range opts {
		o(&ro)
	}

	req, err := buildRequest(ctx, uri, &ro)
	if err != nil {
		var zero T
		return zero, err
	}

	return executor[T](ctx, client, req, &ro)
}

// buildRequest mounts a new http request with all passed options.
func buildRequest(ctx context.Context, uri string, opts *requestOptions) (*http.Request, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	// Add any query param needed, some of the APIs use it for filtering.
	query := u.Query()
	for key, val := range opts.params {
		query.Add(key, val)
	}

	u.RawQuery = query.Encode()

	var bodyReader io.Reader = http.NoBody
	if opts.body != nil {
		b, err := json.Marshal(opts.body)
		if err != nil {
			return nil, err
		}
		bodyReader = strings.NewReader(string(b))
	}

	method := opts.httpMethod
	if method == "" {
		method = http.MethodGet
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if opts.body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if opts.apiKey != "" {
		req.Header.Set(apiTokenHeader, opts.apiKey)
	}

	return req, nil
}

// executor is a wrapper to the do function, holds all checks before actually executing the request.
func executor[T any](ctx context.Context, client *Client, req *http.Request, ro *requestOptions) (T, error) {
	appKey := ratelimit.AppKey(client.routePrefix)
	methodKey := ratelimit.MethodKey(client.routePrefix, ro.apiMethod)
	rlErr := client.limiter.Acquire(ctx, appKey, methodKey)
	if rlErr != nil {
		var zero T
		return zero, rlErr
	}

	result, headers, err := do[T](client, req, ro)

	if headers != nil {
		if s, ok := client.limiter.(ratelimit.Syncer); ok {
			s.Sync(ctx, headers, appKey, methodKey)
		}
	}

	// If too many requests notify rate limiter (If appliable).
	if err != nil {
		var riotErr *RiotError
		if errors.As(err, &riotErr) && riotErr.StatusCode == http.StatusTooManyRequests {
			if n, ok := client.limiter.(ratelimit.Notifier); ok {
				n.NotifyTooManyRequests(ctx, headers, appKey, methodKey)
			}
		}
	}

	return result, err
}

// do Executes the request itself and handles the status and unmarshal.
func do[T any](client *Client, req *http.Request, ro *requestOptions) (T, http.Header, error) {
	var respData T

	logger := client.Logger.With(
		"apiMethod", ro.apiMethod,
		"httpMethod", ro.httpMethod,
		"uri", req.URL.String(),
		"route", client.routePrefix,
	)

	resp, err := client.HTTP.Do(req)
	if err != nil {
		logger.Error("request failed", "error", err)

		if resp != nil {
			return respData, resp.Header, err
		}

		return respData, nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body", "error", err)
		return respData, resp.Header, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logger.Warn("non-OK HTTP status", "status", resp.StatusCode)
		return respData, resp.Header, &RiotError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       strings.TrimSpace(string(body)),
		}
	}

	if err := json.Unmarshal(body, &respData); err != nil {
		logger.Error("failed to unmarshal response body", "error", err)
		return respData, resp.Header, err
	}

	return respData, resp.Header, nil
}
